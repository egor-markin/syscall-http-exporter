package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

// Config represents the overall configuration structure
type Config struct {
	Address   string           `json:"address"`
	Endpoints []EndpointConfig `json:"endpoints"`
}

// EndpointConfig represents configuration for each endpoint
type EndpointConfig struct {
	Command     string `json:"command"`
	Endpoint    string `json:"endpoint"`
	ContentType string `json:"content_type"`
}

func main() {
	// Define command line parameter for config file
	configFile := flag.String("config", "config.json", "Path to the configuration JSON file")
	flag.Parse()

	// Read configuration from JSON file
	config, err := readConfig(*configFile)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}

	// Register HTTP handlers for each endpoint configuration
	for _, endpoint := range config.Endpoints {
		endpoint := endpoint // capture range variable
		http.HandleFunc(endpoint.Endpoint, func(w http.ResponseWriter, r *http.Request) {
			// Split the command into name and arguments
			parts := strings.Fields(endpoint.Command)
			cmdName := parts[0]
			cmdArgs := parts[1:]

			cmd := exec.Command(cmdName, cmdArgs...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", endpoint.ContentType)
			if _, err := w.Write(output); err != nil {
				return
			}
		})

		fmt.Printf("Endpoint '%s' registered at http://%s%s\n", endpoint.Endpoint, config.Address, endpoint.Endpoint)
	}

	// Start the HTTP server
	fmt.Println("Starting syscall-http-exporter server...")
	if err := http.ListenAndServe(config.Address, nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}

// readConfig reads configuration from a JSON file
func readConfig(filename string) (Config, error) {
	var config Config

	// Read JSON file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	// Unmarshal JSON into Config struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
