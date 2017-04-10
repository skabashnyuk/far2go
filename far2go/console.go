package far2go

import (
	"os"
	"errors"
)

type ConsoleMode uint

const (
	ENABLE_PROCESSED_INPUT ConsoleMode = 0x0001
	ENABLE_LINE_INPUT      ConsoleMode = 0x0002
	ENABLE_ECHO_INPUT      ConsoleMode = 0x0004
	ENABLE_WINDOW_INPUT    ConsoleMode = 0x0008
	ENABLE_MOUSE_INPUT     ConsoleMode = 0x0010
	ENABLE_INSERT_MODE     ConsoleMode = 0x0020
	ENABLE_QUICK_EDIT_MODE ConsoleMode = 0x0040
	ENABLE_EXTENDED_FLAGS  ConsoleMode = 0x0080
	ENABLE_AUTO_POSITION   ConsoleMode = 0x0100
)

func GetInputHandle() *os.File {
	return os.Stdin
}

func GetOutputHandle() *os.File {
	return os.Stdout
}

func GetErrorHandle() *os.File {
	return os.Stderr
}

func IsZoomed() bool {
	return false
}

func IsIconic() bool {
	return false
}

func GetSize() (*Coordinate, error) {
	return nil, errors.New("NotImplemented")
}

func SetSize(size Coordinate) error {
	return errors.New("NotImplemented")
}

func SetWindowRect(consoleWindow SmallRect) error {
	return errors.New("NotImplemented")
}

func GetWindowRect() (*SmallRect, error) {
	return nil, errors.New("NotImplemented")
}

func GetWorkingRect() (*SmallRect, error) {
	return nil, errors.New("NotImplemented")
}

func SetTitle(strTitle string) error {
	return errors.New("NotImplemented")
}

func GetTitle() (string, error) {
	return "", errors.New("NotImplemented")
}
func GetKeyboardLayoutName() (string, error) {
	return "", errors.New("NotImplemented")
}

func GetMode(handle *os.File) (*ConsoleMode, error) {
	return nil, errors.New("NotImplemented")
}

func SetMode(handle *os.File, mode ConsoleMode) (error) {
	return errors.New("NotImplemented")
}

func PeekInput(buffer *InputRecord, length uint, numberOfEventsRead uint) (error) {
	return errors.New("NotImplemented")
}

func ReadInput(buffer *InputRecord, length uint, numberOfEventsRead uint) (error) {
	return errors.New("NotImplemented")
}

func WriteInput(buffer *InputRecord, length uint, numberOfEventsRead uint) (error) {
	return errors.New("NotImplemented")
}

func ReadOutput(buffer *CharInfo, bufferSize Coordinate, bufferCoord Coordinate, readRegion *SmallRect) (error) {
	return errors.New("NotImplemented")
}

func WriteOutput(buffer *CharInfo, bufferSize Coordinate, bufferCoord Coordinate, readRegion *SmallRect) (error) {
	return errors.New("NotImplemented")
}

func Write(buffer string, numberOfCharsToWrite bool) (error) {
	return errors.New("NotImplemented")
}

func GetTextAttributes() (uint, error) {
	return 0, errors.New("NotImplemented")
}

func SetTextAttributes(Attributes uint) (error) {
	return errors.New("NotImplemented")
}

func GetCursorInfo() (*ConsoleCursorInfo, error) {
	return nil, errors.New("NotImplemented")
}

func SetCursorInfo(consoleCursorInfo *ConsoleCursorInfo) (error) {
	return errors.New("NotImplemented")
}

func GetCursorPosition() (*Coordinate, error) {
	return nil, errors.New("NotImplemented")
}

func SetCursorPosition(position Coordinate) (error) {
	return errors.New("NotImplemented")
}

func FlushInputBuffer() (error) {
	return errors.New("NotImplemented")
}

func GetNumberOfInputEvents() (uint, error) {
	return 0, errors.New("NotImplemented")
}

func GetLargestWindowSize() *Coordinate {
	return nil
}

func SetActiveScreenBuffer(consoleOutput *os.File) (error) {
	return errors.New("NotImplemented")
}

func ClearExtraRegions(color PaletteColors) (error) {
	return errors.New("NotImplemented")
}

func ScrollWindow(lines int, columns int) (error) {
	return errors.New("NotImplemented")
}

func ResetPosition() (error) {
	return errors.New("NotImplemented")
}
func GetDelta() (int) {
	return 0
}
