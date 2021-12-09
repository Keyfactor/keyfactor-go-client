## keyfactor-go-client
The Keyfactor Go Client is a Go library that provides abstracted
methods of interacting with the Keyfactor API. The primary supported 
functionality is the creation and modification of SSL/TLS certificates,
updating certificate metadata, and creating certificate stores. 

### Installation
1. Initialize new Go module and set the current directory
to the root of the module

```go mod init keyfactor-go-client```
2. Create a ```vendor``` directory and populate it with 
dependencies

```go mod vendor```
3. Build the module

```make build```
