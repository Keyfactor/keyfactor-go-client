package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var showVersion bool

	version := os.Getenv("GITHUB_REF_NAME")
	if version == "" {
		version = "1.0.0"
	}
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.Parse()

	// Show version and exit
	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}
}
