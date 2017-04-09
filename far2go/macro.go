package far2go

type MacroFunctionType struct {
	Name       string
	nParam     uint
	oParam     uint
	Code       int
	fnGUID     string
	BufferSize int
	Buffer     uint
	Syntax     string
	IntFlags   uint
	Func       interface{}
}
type MacroRecord struct {
	Flags       uint
	Key         int
	BufferSize  int
	Buffer      uint
	Src         string
	Description string
	Reserved    []uint
}
type MacroState struct {
	KeyProcess           int
	Executing            int
	MacroPC              int
	ExecLIBPos           int
	MacroWORKCount       int
	UseInternalClipboard bool
	MacroWORK            *MacroRecord
	cRec                 InputRecord
	AllocVarTable        bool
	//locVarTable          *TVarTable
}

const STACKLEVEL = 32

const (
	MACRO_FUNCS  = -3
	MACRO_CONSTS = -2
	MACRO_VARS   = -1

	MACRO_OTHER          = 0
	MACRO_SHELL          = 1
	MACRO_VIEWER         = 2
	MACRO_EDITOR         = 3
	MACRO_DIALOG         = 4
	MACRO_SEARCH         = 5
	MACRO_DISKS          = 6
	MACRO_MAINMENU       = 7
	MACRO_MENU           = 8
	MACRO_HELP           = 9
	MACRO_INFOPANEL      = 10
	MACRO_QVIEWPANEL     = 11
	MACRO_TREEPANEL      = 12
	MACRO_FINDFOLDER     = 13
	MACRO_USERMENU       = 14
	MACRO_AUTOCOMPLETION = 15

	MACRO_COMMON
	MACRO_LAST
)

type KeyMacro struct {
	MacroVersion       uint64
	LastOpCodeUF       uint64
	CMacroFunction     uint
	AllocatedFuncCount uint
	AMacroFunction     *MacroFunctionType
	Recording          int
	InternalInput      int
	IsRedrawEditor     int
	Mode               int
	StartMode          int
	Work               MacroState
	PCStack            [STACKLEVEL]MacroState
	CurPCStack         int
	MacroLIBCount      int
	MacroLIB           *MacroRecord
	IndexMode          [MACRO_LAST][2]int
	RecBufferSize      int
	RecBuffer          uint
	RecSrc             string
	//LockScr            *LockScreen
}

//
//int ReadVarsConst(int ReadMode, FARString &strBuffer);
//int ReadMacroFunction(int ReadMode, FARString &strBuffer);
//int WriteVarsConst(int WriteMode);
//int ReadMacros(int ReadMode, FARString &strBuffer);
//DWORD AssignMacroKey();
//int GetMacroSettings(int Key,DWORD &Flags);
//void InitInternalVars(BOOL InitedRAM=TRUE);
//void InitInternalLIBVars();
//void ReleaseWORKBuffer(BOOL All=FALSE); // удалить временный буфер
//
//DWORD SwitchFlags(DWORD& Flags,DWORD Value);
//FARString &MkRegKeyName(int IdxMacro,FARString &strRegKeyName);
//
//BOOL CheckEditSelected(DWORD CurFlags);
//BOOL CheckInsidePlugin(DWORD CurFlags);
//BOOL CheckPanel(int PanelMode,DWORD CurFlags, BOOL IsPassivePanel);
//BOOL CheckCmdLine(int CmdLength,DWORD Flags);
//BOOL CheckFileFolder(Panel *ActivePanel,DWORD CurFlags, BOOL IsPassivePanel);
//BOOL CheckAll(int CheckMode,DWORD CurFlags);
//void Sort();
//TVar FARPseudoVariable(DWORD Flags,DWORD Code,DWORD& Err);
//DWORD GetOpCode(struct MacroRecord *MR,int PC);
//DWORD SetOpCode(struct MacroRecord *MR,int PC,DWORD OpCode);

//
//static LONG_PTR WINAPI AssignMacroDlgProc(HANDLE hDlg,int Msg,int Param1,LONG_PTR Param2);
//static LONG_PTR WINAPI ParamMacroDlgProc(HANDLE hDlg,int Msg,int Param1,LONG_PTR Param2);

//int ProcessKey(int Key);
//int GetKey();
//int PeekKey();
//bool IsOpCode(DWORD p);
//bool CheckWaitKeyFunc();
//
//int PushState(bool CopyLocalVars=FALSE);
//int PopState();
//int GetLevelState() {return CurPCStack;};
//
//int  IsRecording() {return(Recording);};
//int  IsExecuting() {return(Work.Executing);};
//int  IsExecutingLastKey();
//int  IsDsableOutput() {return CheckCurMacroFlags(MFLAGS_DISABLEOUTPUT);};
//void SetMode(int Mode) {KeyMacro::Mode=Mode;};
//int  GetMode() {return(Mode);};
//
//void DropProcess();
//
//void RunStartMacro();
//
//// Поместить временное строковое представление макроса
//int PostNewMacro(const wchar_t *PlainText,DWORD Flags=0,DWORD AKey=0,BOOL onlyCheck=FALSE);
//// Поместить временный рекорд (бинарное представление)
//int PostNewMacro(struct MacroRecord *MRec,BOOL NeedAddSendFlag=0,BOOL IsPluginSend=FALSE);
//
//int  LoadMacros(BOOL InitedRAM=TRUE,BOOL LoadAll=TRUE);
//void SaveMacros(BOOL AllSaved=TRUE);
//
//int GetStartIndex(int Mode) {return IndexMode[Mode<MACRO_LAST-1?Mode:MACRO_LAST-1][0];}
//// Функция получения индекса нужного макроса в массиве
//int GetIndex(int Key, int Mode, bool UseCommon=true);
//// получение размера, занимаемого указанным макросом
//int GetRecordSize(int Key, int Mode);
//
//bool GetPlainText(FARString& Dest);
//int  GetPlainTextSize();
//
//void SetRedrawEditor(int Sets) {IsRedrawEditor=Sets;}
//
//void RestartAutoMacro(int Mode);
//
//// получить данные о макросе (возвращает статус)
//int GetCurRecord(struct MacroRecord* RBuf=nullptr,int *KeyPos=nullptr);
//// проверить флаги текущего исполняемого макроса.
//BOOL CheckCurMacroFlags(DWORD Flags);
//
//static const wchar_t* GetSubKey(int Mode);
//static int   GetSubKey(const wchar_t *Mode);
//static int   GetMacroKeyInfo(bool FromReg,int Mode,int Pos,FARString &strKeyName,FARString &strDescription);
//static wchar_t *MkTextSequence(DWORD *Buffer,int BufferSize,const wchar_t *Src=nullptr);
//// из строкового представления макроса сделать MacroRecord
//int ParseMacroString(struct MacroRecord *CurMacro,const wchar_t *BufPtr,BOOL onlyCheck=FALSE);
//BOOL GetMacroParseError(DWORD* ErrCode, COORD* ErrPos, FARString *ErrSrc);
//BOOL GetMacroParseError(FARString *Err1, FARString *Err2, FARString *Err3, FARString *Err4);
//
//static void SetMacroConst(const wchar_t *ConstName, const TVar Value);
//static DWORD GetNewOpCode();
//
//static size_t GetCountMacroFunction();
//static const TMacroFunction *GetMacroFunction(size_t Index);
//static void RegisterMacroIntFunction();
//static TMacroFunction *RegisterMacroFunction(const TMacroFunction *tmfunc);
//static bool UnregMacroFunction(size_t Index);

//BOOL WINAPI KeyMacroToText(int Key,FARString &strKeyText0);
//int WINAPI KeyNameMacroToKey(const wchar_t *Name);
//void initMacroVarTable(int global);
//void doneMacroVarTable(int global);
//bool checkMacroConst(const wchar_t *name);
//const wchar_t *eStackAsString(int Pos=0);
//
//inline bool IsMenuArea(int Area){return Area==MACRO_MAINMENU || Area==MACRO_MENU || Area==MACRO_DISKS || Area==MACRO_USERMENU || Area==MACRO_AUTOCOMPLETION;}
