package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	log "github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

// Config object to be loaded with configuration from YAML file
type Config struct {
	Project              string   `yaml:"project,omitempty"`
	Datasets             []string `yaml:"datasets,omitempty"`
	MaxClientConcurrency int      `yaml:"max_client_concurrency,omitempty"`
	LogLevel             string   `yaml:"log_level,omitempty"`
	DBHostname           string   `yaml:"db_hostname,omitempty"`
	DBPort               int32    `yaml:"db_port,omitempty"`
	DBReset              bool     `yaml:"db_reset,omitempty"`
	ServerPort           int32    `yaml:"server_port,omitempty"`
	WebDirectory         string   `yaml:"web_directory,omitempty"`
	CronRefresh          string   `yaml:"cron_refresh,omitempty"`
	Logger               *log.Logger
}

// LoadFromFile return Config object according to a YAML file
func (c *Config) LoadFromFile(file string) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, c)
	}
	return err
}

// PrintConfig print the config
func (c *Config) PrintConfig() {
	currentUser, err := user.Current()
	if err == nil {
		c.Logger.Debugf(fmt.Sprintf("Author: %s", currentUser.Username))
	}
	c.Logger.Debugf(fmt.Sprintf("Project: %s", c.Project))
	credEnv, ok := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !ok {
		credEnv = "default"
	}
	c.Logger.Debugf(fmt.Sprintf("Credentials: %s", credEnv))
}
