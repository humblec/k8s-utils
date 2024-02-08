/*
Copyright 2024 The K8s-utils Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/


package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
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
	// Define flags for input file path and Kubernetes version
	var inputFile string
	var k8sVersion string

	flag.StringVar(&inputFile, "input", "", "Input YAML file path")
	flag.StringVar(&k8sVersion, "k8s-ver", "", "Kubernetes version")

	// Define custom help message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if inputFile == "" && k8sVersion == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Check if input file path is provided
	if inputFile != "" {
		readFromLocalFile(inputFile)
	} else {
		// Construct URL based on Kubernetes version
		url := fmt.Sprintf("https://raw.githubusercontent.com/kubernetes/kubernetes/release-%s/build/dependencies.yaml", k8sVersion)
		readFromURL(url)
	}
}

func readFromLocalFile(inputFile string) {
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

func readFromURL(url string) {
	// Fetch YAML file from URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching YAML file from URL: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Unmarshal YAML into struct
	var dependencies Dependencies
	err = yaml.Unmarshal(body, &dependencies)
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

