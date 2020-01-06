package logging

import (
	"log"
	"os"
	"runtime"
	"strconv"
)

const pathLogs = "logs/logs.log"

// Types of errors
const (
	InfoType     = "[INFO]:"
	CriticalType = "[CRITICAL]:"
	WarningType  = "[WARNING]:"
	ErrorType    = "[ERROR]:"
	DebugType    = "[DEBUG]:"
	RuntimeType  = "[RUNTIME]:"
)

// InfoColor		- 34 Blue
// CriticalColor	- 36 Cyan
// WarningColor		- 33 Yellow
// ErrorColor		- 31 Red
// DebugColor		- 35 Magenta
// RuntimeColor		- 32 Green
//					- 37 White
//					- 30 Black
const (
	InfoColor     = "\033[1;34m%s\033[0m"
	CriticalColor = "\033[1;36m%s\033[0m"
	WarningColor  = "\033[1;33m%s\033[0m"
	ErrorColor    = "\033[1;31m%s\033[0m"
	DebugColor    = "\033[0;35m%s\033[0m"
	RuntimeColor  = "\033[0;32m%s\033[0m"
)

// Info ...
// save bool write to file
func Info(message string, save ...bool) {
	_, fn, line, _ := runtime.Caller(1)
	message = InfoType + fn + ":" + strconv.Itoa(line) + " -> " + message

	if len(save) > 0 && save[0] == true {
		saveLogToFile(message)
	}
	log.Printf(InfoColor, message)
}

// Critical ...
// save bool write to file
func Critical(message string, save ...bool) {
	_, fn, line, _ := runtime.Caller(1)
	message = CriticalType + fn + ":" + strconv.Itoa(line) + " -> " + message

	if len(save) > 0 && save[0] == true {
		saveLogToFile(message)
	}
	log.Printf(CriticalColor, message)
}

// Warning ...
// save bool write to file
func Warning(message string, save ...bool) {
	_, fn, line, _ := runtime.Caller(1)
	message = WarningType + fn + ":" + strconv.Itoa(line) + " -> " + message

	if len(save) > 0 && save[0] == true {
		saveLogToFile(message)
	}
	log.Printf(WarningColor, message)
}

// Error ...
// save bool write to file
func Error(message string, save ...bool) {
	_, fn, line, _ := runtime.Caller(1)
	message = ErrorType + fn + ":" + strconv.Itoa(line) + " -> " + message

	if len(save) > 0 && save[0] == true {
		saveLogToFile(message)
	}
	log.Printf(ErrorColor, message)
}

// Debug ...
// save bool write to file
func Debug(message string, save ...bool) {
	_, fn, line, _ := runtime.Caller(1)
	message = DebugType + fn + ":" + strconv.Itoa(line) + " -> " + message

	if len(save) > 0 && save[0] == true {
		saveLogToFile(message)
	}
	log.Printf(DebugColor, message)
}

// Runtime ...
// save bool write to file
func Runtime(message string, save ...bool) {
	_, fn, line, _ := runtime.Caller(1)
	message = RuntimeType + fn + ":" + strconv.Itoa(line) + " -> " + message

	if len(save) > 0 && save[0] == true {
		saveLogToFile(message)
	}
	log.Printf(RuntimeColor, message)
}

// saveLogToFile ...
func saveLogToFile(message string) {

	if folder, _ := exists(pathLogs); folder == false {
		if err := os.Mkdir("logs", os.ModePerm); err != nil {
			Error(err.Error(), false)
		}
	}
	
	f, err := os.OpenFile(pathLogs, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Error(err.Error(), false)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(message)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
