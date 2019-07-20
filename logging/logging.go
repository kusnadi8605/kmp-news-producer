package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// levelMap
var levelMap = map[string]int{
	"DEBUG": 0,
	"INFO":  1,
	"WARN":  2,
	"ERROR": 3,
}

var prefixMap = map[string]string{
	"DEBUG": "DEBU",
	"INFO":  "INFO",
	"WARN":  "WARN",
	"ERROR": "ERRO",
}

// LogLevel is ...
var LogLevel = levelMap["DEBUG"]

// Filename is ...
var Filename = ""

// Init is ...
func Init(level string, filename string) {
	LogLevel = levelMap[level]
	Filename = filename
}

// SetLogLevel is ...
func SetLogLevel(level string) {
	LogLevel = levelMap[level]
}

// SetFilename is ...
func SetFilename(filename string) {
	Filename = filename
}

// Debugf is ...
func Debugf(msg string, a ...interface{}) {
	internalLog("DEBUG", msg, a)
}

// Debug is ...
func Debug(msg string) {
	Debugf(msg)
}

// Logf is ...
func Logf(msg string, a ...interface{}) {
	internalLog("INFO", msg, a)
}

// Log is ...
func Log(msg string) {
	Logf(msg)
}

// Warnf is ...
func Warnf(msg string, a ...interface{}) {
	internalLog("WARN", msg, a)
}

// Warn is ...
func Warn(msg string) {
	Warnf(msg)
}

// Errorf is ...
func Errorf(msg string, a ...interface{}) {
	internalLog("ERROR", msg, a)
}

// Error is ...
func Error(msg string) {
	Errorf(msg)
}

// internalLog is ...
func internalLog(level string, msg string, a ...interface{}) {
	if LogLevel > levelMap[level] {
		return
	}

	//prepare the message
	outputMsg := fmt.Sprintf(msg, a...)
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")

	// prepare the message
	outputMsg = fmt.Sprintf("%s %s %s", prefixMap[level], timestamp, outputMsg)

	// print to screen and append to log file
	fmt.Println(outputMsg)

	appendToLogFile(outputMsg)
}

// appendToLogFile is ...
func appendToLogFile(msg string) {
	if Filename == "" {
		return
	}

	// append log to file
	f, err := os.OpenFile(Filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Failed opening log file %v", err)
	}
	defer f.Close()
	fmt.Fprintln(f, msg)
}
