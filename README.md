# GO-HTTP-REST-API
- A simple GO HTTP REST API that will accept a shell command and return the output of that command.

## Task
- Write a simple GO HTTP REST API that will accept a shell command and
return the output of that command.
- Create an endpoint with the POST method. api/cmd POST
- Accept the command via query param or JSON body.
Return the output of the command as a response.

**Plus point:**
- Return error if the command is not found, with proper status code.

## Requirements
- Downloading a binary release suitable for your system and then follow the installation instructions in GO documentation 
- In the task directory, intialize the module: `go mod init GO-HTTP-REST-API`
- With the module initialized, install the Gorilla Mux package `go get github.com/gorilla/mux` 
- Create a new file called `main.go` and add the following code
- Run the Go server: `go run main.go`
- You can test the API using curl:
### Testing the API POST REQUEST
- Using JSON body:
`curl -v -X POST -H "Content-Type: application/json" -d '{"command":"ls","args":["-la"]}' http://localhost:8080/api/cmd`

- Using query parameter:
`curl -v -X POST "http://localhost:8080/api/cmd?command=ls%20-la"`


