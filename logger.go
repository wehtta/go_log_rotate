package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
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

// var LogFile *os.File

func init() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	logger = Log4Go{
		LoggerConsole: logrus.New(),
		LoggerFile:    logrus.New(),
	}

	logger.LoggerConsole.Out = os.Stdout
	logger.LoggerConsole.Level = logrus.DebugLevel

	logger.LoggerFile.Out = &lumberjack.Logger{
		Filename:   "./error.log",
		MaxSize:    10, // megabytes
		MaxBackups: 7,
		MaxAge:     2, //days
	}

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
