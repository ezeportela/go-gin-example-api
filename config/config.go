package config

import (
	"os"
	"strings"

	"github.com/ezeportela/meli-challenge/models"
	"github.com/ezeportela/meli-challenge/shared"
	"gopkg.in/yaml.v2"
)

type Config struct {
	MongoURI     string        `json:"mongo_uri"`
	Version      string        `yaml:"version"`
	DatabaseName string        `yaml:"database_name"`
	BasePath     string        `yaml:"base_path"`
	Environments []string      `yaml:"environments"`
	Roles        []models.Role `yaml:"roles"`
}

func (c *Config) LoadConfig(path string) error {
	data, err := shared.ReadFile(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &c)

	return err
}

func (c *Config) LoadEnvironments() {
	envs := make(map[string]string)

	for _, e := range c.Environments {
		envs[strings.ToLower(e)] = os.Getenv(e)
	}

	shared.JsonToStruct(envs, &c)
}

func (c *Config) Setup(path string) error {
	err := c.LoadConfig(path)
	c.LoadEnvironments()
	return err
}
