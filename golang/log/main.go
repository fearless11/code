package main

import (
	"log"
	"os"
	// logrus hook
	// "github.com/sirupsen/logrus"
)

// https://godoc.org/github.com/sirupsen/logrus

func init() {
	// set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	f, err := os.OpenFile("log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// set log to file
		log.SetOutput(f)
	} else {
		log.Print("failed to log to file, using default stderr")
	}
}

func init() {
	// logrus.JSONFormatter{} or logrus.TextFormatter{}
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	// logrus.SetOutput(os.Stdout)
	// logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	log.Println("hello world")
}
