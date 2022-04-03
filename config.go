package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Conf struct {
	Restrictions struct {
		Group int64 `yaml:"group"`
	}
}

func GetConf() Conf {
	ymlFile, err := os.Open("conf.yml")
	if err != nil {
		log.Fatalf("Error reading config.yml: #%v ", err)
	}

	config := Conf{}
	d := yaml.NewDecoder(ymlFile)

	if err != d.Decode(&config) {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}
