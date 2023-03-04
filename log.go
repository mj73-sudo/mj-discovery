package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var LogJson *logrus.Logger
var LogText *logrus.Logger

func LoadLogConfig() {
	//read log level from config file
	logLevel := viper.GetString("log.level")

	LogJson = logrus.New()
	LogJson.SetFormatter(&logrus.JSONFormatter{})
	LogJson.SetOutput(os.Stdout)
	//set log level to log json
	LogJson.SetLevel(getLogLevel(strings.ToUpper(logLevel)))

	LogText = logrus.New()
	LogText.SetOutput(os.Stdout)
	//set log level to log text
	LogText.SetLevel(getLogLevel(strings.ToUpper(logLevel)))
}

func getLogLevel(level string) logrus.Level {
	switch level {
	case "TRACE":
		return logrus.TraceLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "FATAL":
		return logrus.FatalLevel
	}
	return 0
}
