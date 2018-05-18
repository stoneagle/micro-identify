package common

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type DBConf struct {
	Name     string
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Target   string
	MaxIdle  int
	MaxOpen  int
	ShowSQL  bool
	Location string
	Img      string
}

type Conf struct {
	App struct {
		Mode string
	}
	Card struct {
		Database DBConf
		Ipc      struct {
			Model string
			Img   string
		}
	}
}

var onceConfig *Conf = &Conf{}

func GetConfig() *Conf {
	if (Conf{}) == *onceConfig {
		configPath := os.Getenv("ConfigPath")
		if configPath == "" {
			configPath = "../config/.config.yaml"
		}
		yamlFile, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic(err)
		}
		config := &Conf{}
		err = yaml.Unmarshal(yamlFile, config)
		if err != nil {
			panic(err)
		}
		onceConfig = config
	}
	return onceConfig
}
