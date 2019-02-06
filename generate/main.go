package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	var config string
	flag.StringVar(&config, "config", "settings.yaml", "settings")

	flag.Parse()

	b, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatalf("unable to read file %s: %v", config, err)
	}

	var c Console
	if err := yaml.Unmarshal(b, &c); err != nil {
		log.Fatalf("unable to decode file %s: %v", config, err)
	}

	if _, err := fmt.Fprintf(os.Stdout, prefix); err != nil {
		log.Fatal(err)
	}

	if err := c.Generate(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
