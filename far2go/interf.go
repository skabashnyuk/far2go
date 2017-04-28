package far2go

type ConsoleIo struct {
	Oem2Unicode        []rune
	BoxSymbols         []rune
	InitSize, CurSize  *Coordinate
	ScrX, ScrY         int
	PrevScrX, PrevScrY int
	InitialConsoleMode *ConsoleMode
	initTitle          string
	initWindowRect     *SmallRect
	initialSize        *Coordinate
	screenBuf          *ScreenBuf
}

type BoxType uint

const (
	NO_BOX           BoxType = iota
	SINGLE_BOX
	SHORT_SINGLE_BOX
	DOUBLE_BOX
	SHORT_DOUBLE_BOX
)

type BoxDefSymbols uint

const (
	BS_X_B0    BoxDefSymbols = iota
	BS_X_B1
	BS_X_B2
	BS_V1
	BS_R_H1V1
	BS_R_H2V1
	BS_R_H1V2
	BS_RT_H1V2
	BS_RT_H2V1
	BS_R_H2V2
	BS_V2
	BS_RT_H2V2
	BS_RB_H2V2
	BS_RB_H1V2
	BS_RB_H2V1
	BS_RT_H1V1
	BS_LB_H1V1
	BS_B_H1V1
	BS_T_H1V1
	BS_L_H1V1
	BS_H1
	BS_C_H1V1
	BS_L_H2V1
	BS_L_H1V2
	BS_LB_H2V2
	BS_LT_H2V2
	BS_B_H2V2
	BS_T_H2V2
	BS_L_H2V2
	BS_H2
	BS_C_H2V2
	BS_B_H2V1
	BS_B_H1V2
	BS_T_H2V1
	BS_T_H1V2
	BS_LB_H1V2
	BS_LB_H2V1
	BS_LT_H2V1
	BS_LT_H1V2
	BS_C_H1V2
	BS_C_H2V1
	BS_RB_H1V1
	BS_LT_H1V1
	BS_X_DB
	BS_X_DC
	BS_X_DD
	BS_X_DE
	BS_X_DF
)

func newConsoleIo() *ConsoleIo {
	return &ConsoleIo{
		PrevScrX:  -1,
		PrevScrY:  -1,
		screenBuf: NewScreenBuf(),
		BoxSymbols: []rune{
			'░',
			'▒',
			'▓',
			'│',
			'┤',
			'╡',
			'╢',
			'╖',
			'╕',
			'╣',
			'║',
			'╗',
			'╝',
			'╜',
			'╛',
			'┐',
			'└',
			'┴',
			'┬',
			'├',
			'─',
			'┼',
			'╞',
			'╟',
			'╚',
			'╔',
			'╩',
			'╦',
			'╠',
			'═',
			'╬',
			'╧',
			'╨',
			'╤',
			'╥',
			'╙',
			'╘',
			'╒',
			'╓',
			'╫',
			'╪',
			'┘',
			'┌',
			'█',
			'▄',
			'▌',
			'▐',
			'▀',
		},
	}
}

//void ShowTime(int ShowAlways);
func (Obj *ConsoleIo) InitConsole() {
	Obj.InitRecodeOutTable()
	SetControlHandler(Obj, true)
	Obj.InitialConsoleMode = GetMode(GetInputHandle())
	Obj.initTitle = GetTitle()
	Obj.initWindowRect = GetWindowRect()
	Obj.initialSize = GetSize()
	Obj.SetFarConsoleMode(false)

	WindowRect := GetWindowRect()
	Obj.InitSize = Obj.GetVideoMode()
	if Opt.WindowMode {
		ResetPosition()
	} else {
		if WindowRect.Left > 0 || WindowRect.Top > 0 || WindowRect.Right != Obj.InitSize.X-1 || WindowRect.Bottom != Obj.InitSize.Y-1 {
			newSize := Coordinate{
				X: WindowRect.Right - WindowRect.Left + 1,
				Y: WindowRect.Bottom - WindowRect.Top + 1,
			}
			SetSize(newSize)
			Obj.InitSize = Obj.GetVideoMode()

		}
	}
	Obj.CurSize = Obj.GetVideoMode()
	Obj.screenBuf.FillBuf()

}
func (Obj *ConsoleIo) Handle(CtrlType CtrlTypeEvent) {

}

//void CloseConsole();
//void SetFarConsoleMode(BOOL SetsActiveBuffer=FALSE);
func (Obj *ConsoleIo) SetFarConsoleMode(SetsActiveBuffer bool) {
	Mode := ENABLE_WINDOW_INPUT
	if Opt.Mouse {
		Mode |= ENABLE_MOUSE_INPUT | ENABLE_EXTENDED_FLAGS
	} else {
		Mode |= *Obj.InitialConsoleMode & (ENABLE_EXTENDED_FLAGS | ENABLE_QUICK_EDIT_MODE);
	}
	if SetsActiveBuffer {
		SetActiveScreenBuffer(GetOutputHandle())
	}
	Obj.ChangeConsoleMode(Mode)
}

//void ChangeConsoleMode(int Mode);
func (Obj *ConsoleIo) ChangeConsoleMode(Mode ConsoleMode) {
	ConIn := GetInputHandle()
	CurrentConsoleMode := GetMode(ConIn)
	if *CurrentConsoleMode != Mode {
		SetMode(ConIn, Mode)
	}
}

//void FlushInputBuffer();
//void SetVideoMode();
//void ChangeVideoMode(int Maximized);
//void ChangeVideoMode(int NumLines,int NumColumns);
//void GetVideoMode(COORD& Size);
func (Obj *ConsoleIo) GetVideoMode() *Coordinate {
	return nil
}

//void GenerateWINDOW_BUFFER_SIZE_EVENT(int Sx=-1, int Sy=-1);
//void SaveConsoleWindowRect();
//void RestoreConsoleWindowRect();
//
//void GotoXY(int X,int Y);
//int WhereX();
//int WhereY();
//void MoveCursor(int X,int Y);
//void GetCursorPos(SHORT& X,SHORT& Y);
//void SetCursorType(bool Visible, DWORD Size);
//void SetInitialCursorType();
//void GetCursorType(bool& Visible, DWORD& Size);
//void MoveRealCursor(int X,int Y);
//void GetRealCursorPos(SHORT& X,SHORT& Y);
//void ScrollScreen(int Count);
//
//void Text(int X, int Y, int Color, const WCHAR *Str);
//void Text(const WCHAR *Str);
//void Text(int MsgId);
//void VText(const WCHAR *Str);
//void HiText(const WCHAR *Str,int HiColor,int isVertText=0);
//void mprintf(const WCHAR *fmt,...);
//void vmprintf(const WCHAR *fmt,...);
//void PutText(int X1,int Y1,int X2,int Y2,const void *Src);
//void GetText(int X1,int Y1,int X2,int Y2,void *Dest,int DestSize);
//void BoxText(wchar_t Chr);
//void BoxText(const wchar_t *Str,int IsVert=0);
//
//void SetScreen(int X1,int Y1,int X2,int Y2,wchar_t Ch,int Color);
//void MakeShadow(int X1,int Y1,int X2,int Y2);
//void ChangeBlockColor(int X1,int Y1,int X2,int Y2,int Color);
//void SetColor(int Color);
//void SetRealColor(int Color);
//void ClearScreen(int Color);
//int GetColor();
//
//void Box(int x1,int y1,int x2,int y2,int Color,int Type);
//void ScrollBar(int X1,int Y1,int Length, unsigned int Current,unsigned int Total);
//bool ScrollBarRequired(UINT Length, UINT64 ItemsCount);
//bool ScrollBarEx(UINT X1,UINT Y1,UINT Length,UINT64 TopItem,UINT64 ItemsCount);
//void DrawLine(int Length,int Type, const wchar_t* UserSep=nullptr);
//#define ShowSeparator(Length,Type) DrawLine(Length,Type)
//#define ShowUserSeparator(Length,Type,UserSep) DrawLine(Length,Type,UserSep)
//WCHAR* MakeSeparator(int Length,WCHAR *DestStr,int Type=1, const wchar_t* UserSep=nullptr);
//
//void InitRecodeOutTable();
func (Obj *ConsoleIo) InitRecodeOutTable() {
}

//
//int WINAPI TextToCharInfo(const char *Text,WORD Attr, CHAR_INFO *CharInfo, int Length, DWORD Reserved);
//
//inline void SetVidChar(CHAR_INFO& CI,wchar_t Chr)
//{
//CI.Char.UnicodeChar = (Chr >= 0 && (Chr < L'\x20' || Chr == L'\x7f')) ? Oem2Unicode[Chr] : Chr;
//}
//
//int HiStrlen(const wchar_t *Str);
//int HiFindRealPos(const wchar_t *Str, int Pos, BOOL ShowAmp);
//int HiFindNextVisualPos(const wchar_t *Str, int Pos, int Direct);
//FARString& HiText2Str(FARString& strDest, const wchar_t *Str);
//#define RemoveHighlights(Str) RemoveChar(Str,L'&')
//
//bool IsFullscreen();
