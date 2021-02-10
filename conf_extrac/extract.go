package conf_extrac

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Data struct {
	DirsToScan []string `yaml:"dirs_to_scan"`
	Exts       []string `yaml:"extensions_to_handle"`
}

func (c *Data) Extensions() map[string]struct{} {
	exts := make(map[string]struct{})
	for _, ext := range c.Exts {
		exts[ext] = struct{}{}
	}
	return exts
}

func Extract(confPath string) (*Data, error) {
	yamlFile, err := ioutil.ReadFile(confPath)
	conf := &Data{}
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
