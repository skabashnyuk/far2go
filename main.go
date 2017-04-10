package main

import (
	"github.com/skabashnyuk/far2go/far2go"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	//log.Debug("Useful debugging information.")
	//log.Info("Something noteworthy happened!")
	//log.Warn("You should probably take a look at this.")
	//log.Error("Something failed but I'm not quitting.")
	//// Calls os.Exit(1) after logging
	//log.Fatal("Bye.")
	//// Calls panic() after logging
	//log.Panic("I'm bailing.")
	logrus.Debug("Starting controll object")
	CtrlObj := far2go.NewControlObject()
	CtrlObj.CreateFilePanels()
	CtrlObj.Macro.LoadMacros()
}
