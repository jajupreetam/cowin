# Digital Commerce Platform Service

##Service with Golang (1.15.x)
* LineCheck integration
* Hybris integration
* Configuration profiles

## How do I get set up?

* Download and setup the [`Go`](https://golang.org/) - v1.15.x
  
* Folder Structure
  1. Create `workspace` folder
  2. Create `src` folder inside `workspace` folder
  3. Clone this Repository inside `workspace/src` 
  4. Import the Project into Intellij IDE
    
* Configuration     
  Setup System Environment variables          
  a. **GOROOT**  - Path to `Go Installation` Directory   
  b. **GOPATH**  - Path to `workspace` directory 

* Dependencies  
  Run `go get` from `workspace/src/cowin-service` to download the dependencies of project
  
* Go modules
  ``go mod init`` creates a new module, initializing the go.mod file that describes it.<br/>
  ``go build, go test,`` and other package-building commands add new dependencies to go.mod as needed.<br/>
  ``go list -m all`` prints the current module’s dependencies.<br/>
  ``go get`` changes the required version of a dependency (or adds a new dependency).<br/>
  ``go mod tidy`` removes unused dependencies.<br/>

## Build
``go build -o app``

After a successful build, a native binary **_app_** executable should appear in the root directory.

## Run
``./app``  

``PROFILE=t01,l2stdout ./app``

The active profile could be specified either as an environment PROFILE variable or by defining it into the ./config/config.go file. The ENV variable overrides the default ./config/config.go definition.

## Watch
``go get github.com/githubnemo/CompileDaemon``  

The CompileDaemon watches the .go files and invokes a go build if a file changed. https://github.com/githubnemo/CompileDaemon

``CompileDaemon -build="go build -o app" -command="./app"``<br/>
``PROFILE=dev,l2stdout,t02 CompileDaemon -build="go build -o app" -command="./app"``

The active profile could be specified either as an environment PROFILE variable or by defining it into the ./config/config.go file. The ENV variable overrides the default ./config/config.go definition.

## Http Client Proxy
By default (local development), the Http Client Proxy is enabled with the url value ``const HttpClientProxyUrl = "http://localhost:3128"``

The Http Client Proxy is disabled in the OpenShift Pods ``HTTP_CLIENT_PROXY_ENABLE=false PROFILE=t01,l2stdout ./app``

## API
### Postman collection
[TODO]

## Resources
mux: A powerful HTTP router and URL matcher for building Go web servers. https://github.com/gorilla/mux

viper: Viper is a complete configuration solution for Go applications including 12-Factor apps. https://github.com/spf13/viper

CompileDaemon: Watches your .go files in a directory and invokes go build if a file changed. https://github.com/githubnemo/CompileDaemon

## Generate Typescript resources
go run cmd/generate-client-model.go 
tscriptify -package="cowin-service/model" -target="resources/client/model.ts" Product Cart CartModificationList LineCheckEligibility Account

## Format Go files with Gofmt
gofmt -s -w .

