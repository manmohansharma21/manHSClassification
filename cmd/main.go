// go run cmd/main.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// List of queries to classify
	queries := []string{
		"I want a laptop charger",
		"laptop cover",
		"laptop toy",
		"laptop machine",
		"laptop stand",
		"laptop sleeve",
		"laptop charger",
	}

	// Join queries into a single comma-separated string to pass as an argument to the Python script
	queriesString := strings.Join(queries, ",")

	// Run the Python script to process the queries
	//cmd := exec.Command("python3", "scripts/process_query_embedding.py", queriesString) //--> Emebddinga generation using Transformer Python library
	cmd := exec.Command("python3", "scripts/generate_embeddings_openai.py", queriesString) //--> Using openapi for embeddings generation
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and check for errors
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running Python script: %v\n", err)
	}
}
