package main

import (
	"github.com/nsf/termbox-go"
	"github.com/sirupsen/logrus"
)

func draw_all() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(1, 1, '@', termbox.ColorBlack|termbox.AttrBold, termbox.ColorRed)
	termbox.Flush()

}
func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)

	}
	defer termbox.Close()
	draw_all()
	x, y := termbox.Size()
	logrus.WithFields(logrus.Fields{
		"X": x,
		"Y": y,
	}).Info("Console size")
	logrus.Info("Start loop")
	//loop:
	for {
		x, y := termbox.Size()
		logrus.WithFields(logrus.Fields{
			"X": x,
			"Y": y,
		}).Info("Console size")
		logrus.Info("Start loop")
		//		switch ev := termbox.PollEvent(); ev.Type {
		//case termbox.EventKey:
		//	logrus.Info("EventKey event")
		//	switch ev.Key {
		//	case termbox.KeyEsc:
		//		logrus.Info("KeyEsc event")
		//		//				break loop
		//	}
		//case termbox.EventResize:
		//	logrus.Info("Resize event")
		//	draw_all()
		//}
	}

}
