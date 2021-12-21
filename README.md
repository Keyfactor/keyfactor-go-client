## keyfactor-go-client
The Keyfactor Go Client is a Go module that provides abstracted
methods of interacting with the Keyfactor API. The primary supported 
functionality is the creation and modification of SSL/TLS certificates,
updating certificate metadata, and creating certificate stores.

### Usage
This module's primary use is as a client to connect to Keyfactor using the
Keyfactor API. For example, if a plugin for
a 3rd party tool such as HashiCorp Terraform is in development, this Go
module can be imported and used to facilitate a connection to Keyfactor.

To use this module, include
```github.com/Keyfactor/keyfactor-go-client/pkg/keyfactor```, add at least 
one invocation to a function/structure used by the module, then run 
```go mod tidy``` and ```go mod vendor``` to download the package.

#### Example
Let's suppose that a Go project needs to download a certificate from Keyfactor.
The Keyfactor Go Client must be initialized by creating and populating an
```AuthConfig``` structure, and passing it to ```NewKeyfactorClient```. This
function ensures that the provided configuration data is valid by making a test
call to Keyfactor. If no error occurs, a ```Client``` structure is returned with
associated methods for interacting with Keyfactor.
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
* Note: As of the time of writing, a better alternative to ```Basic``` API authentication
  is known for interfacing with the Keyfactor API.

Next, an ```arguments``` structure for the specific API method must be created
containing (at least) the required arguments for the call to Keyfactor. In this
example, a certificate is being downloaded. In this case, we're assuming that
the Keyfactor certificate ID is known. Other options can be specified, such as
the option to include metadata or deployed locations data. 
```go
// Step 2: Create arguments structure
downloadArgs := &keyfactor.DownloadCertArgs{
    CertID:       1860,
    IncludeChain: true,
    CertFormat:   "PEM",
}
```

Finally, call the appropriate ```keyfactor``` method for the required
task. In this case, the ```DownloadCertificate``` method is used. As of the time of writing,
all functions expect (unless otherwise stated) require two arguments as pointers to an ```APIClient```
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
* ```GetTemplate```
* ```RecoverCertificate```
* ```CreateStore```
* ```GetCertStoreType```
* ```DeleteCertificateStore```
* ```GetCertificateStoreByID```