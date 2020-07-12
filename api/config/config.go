package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var Cfg Config

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Connection string `yaml:"connection"`
	} `yaml:"database"`
}

func ReadInConfig() {
	log.Println("Read in Config..")

	f, err := os.Open("config.yml")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		log.Println(err)
	}
}
