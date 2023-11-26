package configuration

import "os"

type AppConfig struct {
	environment string
}

func (config *AppConfig) GetEnvironment() string {
	return config.environment
}

var Config = &AppConfig{environment: os.Getenv("MYAPP_ENV")}
