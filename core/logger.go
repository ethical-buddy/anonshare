package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	debugMode bool
	infoLog   = log.New(os.Stdout, "", 0)
	errorLog  = log.New(os.Stderr, "", 0)
)

const (
	green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
	blue   = "\033[34m"
	reset  = "\033[0m"
)

// InitLogger enables or disables debug logging
func InitLogger(debug bool) {
	debugMode = debug
	Info("Logger initialized. Debug mode:", debug)
}

func Info(v ...any) {
	infoLog.Println(green + "[INFO] " + reset + fmt.Sprint(v...))
}

func Debug(v ...any) {
	if debugMode {
		infoLog.Println(blue + "[DEBUG] " + reset + fmt.Sprint(v...))
	}
}

func Warn(v ...any) {
	infoLog.Println(yellow + "[WARN] " + reset + fmt.Sprint(v...))
}

func Error(v ...any) {
	errorLog.Println(red + "[ERROR] " + reset + fmt.Sprint(v...))
}

