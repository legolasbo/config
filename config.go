package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func getDirectory(appName string) string {
	basePaths := [...]string{"/etc/", "~/.", "./"}

	for _, path := range basePaths {
		checkPath := path + appName + "/"
		if _, err := os.Stat(checkPath); os.IsNotExist(err) {
			continue
		}
		return checkPath
	}

	return "./"
}

// ReadJSON reads the json configuration file with the given file name for the given app name.
func ReadJSON(appName, name string, configuration interface{}) {
	configDir := getDirectory(appName)
	file, err := os.Open(configDir + name + ".json")
	if err != nil {
		file, err = os.Open(configDir + name + ".dist.json")
		if err != nil {
			log.Println(err)
			log.Println("Config file could not be opened. Using default configuration.")
			return
		}
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Could not read configuration")
		panic(err)
	}
}
