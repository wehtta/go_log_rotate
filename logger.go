package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	// "github.com/Sirupsen/logrus/hooks/airbrake"
	"os"
	"time"
)

type Logger logrus.Logger

type Log4Go struct {
	LoggerConsole *logrus.Logger
	LoggerFile    *logrus.Logger
}

func (logger *Log4Go) Info(args ...interface{}) {
	logger.LoggerConsole.WithFields(logrus.Fields{
		"time": time.Now(),
	}).Info(args)
}

// error should be output to both file and console
func (logger *Log4Go) Error(args ...interface{}) {
	logger.LoggerConsole.WithFields(logrus.Fields{
		"time": time.Now(),
	}).Error(args)

	logger.LoggerFile.WithFields(logrus.Fields{
		"time": time.Now(),
	}).Error(args)
}

var logger Log4Go
var LogFile *os.File

func init() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	logfile, err := os.OpenFile("./error.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic(err)
	}

	LogFile = logfile
	logger = Log4Go{
		LoggerConsole: logrus.New(),
		LoggerFile:    logrus.New(),
	}

	logger.LoggerConsole.Out = os.Stdout
	logger.LoggerConsole.Level = logrus.DebugLevel

	logger.LoggerFile.Out = logfile
	logger.LoggerFile.Level = logrus.DebugLevel
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
