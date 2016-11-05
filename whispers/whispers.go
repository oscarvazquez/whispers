package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var environment, filePath string
var create bool

var data string = "# set your environment variables here\n # Make sure not to erase production, development or testing line. This is where you set the environment\n # variables for each type.\n # any variables included in default will be installed for all of your environments\n\ndefault:\n  DEFAULT_NAME: 'default name'\n\ndevelopment:\n #example\n  NAME: 'Oscar Vazquez'\n\nproduction:\n  NAME: 'Secret Name'\n"

func main() {
	parseFlags()
}

func parseFlags() {
	flag.BoolVar(&create, "create", false, "create secrets path and add to gitignore")
	flag.Parse()
	if create {
		createSecrets()
	}
}

func createSecrets() {
	// TODO: create secrets file and add to gitignore and return
	ioutil.WriteFile("./envs.yml", []byte(data), 0644)
	fmt.Println("creating secrets file updating gitignore")
}
