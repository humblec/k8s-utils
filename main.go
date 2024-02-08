package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Dependency struct {
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	RefPaths  []struct {
		Path  string `yaml:"path"`
		Match string `yaml:"match"`
	} `yaml:"refPaths"`
}

type Dependencies struct {
	Dependencies []Dependency `yaml:"dependencies"`
}

func main() {
	// Define flag for input file path
	var inputFile string
	flag.StringVar(&inputFile, "input", "dependencies.yaml", "Input YAML file path")

	// Define custom help message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Read YAML file
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Unmarshal YAML into struct
	var dependencies Dependencies
	err = yaml.Unmarshal(data, &dependencies)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	// Iterate over dependencies and print name and version
	for _, dep := range dependencies.Dependencies {
		fmt.Printf("Name: %s\n", dep.Name)
		fmt.Printf("Version: %s\n", dep.Version)
		fmt.Println("----------")
	}
}

