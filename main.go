package main

import (
	"github.com/skabashnyuk/far2go/far2go"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	far2go.Main()
	//log.Debug("Useful debugging information.")
	//log.Info("Something noteworthy happened!")
	//log.Warn("You should probably take a look at this.")
	//log.Error("Something failed but I'm not quitting.")
	//// Calls os.Exit(1) after logging
	//log.Fatal("Bye.")
	//// Calls panic() after logging
	//log.Panic("I'm bailing.")
	//logrus.Debug("Starting controll object")
	//CtrlObj := far2go.NewControlObject()
	//Console.GetTextAttributes(InitAttributes);
	//SetRealColor(COL_COMMANDLINEUSERSCREEN);
	//CtrlObj.Init()
	//CtrlObj.CreateFilePanels()
	//CtrlObj.Macro.LoadMacros()
	//FrameManager->EnterMainLoop();
	//SetScreen(0,0,ScrX,ScrY,L' ',COL_COMMANDLINEUSERSCREEN);
	//Console.SetTextAttributes(InitAttributes);
	//ScrBuf.ResetShadow();
	//ScrBuf.Flush();
	//MoveRealCursor(0,0);
	//CloseConsole();
}
