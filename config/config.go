package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Config ConfigV001
var ConfigError error

// GauguinConfigVersion is a struct containing just the version number of the Gauguin configuration
type GauguinConfigVersion struct {
	Version string `yaml:"version"`
}

func init() {
	Config, ConfigError = ReadConfigFile()
}

// ReadConfigFile returns a parsed configuration file
func ReadConfigFile() (ConfigV001, error) {
	content, err := ioutil.ReadFile("gauguin.yaml")
	if err != nil {
		return ConfigV001{}, err
	}

	configVersion := getConfigVersion(content)

	// We want to support different versions in the future
	switch configVersion.Version {
	case "0.0.1":
		return ReadV001Config(content), nil
	default:
		fmt.Printf("%s is not a valid version number", configVersion.Version)
		os.Exit(1)
	}

	return ConfigV001{}, nil
}

func getConfigVersion(config []byte) GauguinConfigVersion {
	var version GauguinConfigVersion

	err := yaml.Unmarshal([]byte(config), &version)
	if err != nil {
		panic(err)
	}

	return version
}
