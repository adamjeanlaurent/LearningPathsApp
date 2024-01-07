package utility

import (
	"fmt"
	"os"
	"reflect"
)

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

func (config *ServerConfiguration) Validate() error {
	// none of the values should be empty

	var values reflect.Value = reflect.ValueOf(config)

	for i := 0; i < values.NumField(); i++ {
		var field reflect.Value = values.Field(i)
		var fieldName string = values.Type().Field(i).Name

		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() == 0 {
				return fmt.Errorf("%s is empty", fieldName)
			}
		case reflect.String:
			if field.String() == "" {
				return fmt.Errorf("%s is empty", fieldName)
			}
		}
	}

	return nil
}

func (config *ServerConfiguration) GetEnvironment() string {
	return config.environment
}

func (config *ServerConfiguration) GetJwtSecretKey() []byte {
	return []byte(config.jwtSecretKey)
}
