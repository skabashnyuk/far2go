package far2go

const FCMDOBJ_LOCKUPDATEPANEL = 0x00010000

type PushPopRecord struct {
	Name string
}

type CommandLine struct {
	ScreenObject
	CmdStr            EditControl
	BackgroundScreen  *SaveScreen
	CurDir            string
	LastCmdDir        string
	LastCmdPartLength uint
	LastKey           uint
	ppstack           Stack
}

func (obj *CommandLine) ProcessKey(key KeyCode) (bool) {
	return false
}

func (obj *CommandLine) ProcessMouse(MouseEvent *MouseEventRecord) (bool) {
	return false
}

func (obj *CommandLine) VMProcess(OpCode uint, vparam uint, iParam uint) (int) {
	return 0
}

func (obj *CommandLine) ResizeConsole() {
}
func (obj *CommandLine) GetConsoleLog() (string) {
	return nil
}

func (obj *CommandLine) GetCurDir() (string) {
	return obj.CurDir
}

func (obj *CommandLine) SetCurDir(CurDir string) {
	obj.CurDir = CurDir
}

func (obj *CommandLine) GetString() (string) {
	return obj.CmdStr.GetString()
}


//
//void GetString(FARString &strStr) { CmdStr.GetString(strStr); };
//int GetLength() { return CmdStr.GetLength(); };
//void SetString(const wchar_t *Str,BOOL Redraw=TRUE);
//void InsertString(const wchar_t *Str);
//
//void ExecString(const wchar_t *Str, bool AlwaysWaitFinish, bool SeparateWindow = false, bool DirectRun = false, bool WaitForIdle = false, bool Silent = false, bool RunAs = false);
//
//void ShowViewEditHistory();
//
//void SetCurPos(int Pos, int LeftPos=0);
//int GetCurPos() { return CmdStr.GetCurPos(); };
//int GetLeftPos() { return CmdStr.GetLeftPos(); };
//
//void SetPersistentBlocks(int Mode);
//void SetDelRemovesBlocks(int Mode);
//void SetAutoComplete(int Mode);
//
//void GetSelString(FARString &strStr) { CmdStr.GetSelString(strStr); };
//void GetSelection(int &Start,int &End) { CmdStr.GetSelection(Start,End); };
//void Select(int Start, int End) { CmdStr.Select(Start,End); };
//
//void SaveBackground(int X1,int Y1,int X2,int Y2);
//void SaveBackground();
//void ShowBackground();
//void CorrectRealScreenCoord();
//void LockUpdatePanel(int Mode) {Flags.Change(FCMDOBJ_LOCKUPDATEPANEL,Mode);};
//
//void EnableAC(){return CmdStr.EnableAC();}
//void DisableAC(){return CmdStr.DisableAC();}
//void RevertAC(){return CmdStr.RevertAC();}
