package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"reflect"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ListenAddress string `yaml:"ListenAddress"`
	AccessToken string `yaml:"AccessToken"`
	OwncastAddress string `yaml:"OwncastAddress"`
}

func DefaultConfiguration() Config {
	var defaultConfig = Config { ":8100", "SECRETACCESSTOKEN", "http://localhost:8000" }
	return defaultConfig
}

func CollectDemoBotConfiguration(fileName string) Config {
	_, err := os.Stat(fileName);
	if err != nil {
		log.Print(fmt.Sprintf("Could not find file: %s. Reason: '%s'. Using defaults instead.", fileName, err))
		return DefaultConfiguration()
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Print(fmt.Sprintf("Could not read file: %s. Reason: '%s'. Using defaults instead.", fileName, err))
		return DefaultConfiguration()
	}

	var cfg = DefaultConfiguration()
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < reflect.ValueOf(cfg).NumField(); i++ {
		structFieldName := reflect.ValueOf(&cfg).Elem().Type().Field(i).Name
		environmentVariableName := fmt.Sprintf("%s_%s", "OWNCAST_DEMOBOT", strings.ToUpper(structFieldName))
		environmentVariableValue := os.Getenv(environmentVariableName)

		if environmentVariableValue != "" {
			structField := reflect.ValueOf(&cfg).Elem().FieldByName(structFieldName)
			structField.SetString(environmentVariableValue)
		}
	}
	
	return cfg
}
