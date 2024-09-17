# Keyfactor Command Golang Client (DEPRECATED)

The Keyfactor Command Golang Client is a Go module that provides abstracted methods of interacting with the Keyfactor API. The primary supported functionality is the creation and modification of SSL/TLS certificates, updating certificate metadata, and creating certificate stores.

## Integration status: Deprecated
This client is deprecated as of `2024-09` and will no longer be maintained. Please use the [Keyfactor Command Go Client SDK](https://github.com/Keyfactor/keyfactor-go-client-sdk) for future integrations.

## About the Keyfactor API Client

This API client allows for programmatic management of Keyfactor resources.

## Support for Keyfactor Command Golang Client

Keyfactor Command Golang Client is open source and there is **no SLA** for this tool/library/client. Keyfactor will address issues as resources become available. Keyfactor customers may request escalation by opening up a support ticket through their Keyfactor representative.

###### To report a problem or suggest a new feature, use the **[Issues](../../issues)** tab. If you want to contribute actual bug fixes or proposed enhancements, use the **[Pull requests](../../pulls)** tab.
___




# Keyfactor Command Go Client
The client is a Golang module for the Keyfactor Command Web API

### Usage
This module's primary use is as a client to connect to Keyfactor Command using the Command API. 

To use this module, include
```github.com/Keyfactor/keyfactor-go-client/pkg/keyfactor```, add at least 
one invocation to a function/structure used by the module, then run 
```go mod tidy``` and ```go mod vendor``` to download the package.

#### Example
Let's suppose that a Go project needs to download a certificate from Keyfactor.
The Keyfactor Command Go Client must be initialized by creating and populating an
`AuthConfig` structure, and passing it to `NewKeyfactorClient`. This
function ensures that the provided configuration data is valid by making a test
call to Keyfactor. If no error occurs, a `Client` structure is returned with
associated methods for interacting with Keyfactor Command.
```go
// Step 1: Create authentication structure
clientConfig := &keyfactor.AuthConfig{
    Hostname: HOSTNAME,
    Username: USERNAME,
    Password: PASSWORD,
}
client, err := keyfactor.NewKeyfactorClient(clientConfig)
if err != nil {
    log.Fatal(err)
}
```

Next, an `arguments` structure for the specific API method must be created
containing (at least) the required arguments for the call to Keyfactor Command. In this
example, a certificate is being downloaded. In this case, we're assuming that
the Keyfactor Command certificate ID is known. Other options can be specified, such as
the option to include metadata or deployed locations data. 
```go
// Step 2: Create arguments structure
downloadArgs := &keyfactor.DownloadCertArgs{
    CertID:       1860,
    IncludeChain: true,
    CertFormat:   "PEM",
}
```

Finally, call the appropriate `keyfactor` method for the required
task. In this case, the `DownloadCertificate` method is used. All functions expect (unless otherwise stated) require two 
arguments as pointers to an `APIClient`
structure and an arguments structure. 
```go
// Step 3: Call associated function
response, err := client.DownloadCertificate(downloadArgs)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%#v", response)
```

### Installation
As of now, this module is designed as an imported package, and therefore doesn't
have a module directly designed for installation. Future functionality could be
a CLI that uses the module.

### Supported Methods
* ```EnrollPFX```
* ```DownloadCertificate```
* ```EnrollCSR```
* ```UpdateMetadata```
* ```RevokeCert```
* ```DeployPFXCertificate```
* ```GetCertificateContext```
* ```RecoverCertificate```
* ```CreateStore```
* ```GetCertStoreType```
* ```DeleteCertificateStore```
* ```GetCertificateStoreByID```
* ```AddCertificateToStores```
* ```RemoveCertificateFromStores```
* ```GetSecurityIdentities```
* ```CreateSecurityIdentity```
* ```DeleteSecurityIdentity```
* ```GetSecurityRole```
* ```DeleteSecurityRole```
* ```CreateSecurityRole```
* ```UpdateSecurityRole```
* ```GetTemplate```
* ```UpdateTemplate```

