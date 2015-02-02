package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	// "github.com/Sirupsen/logrus/hooks/airbrake"
	"os"
)

type Logger logrus.Logger

type Log4Go struct {
	LoggerConsole *logrus.Logger
	LoggerFile    *logrus.Logger
}

func (logger *Log4Go) Info(args ...interface{}) {
	logger.LoggerConsole.Info(args)
}

// error should be output to both file and console
func (logger *Log4Go) Error(args ...interface{}) {
	logger.LoggerConsole.Error(args)
	logger.LoggerFile.Error(args)
}

var logger Log4Go

// var logfile *os.File

func init() {
	logfile, err := os.OpenFile("./error.log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	var LoggerConsole = logrus.New()
	LoggerConsole.Out = os.Stdout
	LoggerConsole.Level = logrus.DebugLevel

	var LoggerFile = logrus.New()
	LoggerFile.Out = logfile
	LoggerFile.Level = logrus.DebugLevel

	logger = Log4Go{
		LoggerConsole,
		LoggerFile,
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		// Close(logfile)
	}()

	logger.Info("A group of walrus emerges from the ocean")
	logger.Error("The group's number increased tremendously!")
}
