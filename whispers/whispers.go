package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var filePath string
var create bool
var git bool

var data string = "# set your environment variables here\n # Make sure not to erase production, development or testing line. This is where you set the environment\n # variables for each type.\n # any variables included in default will be installed for all of your environments\n\ndefault:\n  DEFAULT_NAME: 'default name'\n\ndevelopment:\n #example\n  NAME: 'Oscar Vazquez'\n\nproduction:\n  NAME: 'Secret Name'\n"

func main() {
	parseFlags()
}

func parseFlags() {
	flag.BoolVar(&create, "create", false, "create environment yaml file")
	flag.BoolVar(&git, "ignore", false, "Add path to gitignore file")
	flag.StringVar(&filePath, "path", "./envs.yml", "Set the file path for envs yml file")
	flag.Parse()
	if create {
		createSecrets()
	}
	if git {
		addToGit()
	}
}

func createSecrets() {
	// TODO: create secrets file and add to gitignore and return
	ioutil.WriteFile(filePath, []byte(data), 0644)
	fmt.Println("Creating secrets file")
}

func addToGit() {
	fmt.Println("Updating gitignore file")
	f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("\n" + filePath); err != nil {
		panic(err)
	}
}
