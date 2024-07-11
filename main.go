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
	Command string `json:"command"`
}

type CommandResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func executeCommand(cmd string) (string, error) {
	parts := strings.Fields(cmd)
	name := parts[0]
	args := parts[1:]

	out, err := exec.Command(name, args...).CombinedOutput()
	return string(out), err
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	var cmdReq CommandRequest

	if r.Body != nil {
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&cmdReq)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
	}

	if cmdReq.Command == "" {
		cmdReq.Command = r.URL.Query().Get("command")
	}

	if cmdReq.Command == "" {
		http.Error(w, "Command is required", http.StatusBadRequest)
		return
	}

	output, err := executeCommand(cmdReq.Command)
	response := CommandResponse{Output: output}

	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/cmd", commandHandler).Methods("POST")

	// Applying the CORS middleware to our router
	handler := cors.Default().Handler(r)
	// http.Handle("/", r)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", handler)
}
