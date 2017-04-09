package far2go

const PANEL_COLUMNCOUNT = 20

type PanelViewSettings struct {
	ColumnType            [PANEL_COLUMNCOUNT] int
	ColumnWidth           [PANEL_COLUMNCOUNT] int
	ColumnCount           int
	StatusColumnType      [PANEL_COLUMNCOUNT]uint
	StatusColumnWidth     [PANEL_COLUMNCOUNT]int
	StatusColumnCount     int
	FullScreen            int
	AlignExtensions       int
	FolderAlignExtensions int
	FolderUpperCase       int
	FileLowerCase         int
	FileUpperToLowerCase  int
	ColumnWidthType       [PANEL_COLUMNCOUNT]int
	StatusColumnWidthType [PANEL_COLUMNCOUNT]int
}

type PanelType uint

const (
	FILE_PANEL  PanelType = iota
	TREE_PANEL
	QVIEW_PANEL
	INFO_PANEL
)

type SortMode uint

const (
	UNSORTED          SortMode = iota
	BY_NAME
	BY_EXT
	BY_MTIME
	BY_CTIME
	BY_ATIME
	BY_SIZE
	BY_DIZ
	BY_OWNER
	BY_COMPRESSEDSIZE
	BY_NUMLINKS
	BY_NUMSTREAMS
	BY_STREAMSSIZE
	BY_FULLNAME
	BY_CHTIME
	BY_CUSTOMDATA
)

const (
	VIEW_0 = iota
	VIEW_1
	VIEW_2
	VIEW_3
	VIEW_4
	VIEW_5
	VIEW_6
	VIEW_7
	VIEW_8
	VIEW_9
)

const (
	DRIVE_SHOW_TYPE       = 1 << iota
	DRIVE_SHOW_NETNAME
	DRIVE_SHOW_LABEL
	DRIVE_SHOW_FILESYSTEM
	DRIVE_SHOW_SIZE
	DRIVE_SHOW_REMOVABLE
	DRIVE_SHOW_PLUGINS
	DRIVE_SHOW_CDROM
	DRIVE_SHOW_SIZE_FLOAT
	DRIVE_SHOW_REMOTE
)

const (
	UPDATE_KEEP_SELECTION  = 1 << iota
	UPDATE_SECONDARY
	UPDATE_IGNORE_VISIBLE
	UPDATE_DRAW_MESSAGE
	UPDATE_CAN_BE_ANNOYING
)

const (
	NORMAL_PANEL = iota
	PLUGIN_PANEL
)

const (
	DRIVE_DEL_FAIL    = iota
	DRIVE_DEL_SUCCESS
	DRIVE_DEL_EJECT
	DRIVE_DEL_NONE
)

const (
	UIC_UPDATE_NORMAL             = iota
	UIC_UPDATE_FORCE
	UIC_UPDATE_FORCE_NOTIFICATION
)

type Panel struct {
	ScreenObject
	ViewSettings            PanelViewSettings
	ProcessingPluginCommand int
	curDir                  string
	focus                   int
	panelType               PanelType
	enableUpdate            int
	panelMode               int
	sortMode                SortMode
	sortOrder               int
	sortGroups              int
	prevViewMode, viewMode  int
	curTopFile              int
	curFile                 int
	numericSort             int
	caseSensitiveSort       int
	directoriesFirst        int
	modalMode               int
	pluginCommand           int
	pluginParam             int
}

//virtual int SendKeyToPlugin(DWORD Key,BOOL Pred=FALSE) {return FALSE;};
//virtual BOOL SetCurDir(const wchar_t *NewDir,int ClosePlugin);
//virtual void ChangeDirToCurrent();
//
//virtual int GetCurDir(FARString &strCurDir);
//
//virtual int GetSelCount() {return 0;};
//virtual int GetRealSelCount() {return 0;};
//virtual int GetSelName(FARString *strName,DWORD &FileAttr,DWORD &FileMode,FAR_FIND_DATA_EX *fd=nullptr) {return FALSE;};
//virtual void UngetSelName() {};
//virtual void ClearLastGetSelection() {};
//virtual uint64_t GetLastSelectedSize() {return (uint64_t)(-1);};
//virtual int GetLastSelectedItem(struct FileListItem *LastItem) {return 0;};
//
//virtual int GetCurName(FARString &strName);
//virtual int GetCurBaseName(FARString &strName);
//virtual int GetFileName(FARString &strName,int Pos,DWORD &FileAttr) {return FALSE;};
//
//virtual int GetCurrentPos() {return 0;};
//virtual void SetFocus();
//virtual void KillFocus();
//virtual void Update(int Mode) {};
//virtual int UpdateIfChanged(int UpdateMode) {return 0;};
//virtual void UpdateIfRequired() {};
//
//virtual void CloseChangeNotification() {};
//virtual int FindPartName(const wchar_t *Name,int Next,int Direct=1,int ExcludeSets=0) {return FALSE;}
//
//virtual int GoToFile(long idxItem) {return TRUE;};
//virtual int GoToFile(const wchar_t *Name,BOOL OnlyPartName=FALSE) {return TRUE;};
//virtual long FindFile(const wchar_t *Name,BOOL OnlyPartName=FALSE) {return -1;};
//
//virtual int IsSelected(const wchar_t *Name) {return FALSE;};
//virtual int IsSelected(long indItem) {return FALSE;};
//
//virtual long FindFirst(const wchar_t *Name) {return -1;}
//virtual long FindNext(int StartPos, const wchar_t *Name) {return -1;}
//
//virtual void SetSelectedFirstMode(int) {};
//virtual int GetSelectedFirstMode() {return 0;};
//int GetMode() {return(PanelMode);};
//void SetMode(int Mode) {PanelMode=Mode;};
//int GetModalMode() {return(ModalMode);};
//void SetModalMode(int ModalMode) {Panel::ModalMode=ModalMode;};
//int GetViewMode() {return(ViewMode);};
//virtual void SetViewMode(int ViewMode);
//virtual int GetPrevViewMode() {return(PrevViewMode);};
//void SetPrevViewMode(int PrevViewMode) {Panel::PrevViewMode=PrevViewMode;};
//virtual int GetPrevSortMode() {return(SortMode);};
//virtual int GetPrevSortOrder() {return(SortOrder);};
//int GetSortMode() {return(SortMode);};
//virtual int GetPrevNumericSort() {return NumericSort;};
//int GetNumericSort() { return NumericSort; }
//void SetNumericSort(int Mode) { NumericSort=Mode; }
//virtual void ChangeNumericSort(int Mode) { SetNumericSort(Mode); }
//virtual int GetPrevCaseSensitiveSort() {return CaseSensitiveSort;}
//int GetCaseSensitiveSort() {return CaseSensitiveSort;}
//void SetCaseSensitiveSort(int Mode) {CaseSensitiveSort = Mode;}
//virtual void ChangeCaseSensitiveSort(int Mode) {SetCaseSensitiveSort(Mode);}
//virtual int GetPrevDirectoriesFirst() {return DirectoriesFirst;};
//int GetDirectoriesFirst() { return DirectoriesFirst; }
//void SetDirectoriesFirst(int Mode) { DirectoriesFirst=Mode; }
//virtual void ChangeDirectoriesFirst(int Mode) { SetDirectoriesFirst(Mode); }
//virtual void SetSortMode(int SortMode) {Panel::SortMode=SortMode;};
//int GetSortOrder() {return(SortOrder);};
//void SetSortOrder(int SortOrder) {Panel::SortOrder=SortOrder;};
//virtual void ChangeSortOrder(int NewOrder) {SetSortOrder(NewOrder);};
//int GetSortGroups() {return(SortGroups);};
//void SetSortGroups(int SortGroups) {Panel::SortGroups=SortGroups;};
//void InitCurDir(const wchar_t *CurDir);
//virtual void CloseFile() {};
//virtual void UpdateViewPanel() {};
//virtual void CompareDir() {};
//virtual void MoveToMouse(MOUSE_EVENT_RECORD *MouseEvent) {};
//virtual void ClearSelection() {};
//virtual void SaveSelection() {};
//virtual void RestoreSelection() {};
//virtual void SortFileList(int KeepPosition) {};
//virtual void EditFilter() {};
//virtual bool FileInFilter(long idxItem) {return true;};
//virtual void ReadDiz(struct PluginPanelItem *ItemList=nullptr,int ItemLength=0, DWORD dwFlags=0) {};
//virtual void DeleteDiz(const wchar_t *Name) {};
//virtual void GetDizName(FARString &strDizName) {};
//virtual void FlushDiz() {};
//virtual void CopyDiz(const wchar_t *Name,const wchar_t *DestName, DizList *DestDiz) {};
//virtual int IsFullScreen() {return ViewSettings.FullScreen;};
//virtual int IsDizDisplayed() {return FALSE;};
//virtual int IsColumnDisplayed(int Type) {return FALSE;};
//virtual int GetColumnsCount() { return 1;};
//virtual void SetReturnCurrentFile(int Mode) {};
//virtual void QViewDelTempName() {};
//virtual void GetOpenPluginInfo(struct OpenPluginInfo *Info) {};
//virtual void SetPluginMode(HANDLE hPlugin,const wchar_t *PluginFile,bool SendOnFocus=false) {};
//virtual void SetPluginModified() {};
//virtual int ProcessPluginEvent(int Event,void *Param) {return FALSE;};
//virtual HANDLE GetPluginHandle() {return(INVALID_HANDLE_VALUE);};
//virtual void SetTitle();
//virtual FARString &GetTitle(FARString &Title,int SubLen=-1,int TruncSize=0);
//
//virtual int64_t VMProcess(int OpCode,void *vParam=nullptr,int64_t iParam=0);
//
//virtual void IfGoHome(wchar_t Drive) {};
//virtual BOOL UpdateKeyBar() { return FALSE; };
//
//virtual long GetFileCount() {return 0;}
//virtual BOOL GetItem(int,void *) {return FALSE;};
//
//bool ExecShortcutFolder(int Pos);
//bool SaveShortcutFolder(int Pos);
//
//static void EndDrag();
//virtual void Hide();
//virtual void Show();
//int SetPluginCommand(int Command,int Param1,LONG_PTR Param2);
//int PanelProcessMouse(MOUSE_EVENT_RECORD *MouseEvent,int &RetCode);
//void ChangeDisk();
//int GetFocus() {return(Focus);};
//int GetType() {return(Type);};
//void SetUpdateMode(int Mode) {EnableUpdate=Mode;};
//bool MakeListFile(FARString &strListFileName,const wchar_t *Modifers=nullptr);
//int SetCurPath();
//
//BOOL NeedUpdatePanel(Panel *AnotherPanel);
