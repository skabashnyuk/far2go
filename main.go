package main

import (
	"github.com/skabashnyuk/far2go/far2go"
	log "github.com/Sirupsen/logrus"

)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")

	CtrlObj := far2go.NewControlObject()
	CtrlObj.CreateFilePanels()
}
