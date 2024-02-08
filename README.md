# Dependencies Parser

Dependencies Parser is a Go program designed to parse YAML files containing dependency information for Kubernetes projects. It can read from a local file or fetch the YAML file directly from GitHub based on the specified Kubernetes version.

## Features

- Parse dependency information from a local YAML file or fetch from GitHub.
- Supports specifying the Kubernetes version to fetch the appropriate YAML file.
- Print the name and version of each component specified in the YAML file.

## Usage

### Installation

1. Make sure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).

2. Clone the repository:

```
git clone https://github.com/humblec/k8s-utils/dependencies-parser.git

```
# Dependencies Parser

Dependencies Parser is a Go program designed to parse YAML files containing dependency information for Kubernetes projects. It can read from a local file or fetch the YAML file directly from GitHub based on the specified Kubernetes version.


## Features

- Parse dependency information from a local YAML file or fetch from GitHub.
- Supports specifying the Kubernetes version to fetch the appropriate YAML file.
- Print the name and version of each component specified in the YAML file.

## Usage

### Installation

1. Make sure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).

2. Clone the repository:

git clone https://github.com/humblec/k8s-utils/dependencies-parser.git

css


3. Navigate to the project directory:

cd dependencies-parser

markdown


4. Build the program:

go build

bash


### Run

Run the program using the following command:

./dependencies-parser [options]

csharp


Replace `[options]` with the desired command-line options.

### Options

- `-input`: Specifies the path to the local YAML file containing dependency information. If not provided, the program fetches the YAML file from GitHub based on the Kubernetes version.
- `-k8s-ver`: Specifies the Kubernetes version. If provided, the program fetches the YAML file from GitHub based on this version.

### Examples

1. Parse dependency information from a local YAML file:

./dependencies-parser -input dependencies.yaml

sql


2. Fetch dependency information from GitHub for Kubernetes version 1.28:

./dependencies-parser -k8s-ver 1.28

sql


3. Fetch dependency information from GitHub for Kubernetes version v1.29:

./dependencies-parser -k8s-ver v1.29

csharp


## License

