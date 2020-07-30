package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gronka/tg"
	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	InitialPath string `yaml:"initialPath"`
}

var Conf ConfigStruct

func GenerateConfig(initialPath string) ConfigStruct {
	conf := ConfigStruct{}

	home := os.Getenv("HOME")
	configPath := filepath.Join(home, ".config", "butter.yaml")
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		tg.Warn("Failed to open config file from: " + configPath)
	} else {
		err = yaml.Unmarshal(configFile, conf)
		if err != nil {
			tg.Warn("Failed to parse config file from: " + configPath)
		}
	}

	if conf.InitialPath == "CURRENT" {
		conf.InitialPath, err = os.Getwd()
		if err != nil {
			tg.Warn("Failed to get working directory")
		}
	}
	if initialPath != "" {
		conf.InitialPath = initialPath
	}
	if conf.InitialPath == "" {
		conf.InitialPath = home
	}

	return conf
}
