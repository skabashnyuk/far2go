package far2go

// Младший байт (маска 0xFF) юзается классом ScreenObject!!!
const (
	FEDITLINE_MARKINGBLOCK     = 0x00000100
	FEDITLINE_DROPDOWNBOX      = 0x00000200
	FEDITLINE_CLEARFLAG        = 0x00000400
	FEDITLINE_PASSWORDMODE     = 0x00000800
	FEDITLINE_EDITBEYONDEND    = 0x00001000
	FEDITLINE_EDITORMODE       = 0x00002000
	FEDITLINE_OVERTYPE         = 0x00004000
	FEDITLINE_DELREMOVESBLOCKS = 0x00008000 // Del удаляет блоки (Opt.EditorDelRemovesBlocks)
	FEDITLINE_PERSISTENTBLOCKS = 0x00010000 // Постоянные блоки (Opt.EditorPersistentBlocks)
	FEDITLINE_SHOWWHITESPACE   = 0x00020000
	FEDITLINE_READONLY         = 0x00040000
	FEDITLINE_CURSORVISIBLE    = 0x00080000
	// Если ни один из FEDITLINE_PARENT_ не указан (или указаны оба) то Edit
	// явно не в диалоге юзается.
	FEDITLINE_PARENT_SINGLELINE = 0x00100000 // обычная строка ввода в диалоге
	FEDITLINE_PARENT_MULTILINE  = 0x00200000 // для будущего Memo-Edit (DI_EDITOR или DIF_MULTILINE)
	FEDITLINE_PARENT_EDITOR     = 0x00400000 // "вверху" обычный редактор
)

type ColorItem struct {
	StartPos int
	EndPos   int
	Color    int
}

const (
	SETCP_NOERROR    = 0x00000000
	SETCP_WC2MBERROR = 0x00000001
	SETCP_MB2WCERROR = 0x00000002
	SETCP_OTHERERROR = 0x10000000
)

type EditChangeFunc func()
type Callback struct {
	Active     bool
	M_Callback *EditChangeFunc
}

type Edit struct {
	ScreenObject
	Next      *Edit
	Prev      *Edit
	Str       string
	StrSize   int
	MaxLength int

	Mask           string
	ColorList      *ColorItem
	ColorCount     int
	Color          int
	SelColor       int
	ColorUnChanged int

	LeftPos    int
	CurPos     int
	PrevCurPos int

	TabSize int

	TabExpandMode int
	MSelStart     int
	SelStart      int
	SelEnd        int
	EndType       int
	CursorSize    int
	CursorPos     int

	WordDiv  string
	CodePage uint

	M_Callback Callback
}

func (obj *Edit) GetString() string {
	return obj.Str
}

//void  GetString(wchar_t *Str, int MaxSize);


//virtual void   DisplayObject();
//int    InsertKey(int Key);
//int    RecurseProcessKey(int Key);
//void   DeleteBlock();
//void   ApplyColor();
//int    GetNextCursorPos(int Position,int Where);
//void   RefreshStrByMask(int InitMode=FALSE);
//int KeyMatchedMask(int Key);
//
//int ProcessCtrlQ();
//int ProcessInsDate(const wchar_t *Str);
//int ProcessInsPlainText(const wchar_t *Str);
//
//int CheckCharMask(wchar_t Chr);
//int ProcessInsPath(int Key,int PrevSelStart=-1,int PrevSelEnd=0);
//
//int RealPosToTab(int PrevLength, int PrevPos, int Pos, int* CorrectPos);
//
//inline const wchar_t* WordDiv(void) {return strWordDiv->CPtr();};
//DWORD SetCodePage(UINT codepage);  //BUGBUG
//UINT GetCodePage();  //BUGBUG
//
//virtual void  FastShow();
//virtual int   ProcessKey(int Key);
//virtual int   ProcessMouse(MOUSE_EVENT_RECORD *MouseEvent);
//virtual int64_t VMProcess(int OpCode,void *vParam=nullptr,int64_t iParam=0);
//
////   ! Функция установки текущих Color,SelColor и ColorUnChanged!
//void  SetObjectColor(int Color,int SelColor=0xf,int ColorUnChanged=COL_DIALOGEDITUNCHANGED);
////   + Функция получения текущих Color,SelColor
//long  GetObjectColor() {return MAKELONG(Color,SelColor);}
//int   GetObjectColorUnChanged() {return ColorUnChanged;}
//
//void SetTabSize(int NewSize) { TabSize=NewSize; }
//int  GetTabSize() {return TabSize; }
//
//void SetDelRemovesBlocks(int Mode) {Flags.Change(FEDITLINE_DELREMOVESBLOCKS,Mode);}
//int  GetDelRemovesBlocks() {return Flags.Check(FEDITLINE_DELREMOVESBLOCKS); }
//
//void SetPersistentBlocks(int Mode) {Flags.Change(FEDITLINE_PERSISTENTBLOCKS,Mode);}
//int  GetPersistentBlocks() {return Flags.Check(FEDITLINE_PERSISTENTBLOCKS); }
//
//void SetShowWhiteSpace(int Mode) {Flags.Change(FEDITLINE_SHOWWHITESPACE,Mode);}
//
//void  GetString(wchar_t *Str, int MaxSize);
//void  GetString(FARString &strStr);
//
//const wchar_t* GetStringAddr();
//
//void  SetHiString(const wchar_t *Str);
//void  SetString(const wchar_t *Str,int Length=-1);
//
//void  SetBinaryString(const wchar_t *Str,int Length);
//
//void  GetBinaryString(const wchar_t **Str, const wchar_t **EOL,int &Length);
//
//void  SetEOL(const wchar_t *EOL);
//const wchar_t *GetEOL();
//
//int   GetSelString(wchar_t *Str,int MaxSize);
//int   GetSelString(FARString &strStr);
//
//int   GetLength();
//
//void  InsertString(const wchar_t *Str);
//void  InsertBinaryString(const wchar_t *Str,int Length);
//
//int   Search(const FARString& Str,FARString& ReplaceStr,int Position,int Case,int WholeWords,int Reverse,int Regexp, int *SearchLength);
//
//void  SetClearFlag(int Flag) {Flags.Change(FEDITLINE_CLEARFLAG,Flag);}
//int   GetClearFlag() {return Flags.Check(FEDITLINE_CLEARFLAG);}
//void  SetCurPos(int NewPos) {CurPos=NewPos; PrevCurPos=NewPos;}
//int   GetCurPos() {return(CurPos);}
//int   GetTabCurPos();
//void  SetTabCurPos(int NewPos);
//int   GetLeftPos() {return(LeftPos);}
//void  SetLeftPos(int NewPos) {LeftPos=NewPos;}
//void  SetPasswordMode(int Mode) {Flags.Change(FEDITLINE_PASSWORDMODE,Mode);};
//void  SetMaxLength(int Length) {MaxLength=Length;};
//
//// Получение максимального значения строки для потребностей Dialod API
//int   GetMaxLength() {return MaxLength;};
//
//void  SetInputMask(const wchar_t *InputMask);
//const wchar_t* GetInputMask() {return Mask;}
//
//void  SetOvertypeMode(int Mode) {Flags.Change(FEDITLINE_OVERTYPE,Mode);};
//int   GetOvertypeMode() {return Flags.Check(FEDITLINE_OVERTYPE);};
//
//void  SetConvertTabs(int Mode) { TabExpandMode = Mode;};
//int   GetConvertTabs() {return TabExpandMode;};
//
//int   RealPosToTab(int Pos);
//int   TabPosToReal(int Pos);
//void  Select(int Start,int End);
//void  AddSelect(int Start,int End);
//void  GetSelection(int &Start,int &End);
//BOOL  IsSelection() {return  SelStart==-1 && !SelEnd?FALSE:TRUE; };
//void  GetRealSelection(int &Start,int &End);
//void  SetEditBeyondEnd(int Mode) {Flags.Change(FEDITLINE_EDITBEYONDEND,Mode);};
//void  SetEditorMode(int Mode) {Flags.Change(FEDITLINE_EDITORMODE,Mode);};
//void  ReplaceTabs();
//
//void  InsertTab();
//
//void  AddColor(ColorItem *col);
//int   DeleteColor(int ColorPos);
//int   GetColor(ColorItem *col,int Item);
//
//void Xlat(bool All=false);
//
//void SetDialogParent(DWORD Sets);
//void SetCursorType(bool Visible, DWORD Size);
//void GetCursorType(bool& Visible, DWORD& Size);
//int  GetReadOnly() {return Flags.Check(FEDITLINE_READONLY);}
//void SetReadOnly(int NewReadOnly) {Flags.Change(FEDITLINE_READONLY,NewReadOnly);}
//int  GetDropDownBox() {return Flags.Check(FEDITLINE_DROPDOWNBOX);}
//void SetDropDownBox(int NewDropDownBox) {Flags.Change(FEDITLINE_DROPDOWNBOX,NewDropDownBox);}
//void SetWordDiv(const FARString& WordDiv) {strWordDiv=&WordDiv;}
//virtual void Changed(bool DelBlock=false);

type EditControl struct {
	Edit
	pCustomCompletionList Stack
	pHistory              *History
	pList                 *FarList
	Selection             bool
	SelectionStart        int
	ECFlags               BitFlags
	ACState               bool
}

//void SetMenuPos(VMenu& menu);
//int AutoCompleteProc(bool Manual,bool DelBlock,int& BackKey);
//void PopulateCompletionMenu(VMenu &ComplMenu, const FARString &strFilter);
//virtual int ProcessMouse(MOUSE_EVENT_RECORD *MouseEvent);
//virtual void Show();
//virtual void Changed(bool DelBlock=false);
//void SetCallbackState(bool Enable){m_Callback.Active=Enable;}
//
//void AutoComplete(bool Manual,bool DelBlock);
//void EnableAC(bool Permanent=false);
//void DisableAC(bool Permanent=false);
//void RevertAC(){ACState?EnableAC():DisableAC();}
//void ShowCustomCompletionList(const std::vector<std::string> &list);
