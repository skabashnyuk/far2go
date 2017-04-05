package far2go

type Coordinate struct {
	X, Y uint
}

type SmallRect struct {
	Left, Top, Right, Bottom uint
}

type MouseEventRecord struct {
	Coordinate
	ButtonState     uint
	ControlKeyState uint
	EventFlags      uint
}

type EventType uint

const (
	KEY_EVENT                EventType = 0x0001 // Event contains key event record
	MOUSE_EVENT              EventType = 0x0002 // Event contains mouse event record
	WINDOW_BUFFER_SIZE_EVENT EventType = 0x0004 // Event contains window change event record
	MENU_EVENT               EventType = 0x0008 // Event contains menu event record
	FOCUS_EVENT              EventType = 0x0010 // event contains focus change
)

type ControlKeyState uint

const (
	RIGHT_ALT_PRESSED  ControlKeyState = 0x0001     // the right alt key is pressed.
	LEFT_ALT_PRESSED   ControlKeyState = 0x0002     // the left alt key is pressed.
	RIGHT_CTRL_PRESSED ControlKeyState = 0x0004     // the right ctrl key is pressed.
	LEFT_CTRL_PRESSED  ControlKeyState = 0x0008     // the left ctrl key is pressed.
	SHIFT_PRESSED      ControlKeyState = 0x0010     // the shift key is pressed.
	NUMLOCK_ON         ControlKeyState = 0x0020     // the numlock light is on.
	SCROLLLOCK_ON      ControlKeyState = 0x0040     // the scrolllock light is on.
	CAPSLOCK_ON        ControlKeyState = 0x0080     // the capslock light is on.
	ENHANCED_KEY       ControlKeyState = 0x0100     // the key is enhanced.
	NLS_DBCSCHAR       ControlKeyState = 0x00010000 // DBCS for JPN: SBCS/DBCS mode.
	NLS_ALPHANUMERIC   ControlKeyState = 0x00000000 // DBCS for JPN: Alphanumeric mode.
	NLS_KATAKANA       ControlKeyState = 0x00020000 // DBCS for JPN: Katakana mode.
	NLS_HIRAGANA       ControlKeyState = 0x00040000 // DBCS for JPN: Hiragana mode.
	NLS_ROMAN          ControlKeyState = 0x00400000 // DBCS for JPN: Roman/Noroman mode.
	NLS_IME_CONVERSION ControlKeyState = 0x00800000 // DBCS for JPN: IME conversion.
	NLS_IME_DISABLE    ControlKeyState = 0x20000000 // DBCS for JPN: IME enable/disable.
)

type KeyEventRecord struct {
	KeyDown         bool
	RepeatCount     uint
	VirtualKeyCode  VirtualKey
	VirtualScanCode VirtualKey
	Char            rune
	ControlKeyState ControlKeyState
}
type WindowBufferSizeRecord struct {
	Size Coordinate
}

type MenuEventRecord struct {
	CommandId uint
}

type FocusEventRecord struct {
	SetFocus bool
}
type Event struct {
	MouseEvent            MouseEventRecord
	KeyEvent              KeyEventRecord
	WindowBufferSizeEvent WindowBufferSizeRecord
	MenuEvent             MenuEventRecord
	FocusEvent            FocusEventRecord
}

type InputRecord struct {
	EventType EventType
	Event
}

type CharInfo struct {
	Char       rune
	Attributes uint
}
type ConsoleCursorInfo struct {
	Size uint
	Visible bool
}

type UIElement interface {
	ProcessKey(key int) (int)
	ProcessMouse(event *MouseEventRecord) (int)
	Hide()
	Show()
	ShowConsoleTitle()
	SetPosition(c1, c2 Coordinate)
	GetPosition() (c1, c2 *Coordinate)
	SetScreenPosition()
	ResizeConsole()
}
