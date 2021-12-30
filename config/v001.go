package config

import "gopkg.in/yaml.v2"

// ConfigV001 describes the V 0.0.1 configuration yaml
type ConfigV001 struct {
	Routes []ConfigV001Route `yaml:"routes"`
}

// ConfigV001Routes describes the V 0.0.1 configuration routes
type ConfigV001Route struct {
	Path     	 string   `yaml:"path"`
	Params   	 []string `yaml:"params"`
	Size     	 string   `yaml:"size"`
	Template 	 string   `yaml:"template"`
	CacheControl string   `yaml:"cache-control"`
}

// ReadV001Config returns a parsed V 0.0.1 configuration struct
func ReadV001Config(config []byte) ConfigV001 {
	var configuration ConfigV001
	err := yaml.Unmarshal(config, &configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
