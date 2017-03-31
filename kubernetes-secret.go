package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var Resource struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Type       string `json:"type"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Data map[string]string `json:"data"`
}

var Env bool
var Delimiter string
var FileLinkIndicator string

func init() {
	Resource.APIVersion = "v1"
	Resource.Kind = "Secret"
	Resource.Type = "Opaque"
	Resource.Data = make(map[string]string)
	flag.StringVar(&Resource.Metadata.Name, "n", "secret", "name")
	flag.StringVar(&Resource.Metadata.Namespace, "ns", "default", "namespace")
	flag.BoolVar(&Env, "e", true, "delimited key/value pairs as input")
	flag.StringVar(&Delimiter, "d", "=", "delimiter (if -e is specified)")
	flag.StringVar(&FileLinkIndicator, "f", ">>>", "file link indicator (if -f is specified)")
}

func main() {
	flag.Parse()

	// Create a buffered reader from stdin.
	r := bufio.NewReader(os.Stdin)

	// Check if we should read this in as delimited values.
	if Env {
		for {
			// Get a value by reading bytes from stdin up to the next newline.
			// (Newline characters aren't allowed in secrets.)
			line, err := r.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}

			// Parse the components.
			components := strings.SplitN(line, Delimiter, 2)
			if len(components) != 2 {
				log.Fatalln("invalid input")
			}

			// Encode and add the value to the resource.
			varname := strings.TrimSpace(components[0])
			value := strings.TrimSpace(components[1])

			if strings.HasPrefix(value, FileLinkIndicator) {
				file := strings.TrimPrefix(value, FileLinkIndicator)
				bytes, err := ioutil.ReadFile(file)
				if err != nil {
					fmt.Print(err)
				}
				Resource.Data[varname] = base64.StdEncoding.EncodeToString(bytes)
			} else {
				Resource.Data[varname] = base64.StdEncoding.EncodeToString([]byte(strings.TrimSpace(components[1])))
			}
		}
	} else {
		// Use the command line arguments as keys.
		for _, arg := range flag.Args() {
			// Get a value by reading bytes from stdin up to the next newline.
			// (Newline characters aren't allowed in secrets.)
			value, err := r.ReadBytes('\n')
			if err != nil {
				log.Fatal(err)
			}

			// Encode and add the value to the resource.
			Resource.Data[arg] = base64.StdEncoding.EncodeToString(value)
		}
	}

	// Encode the resource as JSON and write it to stdout.
	b, err := json.Marshal(Resource)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(b)
}
