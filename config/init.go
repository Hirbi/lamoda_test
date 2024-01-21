package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Server struct {
		Port string `yaml:port`
	} `yaml: server`
	Database struct {
		Driver   string `yaml:driver`
		User     string `yaml:user`
		Password string `yaml:password`
		Name     string `yaml:name`
		Host     string `yaml:host`
		Port     string `yaml:port`
	} `yaml: database`
}

var PORT = 8088

var Config config

func InitConifg() error {
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return err
	}
	yaml.Unmarshal(data, &Config)
	fmt.Printf("%+v\n", Config)
	return nil
}
