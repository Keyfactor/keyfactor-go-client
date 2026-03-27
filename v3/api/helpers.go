package api

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/spbsoluble/go-pkcs12"
	"github.com/youmark/pkcs8"
	"reflect"
)

// UnpackPEM extracts the private key, certificate, and CA certificates from PEM-formatted data.
// If the private key is encrypted (PKCS#8 encrypted format), it will be decrypted using the provided password.
//
// Parameters:
//   - pemData: The PEM data as a string, *string, or []byte (may be base64-encoded)
//   - password: The password used for decrypting the encrypted private key
//
// Returns:
//   - privateKey: The decrypted private key in PEM format
//   - certificate: The leaf certificate in PEM format
//   - caCertificates: A slice of CA certificates in PEM format (if any)
//   - err: An error that describes why the unpacking failed, if any
func UnpackPEM(pemData interface{}, password string) (
	privateKey, certificate string,
	caCertificates []string,
	err error,
) {
	var pemBytes []byte

	// Convert pemData to []byte
	switch v := pemData.(type) {
	case string:
		pemBytes = []byte(v)
		// Try base64 decode
		if decoded, decodeErr := base64.StdEncoding.DecodeString(v); decodeErr == nil && len(decoded) > 0 {
			pemBytes = decoded
		}
	case *string:
		if v == nil {
			err = fmt.Errorf("pemData pointer is nil")
			return
		}
		pemBytes = []byte(*v)
		// Try base64 decode
		if decoded, decodeErr := base64.StdEncoding.DecodeString(*v); decodeErr == nil && len(decoded) > 0 {
			pemBytes = decoded
		}
	case []byte:
		pemBytes = v
		// Try base64 decode
		if decoded, decodeErr := base64.StdEncoding.DecodeString(string(v)); decodeErr == nil && len(decoded) > 0 {
			pemBytes = decoded
		}
	default:
		err = fmt.Errorf("invalid pemData type: expected string, *string, or []byte, got %T", pemData)
		return
	}

	var certificates []string
	var encryptedKeyBlock *pem.Block

	// Parse all PEM blocks
	remaining := pemBytes
	for {
		var block *pem.Block
		block, remaining = pem.Decode(remaining)
		if block == nil {
			break
		}

		switch block.Type {
		case "CERTIFICATE":
			certPEM := string(pem.EncodeToMemory(block))
			certificates = append(certificates, certPEM)
		case "ENCRYPTED PRIVATE KEY":
			encryptedKeyBlock = block
		case "RSA PRIVATE KEY", "EC PRIVATE KEY", "PRIVATE KEY":
			// Already unencrypted
			privateKey = string(pem.EncodeToMemory(block))
		}
	}

	// If we found an encrypted private key, decrypt it
	if encryptedKeyBlock != nil && privateKey == "" {
		decryptedKey, decryptErr := DecryptPKCS8PrivateKey(encryptedKeyBlock.Bytes, password)
		if decryptErr != nil {
			err = fmt.Errorf("failed to decrypt private key: %v", decryptErr)
			return
		}
		privateKey = decryptedKey
	}

	// Assign certificates: first is leaf, rest are CA chain
	if len(certificates) > 0 {
		certificate = certificates[0]
		if len(certificates) > 1 {
			caCertificates = certificates[1:]
		}
	}

	return privateKey, certificate, caCertificates, nil
}

// DecryptPKCS8PrivateKey decrypts a PKCS#8 encrypted private key and returns it in PEM format.
// Uses the github.com/youmark/pkcs8 package which supports PBES2 encryption schemes
// including AES-128-CBC, AES-192-CBC, AES-256-CBC, AES-128-GCM, AES-192-GCM, AES-256-GCM.
func DecryptPKCS8PrivateKey(encryptedKey []byte, password string) (string, error) {
	// Use the pkcs8 package to parse and decrypt the encrypted PKCS#8 key
	// This handles PBES2 encryption with various algorithms
	parsedKey, err := pkcs8.ParsePKCS8PrivateKey(encryptedKey, []byte(password))
	if err != nil {
		return "", fmt.Errorf("failed to decrypt PKCS#8 key: %v", err)
	}

	// Encode the decrypted key back to PEM format
	pemBlock, encodeErr := EncodePrivateKey(parsedKey)
	if encodeErr != nil {
		return "", fmt.Errorf("failed to encode decrypted key: %v", encodeErr)
	}

	return string(pem.EncodeToMemory(pemBlock)), nil
}

// UnpackPkcs12 extracts the private key, certificate, and CA certificates from a PKCS#12/PFX file.
// Parameters:
//   - pfxData: The byte slice containing the PKCS#12/PFX file data.
//   - password: The password used for decrypting the PKCS#12/PFX file.
//
// Returns:
//   - privateKey: The private key extracted from the PFX file, in PEM format.
//   - certificate: The certificate extracted from the PFX file, in PEM format.
//   - caCertificates: A slice of CA certificates extracted from the PFX file, in PEM format (if any).
//   - err: An error that describes why the unpacking failed, if any.
func UnpackPkcs12(pfxData interface{}, password string) (
	privateKey, certificate string,
	caCertificates []string,
	err error,
) {
	// Convert pfxData to []byte, if necessary
	var pfxBytes []byte

	switch v := pfxData.(type) {
	case string:
		// attempt to base64 decode first
		pfxBytes = []byte(v) // Convert string to []byte
		decoded, decodeErr := base64.StdEncoding.DecodeString(v)
		if decodeErr == nil && len(decoded) > 0 {
			pfxBytes = decoded
			break
		}

	case *string:
		if v == nil {
			err = fmt.Errorf("pfxData pointer is nil")
			return
		}
		// attempt to base64 decode first
		pfxBytes = []byte(*v) // Convert *string to []byte
		decoded, decodeErr := base64.StdEncoding.DecodeString(*v)
		if decodeErr == nil && len(decoded) > 0 {
			pfxBytes = decoded
			break
		}
		break
	case []byte:
		pfxBytes = v
	default:
		err = fmt.Errorf(
			"invalid pfxData type: expected string or []byte, got %s (type %T)",
			reflect.ValueOf(pfxData),
			pfxData,
		)
		return
	}

	// Decode the PKCS#12 data
	parsedKey, parsedCert, parsedCAs, pkcs12Err := pkcs12.DecodeChain(pfxBytes, password)
	if pkcs12Err != nil {
		err = fmt.Errorf("failed to decode PKCS#12 data: %v", pkcs12Err)
		return
	}

	// PEM-encode the private key
	privateKeyBlock, keyErr := EncodePrivateKey(parsedKey)
	if keyErr != nil {
		err = fmt.Errorf("failed to encode private key: %v", keyErr)
		return
	}
	privateKey = string(pem.EncodeToMemory(privateKeyBlock))

	// PEM-encode the certificate
	certificateBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: parsedCert.Raw,
	}
	certificate = string(pem.EncodeToMemory(certificateBlock))

	// PEM-encode the CA certificates (if any)
	for _, caCert := range parsedCAs {
		caCertBlock := &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: caCert.Raw,
		}
		caCertificates = append(caCertificates, string(pem.EncodeToMemory(caCertBlock)))
	}

	return privateKey, certificate, caCertificates, nil
}

// EncodePrivateKey determines the type of private key (RSA or ECDSA) and encodes it as a PEM block.
// Parameters:
//   - key: The private key to encode.
//
// Returns:
//   - pemBlock: The PEM block representation of the private key.
//   - err: An error if the private key type is unsupported or invalid.
func EncodePrivateKey(key interface{}) (*pem.Block, error) {
	switch k := key.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(k),
		}, nil
	case *ecdsa.PrivateKey:
		encodedKey, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return nil, fmt.Errorf("failed to encode ECDSA private key: %v", err)
		}
		return &pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: encodedKey,
		}, nil
	case ed25519.PrivateKey:
		return &pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: k,
		}, nil
	case *pkcs12.OpaquePrivateKey:
		// Algorithm not supported by Go's x509 (e.g. Ed448, OID 1.3.101.113).
		// The DER is already valid PKCS#8; wrap it directly.
		return &pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: k.DER,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported private key type: %T", key)
	}
}

// GetCertificateThumbprint computes the thumbprint (SHA-1 hash) of an x509 certificate.
//
// The thumbprint is calculated by hashing the raw DER-encoded certificate data
// using the SHA-1 algorithm.
//
// Parameters:
//   - cert: A pointer to an x509.Certificate object.
//
// Returns:
//   - A string representing the hexadecimal-encoded thumbprint of the certificate.
//   - An error, which will be nil if the computation succeeds.
//
// Example:
//
//	thumbprint, err := GetCertificateThumbprint(cert)
//	if err != nil {
//	    log.Fatalf("error computing thumbprint: %v", err)
//	}
//	fmt.Println("Certificate Thumbprint:", thumbprint)
func GetCertificateThumbprint(cert *x509.Certificate) (string, error) {
	// Compute the SHA-1 hash of the certificate's raw DER data
	hash := sha1.Sum(cert.Raw)

	// Convert the hash to a hexadecimal string
	thumbprint := hex.EncodeToString(hash[:])

	return thumbprint, nil
}
