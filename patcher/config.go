package patcher

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Url UrlConfig `yaml:"url"`
}

type UrlConfig struct {
	PatchInfo string `yaml:"PATCH_INFO"`
}

func LoadConfig(path string, config *Config) error {
	bytes, err := read(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bytes, config); err != nil {
		return err
	}
	return nil
}

func read(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
