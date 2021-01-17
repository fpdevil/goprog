package main

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/fpdevil/goprog/random-stuff/generic/logger/hooks"
	"github.com/sirupsen/logrus"
)

func main() {
	// do not send logs anywhere by default
	logrus.SetOutput(ioutil.Discard)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// the custom hook will trigger and flush them to stderr
	// only when an error comes
	logrus.AddHook(&hooks.Hook{Writer: os.Stderr})
	updateAPI("making a call to the API's")
}

func updateAPI(msg string) {
	txnID := "abcd-1234-0x1xff"
	logger := logrus.WithFields(logrus.Fields{"txnID": txnID})

	api(msg, logger)
	updateProfile(msg, logger)
}

func api(msg string, logger *logrus.Entry) {
	logger.Info("api call started")
	// call goes to api hub
	logger.Info("api hub invoked")
}

func updateProfile(msg string, logger *logrus.Entry) {
	logger.Info("sending profile information started")
	err := callProfileAPI(msg)
	if err != nil {
		logger.Errorf("calling Profile api failed error: %s", err)
	}
	logger.Info("Sent to callProfileApi")
}

func callProfileAPI(msg string) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(10) > 6 {
		return errors.New("HTTP 400 Error")
	}
	return nil
}
