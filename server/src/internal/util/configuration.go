package configuration

import "os"

var ENV string = os.Getenv("MYAPP_ENV")
