package whispers

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ParseConfigs(filePath string, environment string) {
	filename, _ := filepath.Abs(filePath)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	m := make(map[interface{}]map[string]string)
	err = yaml.Unmarshal([]byte(yamlFile), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for key, envs := range m {
		if key == environment || key == "default" {
			for x, j := range envs {
				os.Setenv(x, j)
			}
		}
		// TODO: Should I unset other environment variables?
	}
}
