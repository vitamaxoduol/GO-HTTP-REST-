package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type CommandRequest struct {
	Command string   `json:"command"`
	Args    []string `json:"args,omitempty"`
}

type CommandResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

// function to execute commands
func executeCommand(cmd string, args []string) (string, error) {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	return string(out), err
}

// handler function for the API
func commandHandler(w http.ResponseWriter, r *http.Request) {
	var cmdReq CommandRequest

	// Attempt to decode the JSON body if present
	if r.Body != nil {
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&cmdReq)
		if err != nil && err.Error() != "EOF" { // Ignore EOF error for empty body
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
	}

	// Log request details for debugging
	fmt.Println("Received request:", r.Method, r.URL)

	// Determine the command and args from JSON body or query parameter
	var cmd string
	var args []string

	if cmdReq.Command != "" {
		cmd = cmdReq.Command
		args = cmdReq.Args
	} else {
		cmd = r.URL.Query().Get("command")
		if cmd != "" {
			parts := strings.Fields(cmd) // Split the command and arguments
			cmd = parts[0]               // Command is the first part
			if len(parts) > 1 {
				args = parts[1:] // Remaining parts are arguments
			}
		}
	}

	// If no command was provided, return an error
	if cmd == "" {
		http.Error(w, "Command is required", http.StatusBadRequest)
		return
	}

	output, err := executeCommand(cmd, args)
	response := CommandResponse{Output: output}

	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} // Handle the error gracefully

	w.Header().Set("Content-Type", "application/json") // Set other headers as needed
	json.NewEncoder(w).Encode(response)
}

// main function to start the server
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/cmd", commandHandler).Methods("POST")

	// Apply the CORS middleware to our router
	handler := cors.Default().Handler(r)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", handler)
}
