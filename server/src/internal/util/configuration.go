package configuration

import "os"

var (
	environment  = os.Getenv("MYAPP_ENV")
	jwtSecretKey = os.Getenv("JWT_SECRET_KEY")
)

func GetEnvironment() string {
	return environment
}

func GetJwtSecretKey() string {
	return jwtSecretKey
}
