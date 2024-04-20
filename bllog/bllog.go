// Package bllog sets up logging functions
package bllog

import (
	"log"
	"os"
)

var projectLogFile *os.File

// InitProjectLog sets up a project log that logs to a file and returns a func that will close the log file.
func InitProjectLog(filename string) {
	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	projectLogFile = logFile
}

// CloseLog closes the project log.
func CloseLog() {
	projectLogFile.Close()
}
