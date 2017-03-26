package far2go

type Coordinate struct {
	X, Y uint
}
type MouseEvent struct {
	Coordinate
	dwButtonState     uint
	dwControlKeyState uint
	dwEventFlags      uint
}

type UIElement interface {
	ProcessKey(key int) (int)
	ProcessMouse(event *MouseEvent) (int)
	Hide()
	Show()
	ShowConsoleTitle()
	SetPosition(c1, c2 Coordinate)
	GetPosition() (c1, c2 *Coordinate)
	SetScreenPosition()
	ResizeConsole()
}
