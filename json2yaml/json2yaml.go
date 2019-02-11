package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	var text string
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = text + scanner.Text()
	}

	if scanner.Err() != nil {
		fmt.Println("Unable to read input")
		os.Exit(1)
	}

	var object map[string]interface{}
	err = json.Unmarshal([]byte(text), &object)
	if err != nil {
		fmt.Printf("This isn't valid JSON\n")
		os.Exit(1)
	}

	o, err := yaml.Marshal(&object)
	if err != nil {
		fmt.Printf("Error converting: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(string(o))
}
