package far2go

const FCMDOBJ_LOCKUPDATEPANEL = 0x00010000

type PushPopRecord struct {
	Name string
}

type CommandLine struct {
	ScreenObject
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


//virtual int64_t VMProcess(int OpCode,void *vParam=nullptr,int64_t iParam=0);
//
//virtual void ResizeConsole();
//
//std::string GetConsoleLog();
//int GetCurDir(FARString &strCurDir);
//BOOL SetCurDir(const wchar_t *CurDir);
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
