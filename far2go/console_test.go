package far2go

import (
	"testing"
	"github.com/sirupsen/logrus"
	"github.com/nsf/termbox-go"
)

func TestSize(t *testing.T) {
loop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(1, 1, '@', termbox.ColorBlack|termbox.AttrBold, termbox.ColorRed)
		termbox.Flush()
		coordinate := GetSize()
		logrus.WithFields(logrus.Fields{
			"X": coordinate.X,
			"Y": coordinate.Y,
		}).Info("Console size")
		logrus.Info("Start loop")
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			logrus.Info("EventKey event")
			switch ev.Key {
			case termbox.KeyEsc:
				logrus.Info("KeyEsc event")
				break loop
			}
		case termbox.EventResize:
			logrus.Info("Resize event")

		}
	}

}
