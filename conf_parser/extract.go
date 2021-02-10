package conf_parser

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	DirsToScan         []string `yaml:"dirs_to_scan"`
	Exts               []string `yaml:"extensions_to_handle"`
	TargetFolderToCopy string   `yaml:"target_folder_to_copy"`
}

func (c *Config) Extensions() map[string]struct{} {
	exts := make(map[string]struct{})
	for _, ext := range c.Exts {
		exts[ext] = struct{}{}
	}
	return exts
}

func Extract(confPath string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(confPath)
	conf := &Config{}
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
