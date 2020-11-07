package config

import "gopkg.in/yaml.v2"

// ConfigV001 describes the V 0.0.1 configuration yaml
type ConfigV001 struct {
	Routes []ConfigV001Route `yaml:"routes" json:"routes"`
}

// ConfigV001Routes describes the V 0.0.1 configuration routes
type ConfigV001Route struct {
	Path     string   `yaml:"path" json:"path"`
	Params   []string `yaml:"params" json:"params"`
	Size     string   `yaml:"size" json:"size"`
	Template string   `yaml:"template" json:"template"`
	Name     string   `yaml:"name" json:"name"`
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
