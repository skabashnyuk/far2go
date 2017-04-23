package far2go

import (
	"os"
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

func GetHandle() *os.File {
	return os.Stderr
}

func IsZoomed() bool {
	return false
}

func IsIconic() bool {
	return false
}

func GetSize() (*Coordinate) {
	return nil
}

func SetSize(size Coordinate) {

}

func SetWindowRect(consoleWindow SmallRect) {

}

func GetWindowRect() (*SmallRect) {
	return nil
}

func GetWorkingRect() (*SmallRect) {
	return nil
}

func SetTitle(strTitle string) {
}

func GetTitle() (string) {
	return ""
}
func GetKeyboardLayoutName() string {
	return ""
}

func SetControlHandler(HandlerRoutine ControlHandler, Add bool) {

}

func GetMode(handle *os.File) (*ConsoleMode) {
	return nil
}

func SetMode(handle *os.File, mode ConsoleMode) () {

}

func PeekInput(buffer *InputRecord, length uint, numberOfEventsRead uint) () {

}

func ReadInput(buffer *InputRecord, length uint, numberOfEventsRead uint) () {

}

func WriteInput(buffer *InputRecord, length uint, numberOfEventsRead uint) () {

}

func ReadOutput(buffer *CharInfo, bufferSize Coordinate, bufferCoord Coordinate, readRegion *SmallRect) () {

}

func WriteOutput(buffer *CharInfo, bufferSize Coordinate, bufferCoord Coordinate, readRegion *SmallRect) () {

}

func Write(buffer string, numberOfCharsToWrite bool) () {

}

func GetTextAttributes() (uint) {
	return 0
}

func SetTextAttributes(Attributes uint) () {

}

func GetCursorInfo() (*ConsoleCursorInfo) {
	return nil
}

func SetCursorInfo(consoleCursorInfo *ConsoleCursorInfo) () {

}

func GetCursorPosition() (*Coordinate) {
	return nil
}

func SetCursorPosition(position Coordinate) () {

}

func FlushInputBuffer() () {

}

func GetNumberOfInputEvents() (uint) {
	return 0
}

func GetLargestWindowSize() *Coordinate {
	return nil
}

func SetActiveScreenBuffer(consoleOutput *os.File) () {

}

func ClearExtraRegions(color PaletteColor) () {

}

func ScrollWindow(lines int, columns int) () {

}

func ResetPosition() () {

}
func GetDelta() (int) {
	return 0
}
