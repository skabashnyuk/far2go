package far2go

import (
	"os"
	"errors"
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

func GetSize() (Coordinate, error) {
	return nil, errors.New("NotImplemented")
}

func SetSize(size Coordinate) error {
	return errors.New("NotImplemented")
}

func SetWindowRect(consoleWindow SmallRect) error {
	return errors.New("NotImplemented")
}

func GetWindowRect() (SmallRect, error) {
	return nil, errors.New("NotImplemented")
}

func GetWorkingRect() (SmallRect, error) {
	return nil, errors.New("NotImplemented")
}

func SetTitle(strTitle string) error {
	return errors.New("NotImplemented")
}

func GetTitle() (string, error) {
	return nil, errors.New("NotImplemented")
}
func GetKeyboardLayoutName() (string, error) {
	return nil, errors.New("NotImplemented")
}


//
//bool GetMode(HANDLE ConsoleHandle, DWORD& Mode);
//bool SetMode(HANDLE ConsoleHandle, DWORD Mode);
//
//bool PeekInput(INPUT_RECORD& Buffer, DWORD Length, DWORD& NumberOfEventsRead);
//bool ReadInput(INPUT_RECORD& Buffer, DWORD Length, DWORD& NumberOfEventsRead);
//bool WriteInput(INPUT_RECORD& Buffer, DWORD Length, DWORD& NumberOfEventsWritten);
//bool ReadOutput(CHAR_INFO& Buffer, COORD BufferSize, COORD BufferCoord, SMALL_RECT& ReadRegion);
//bool WriteOutput(const CHAR_INFO& Buffer, COORD BufferSize, COORD BufferCoord, SMALL_RECT& WriteRegion);
//bool Write(LPCWSTR Buffer, DWORD NumberOfCharsToWrite);
//
//bool GetTextAttributes(WORD& Attributes);
//bool SetTextAttributes(WORD Attributes);
//
//bool GetCursorInfo(CONSOLE_CURSOR_INFO& ConsoleCursorInfo);
//bool SetCursorInfo(const CONSOLE_CURSOR_INFO& ConsoleCursorInfo);
//
//bool GetCursorPosition(COORD& Position);
//bool SetCursorPosition(COORD Position);
//
//bool FlushInputBuffer();
//
//bool GetNumberOfInputEvents(DWORD& NumberOfEvents);
//
//COORD GetLargestWindowSize();
//
//bool SetActiveScreenBuffer(HANDLE ConsoleOutput);
//
//bool ClearExtraRegions(WORD Color);
//
//bool ScrollWindow(int Lines,int Columns=0);
//bool ResetPosition();
//int GetDelta()
