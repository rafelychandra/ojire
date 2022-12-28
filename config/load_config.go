package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

type HandlerLoadConfig struct {
	Env string
}

func (hlc *HandlerLoadConfig) LoadConfig() (*Config, error) {
	_, filename, _, _ := runtime.Caller(1)
	pathENV := path.Join(path.Dir(filename), "../"+env[hlc.Env])
	_, err := os.Stat(pathENV)
	if err != nil {
		return nil, err
	}

	var c *Config
	yamlFile, err := ioutil.ReadFile(pathENV)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
