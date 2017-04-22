package far2go

type FilePanels struct {
	Frame
	LastLeftFilePanel,
	LastRightFilePanel,
	LeftPanel,
	RightPanel,
	ActivePanel *Panel

	MainKeyBar                  KeyBar
	TopMenuBar                  MenuBar
	LastLeftType, LastRightType int

	LeftStateBeforeHide, RightStateBeforeHide int
}

func newFilePanels() *FilePanels {
	return &FilePanels{}
}

//void Init();
//
//Panel* CreatePanel(int Type);
//void   DeletePanel(Panel *Deleted);
//Panel* GetAnotherPanel(Panel *Current);
//Panel* ChangePanelToFilled(Panel *Current,int NewType);
//Panel* ChangePanel(Panel *Current,int NewType,int CreateNew,int Force);
//void   SetPanelPositions(int LeftFullScreen,int RightFullScreen);
////    void   SetPanelPositions();
//
//void   SetupKeyBar();
//
//virtual int ProcessKey(int Key);
//virtual int ProcessMouse(MOUSE_EVENT_RECORD *MouseEvent);
//virtual int64_t VMProcess(int OpCode,void *vParam=nullptr,int64_t iParam=0);
//
//int SetAnhoterPanelFocus();
//int SwapPanels();
//int ChangePanelViewMode(Panel *Current,int Mode,BOOL RefreshFrame);
//
//virtual void SetScreenPosition();
//
//void Update();
//
//virtual int GetTypeAndName(FARString &strType, FARString &strName);
//virtual int GetType() { return MODALTYPE_PANELS; }
//virtual const wchar_t *GetTypeName() {return L"[FilePanels]";};
//
//virtual void OnChangeFocus(int focus);
//
//virtual void RedrawKeyBar(); // virtual
//virtual void ShowConsoleTitle();
//virtual void ResizeConsole();
//virtual int FastHide();
//virtual void Refresh();
//void GoToFile(const wchar_t *FileName);
//
//virtual int GetMacroMode();
