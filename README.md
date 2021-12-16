## keyfactor-go-client
The Keyfactor Go Client is a Go module that provides abstracted
methods of interacting with the Keyfactor API. The primary supported 
functionality is the creation and modification of SSL/TLS certificates,
updating certificate metadata, and creating certificate stores.

### Usage
For now, this module's primary use is as an imported Go module for Go 
projects that need to use Keyfactor features. For example, if a plugin for 
a 3rd party tool such as HashiCorp Terraform is in development, this Go
module can be imported and used as a 'client library.' Using the module is
the same as including any other Go module.

To include this library in a Go project, include 
```github.com/Keyfactor/keyfactor-go-client```, add at least one invocation
to a function/structure used by the module, then run ```go mod tidy``` and
```go mod vendor``` to download the package.

#### Example
Let's suppose that a Go project needs to download a certificate from Keyfactor.
The Keyfactor Go utilities module is based around an authentication structure
called ```APIClient```. This structure expects a ```Hostname```, ```Username```,
and ```Password``` at the very least. This structure is used by all subsequent
utility functions. An instance of this struct must be created before using
any module functions.
```go
// Step 1: Create authentication structure
clientConfig := &keyfactor.APIClient{
    Hostname: HOSTNAME,
    Username: USERNAME,
    Password: PASSWORD,
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

Finally, call the appropriate ```keyfactor``` utilities function for the required
task. In this case, the ```DownloadCertificate``` method is used. As of the time of writing,
all functions expect (unless otherwise stated) require two arguments as pointers to an ```APIClient```
structure and an arguments structure. 
```go
// Step 3: Call associated function
response, err := keyfactor.DownloadCertificate(getCertContextArgs, clientConfig)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%#v", response)
```

### Installation
As of now, this module is designed as an imported package, and therefore doesn't
have a module directly designed for installation. Future functionality could be
a CLI that uses the module.