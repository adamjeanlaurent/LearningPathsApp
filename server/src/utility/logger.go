package utility

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/fatih/color"
)

func getFilenameFromPath(filePath string) string {
	return filepath.Base(filePath)
}

func getLogHeading() string {
	currentTime := time.Now()

	_, file, line, _ := runtime.Caller(1)

	var fileName = getFilenameFromPath(file)

	return fmt.Sprintf("Log: [%s] %s:%d : ", currentTime.Format("2006-01-02 15:04:05"), fileName, line)
}

func LogError(err error) {

	var logHeading string = getLogHeading()
	fmt.Print(logHeading)
	color.Red(err.Error())
}
