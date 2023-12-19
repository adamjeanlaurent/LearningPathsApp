package utility

import "os"

type Configuration interface {
	GetEnvironment() string
	GetJwtSecretKey() []byte
}

type ServerConfiguration struct {
	environment  string
	jwtSecretKey string
}

func NewServerConfiguration() *ServerConfiguration {
	return &ServerConfiguration{
		environment:  os.Getenv("MYAPP_ENV"),
		jwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}

func (config *ServerConfiguration) GetEnvironment() string {
	return config.environment
}

func (config *ServerConfiguration) GetJwtSecretKey() []byte {
	return []byte(config.jwtSecretKey)
}
