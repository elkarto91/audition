package models

//elkarto91@Author : Karthik
//Using Logrus logger utility, it gives quite a good look for the logs

import "github.com/sirupsen/logrus"

func SetLoggerText() *logrus.Logger {

	logger := logrus.New()
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05:01"
	Formatter.FullTimestamp = true
	Formatter.ForceColors = true
	Formatter.ForceQuote = true
	Formatter.DisableSorting = false // order!
	logrus.SetFormatter(Formatter)
	logrus.SetLevel(logrus.DebugLevel)
	logger.Formatter = Formatter
	return logger
}
