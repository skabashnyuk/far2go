package far2go

const (
	SBFLAGS_FLUSHED        = 0x00000001
	SBFLAGS_FLUSHEDCURPOS  = 0x00000002
	SBFLAGS_FLUSHEDCURTYPE = 0x00000004
	SBFLAGS_USESHADOW      = 0x00000008
)

type ScreenBuf struct {
	sBFlags           *BitFlags
	buf               []CharInfo
	shadow            []CharInfo
	macroChar         CharInfo
	macroCharUsed     bool
	elevationChar     CharInfo
	elevationCharUsed bool
	bufX, bufY        uint
	curX, curY        uint
	curVisible        bool
	curSize           uint
}

func NewScreenBuf() *ScreenBuf {
	return &ScreenBuf{
		sBFlags: NewBitFlags(SBFLAGS_FLUSHED | SBFLAGS_FLUSHEDCURPOS | SBFLAGS_FLUSHEDCURTYPE),
	}
}

//void AllocBuf(int X,int Y);
//void Lock();
//void Unlock();
//int  GetLockCount() {return(LockCount);};
//void SetLockCount(int Count) {LockCount=Count;};
//void ResetShadow();
//void MoveCursor(int X,int Y);
//void GetCursorPos(SHORT& X,SHORT& Y);
//void SetCursorType(bool Visible, DWORD Size);
//void GetCursorType(bool& Visible, DWORD& Size);
//
//void FillBuf();
func (Obj *ScreenBuf) FillBuf() {
	BufferSize := Coordinate{Obj.bufX, Obj.bufY}
	BufferCoord := Coordinate{0, 0}
	ReadRegion := SmallRect{Obj.bufX, Obj.bufY, Obj.bufX - 1, Obj.bufY - 1 }
	ReadOutput(Obj.buf, BufferSize, BufferCoord, ReadRegion)
	length := Obj.bufX * Obj.bufY

	//TODO test copy of slices
	copy(Obj.shadow[:length], Obj.buf[:length])
	Obj.sBFlags.Set(SBFLAGS_USESHADOW)

	CursorPosition := GetCursorPosition()
	Obj.curX = CursorPosition.X
	Obj.curY = CursorPosition.Y
}

//void Read(int X1,int Y1,int X2,int Y2,CHAR_INFO *Text,int MaxTextLength);
//void Write(int X,int Y,const CHAR_INFO *Text,int TextLength);
//void RestoreMacroChar();
//void RestoreElevationChar();
//
//void ApplyColorMask(int X1,int Y1,int X2,int Y2,WORD ColorMask);
//void ApplyColor(int X1,int Y1,int X2,int Y2,WORD Color);
//void ApplyColor(int X1,int Y1,int X2,int Y2,int Color,WORD ExceptColor);
//void FillRect(int X1,int Y1,int X2,int Y2,WCHAR Ch,WORD Color);
//
//void Scroll(int);
//void Flush();
