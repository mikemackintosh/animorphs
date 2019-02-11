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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = text + scanner.Text()
	}

	if scanner.Err() != nil {
    		fmt.Println("Unable to read input")
		os.Exit(1)
	}

	var object interface{}
	_ = json.Unmarshal([]byte(text), &object)

	o, _ := yaml.Marshal(&object)
	fmt.Println(string(o))
}
