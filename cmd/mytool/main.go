package main

import (
	"commandTest/internal/cli"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("mytool: ")
	log.SetOutput(os.Stderr)
}

func main() {
	rootCmd := cli.NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
