package far2go

import (
	"github.com/spf13/viper"
	"fmt"
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

const (
	WordDiv0 = "~!%^&*()+|{}:\"<>?`-=\\[];',./"

	WordDivForXlat0 = " \t!#$%^&*()+|=\\/@?"

	//FARString strKeyNameConsoleDetachKey;
	szCtrlShiftX   = "CtrlShiftX"
	szCtrlDot      = "Ctrl."
	szCtrlShiftDot = "CtrlShift."

	// KeyName
	NKeyColors              = "Colors"
	NKeyScreen              = "Screen"
	NKeyCmdline             = "Cmdline"
	NKeyInterface           = "Interface"
	NKeyInterfaceCompletion = "Interface/Completion"
	NKeyViewer              = "Viewer"
	NKeyDialog              = "Dialog"
	NKeyEditor              = "Editor"
	NKeyXLat                = "XLat"
	NKeySystem              = "System"
	NKeySystemExecutor      = "System/Executor"
	NKeySystemNowell        = "System/Nowell"
	NKeyHelp                = "Help"
	NKeyLanguage            = "Language"
	NKeyConfirmations       = "Confirmations"
	NKeyPluginConfirmations = "PluginConfirmations"
	NKeyPanel               = "Panel"
	NKeyPanelLeft           = "Panel/Left"
	NKeyPanelRight          = "Panel/Right"
	NKeyPanelLayout         = "Panel/Layout"
	NKeyPanelTree           = "Panel/Tree"
	NKeyPanelInfo           = "Panel/Info"
	NKeyLayout              = "Layout"
	NKeyDescriptions        = "Descriptions"
	NKeyKeyMacros           = "KeyMacros"
	NKeyPolicies            = "Policies"
	NKeyFileFilter          = "OperationsFilter"
	NKeySavedHistory        = "SavedHistory"
	NKeySavedViewHistory    = "SavedViewHistory"
	NKeySavedFolderHistory  = "SavedFolderHistory"
	NKeySavedDialogHistory  = "SavedDialogHistory"
	NKeyCodePages           = "CodePages"
	NParamHistoryCount      = "HistoryCount"
	NKeyVMenu               = "VMenu"

	constBatchExt = ".BAT;.CMD;"
)
const (
	CASR_PANEL  = 0x0001
	CASR_EDITOR = 0x0002
	CASR_VIEWER = 0x0004
	CASR_HELP   = 0x0008
	CASR_DIALOG = 0x0010
)
const CP_AUTODETECT = -1

type ExcludeCmdHistoryType uint

const (
	EXCLUDECMDHISTORY_NOTWINASS  ExcludeCmdHistoryType = 0x00000001
	EXCLUDECMDHISTORY_NOTFARASS  ExcludeCmdHistoryType = 0x00000002
	EXCLUDECMDHISTORY_NOTPANEL   ExcludeCmdHistoryType = 0x00000004
	EXCLUDECMDHISTORY_NOTCMDLINE ExcludeCmdHistoryType = 0x00000008
	//EXCLUDECMDHISTORY_NOTAPPLYCMD   = 0x00000010
)

type QuoteDNameType uint

const (
	QUOTEDNAME_INSERT    QuoteDNameType = 0x00000001
	QUOTEDNAME_CLIPBOARD QuoteDNameType = 0x00000002
)
const DMOUSEBUTTON_LEFT = 0x00000001
const DMOUSEBUTTON_RIGHT = 0x00000002

const (
	VMENUCLICK_IGNORE = iota
	VMENUCLICK_CANCEL
	VMENUCLICK_APPLY
)

type DizUpdateType uint

const (
	DIZ_NOT_UPDATE          DizUpdateType = iota
	DIZ_UPDATE_IF_DISPLAYED
	DIZ_UPDATE_ALWAYS
)

type PanelOptions struct {
	Type              int
	Visible           int
	Focus             int
	ViewMode          int
	SortMode          int
	SortOrder         int
	SortGroups        int
	NumericSort       int
	CaseSensitiveSort int
	DirectoriesFirst  int
}

type AutoCompleteOptions struct {
	ShowList         int
	ModalList        int
	AppendCompletion int
	Exceptions       string
}

type PluginConfirmation struct {
	OpenFilePlugin      int
	StandardAssociation int
	EvenIfOnlyOnePlugin int
	SetFindList         int
	Prefix              int
}

type Confirmation struct {
	Copy                int
	Move                int
	RO                  int
	Drag                int
	Delete              int
	DeleteFolder        int
	Exit                int
	Esc                 int
	EscTwiceToInterrupt int
	RemoveConnection    int
	AllowReedit         int
	HistoryClear        int
	RemoveSUBST         int
	RemoveHotPlug       int
	DetachVHD           int
}

type DizOptions struct {
	strListNames  string
	ROUpdate      int
	UpdateMode    DizUpdateType
	SetHidden     int
	StartPos      int
	AnsiByDefault int
	SaveInUTF     int
}

type CodeXLAT struct {
	Flags             uint64
	strWordDivForXlat string
	Table             []string
	Rules             []string
}

type EditorOptions struct {
	TabSize                 int
	ExpandTabs              int
	PersistentBlocks        int
	DelRemovesBlocks        int
	AutoIndent              int
	AutoDetectCodePage      int
	UTF8CodePageForNewFile  int
	UTF8CodePageAsDefault   int
	CursorBeyondEOL         int
	BSLikeDel               int
	CharCodeBase            int
	SavePos                 int
	SaveShortPos            int
	F7Rules                 int
	AllowEmptySpaceAfterEof int
	ReadOnlyLock            int
	UndoSize                int
	UseExternalEditor       int
	FileSizeLimitLo         uint64
	FileSizeLimitHi         uint64
	ShowKeyBar              int
	ShowTitleBar            int
	ShowScrollBar           int
	EditOpenedForWrite      int
	SearchSelFound          int
	SearchRegexp            int
	SearchPickUpWord        int
	ShowWhiteSpace          int
	strWordDiv              string
}

type ViewerOptions struct {
	TabSize               int
	AutoDetectCodePage    int
	ShowScrollbar         int
	ShowArrows            int
	PersistentBlocks      int
	ViewerIsWrap          int
	ViewerWrap            int
	SavePos               int
	SaveShortPos          int
	UseExternalViewer     int
	ShowKeyBar            int
	UTF8CodePageAsDefault int
	ShowTitleBar          int
	SearchRegexp          int
}

type PoliciesOptions struct {
	DisabledOptions  int
	ShowHiddenDrives int
}

type DialogsOptions struct {
	EditBlock         int
	EditHistory       int
	AutoComplete      int
	EULBsClear        int
	SelectFromHistory int
	EditLine          int
	MouseButton       int
	DelRemovesBlocks  int
	CBoxMaxHeight     int
}

type VMenuOptions struct {
	LBtnClick int
	RBtnClick int
	MBtnClick int
}

type CommandLineOptions struct {
	EditBlock        int
	DelRemovesBlocks int
	AutoComplete     int
	UsePromptFormat  int
	strPromptFormat  string
}

type NowellOptions struct {
	MoveRO int
}

type ScreenSizes struct {
	DeltaXY        Coordinate
	WScreenSizeSet int
	WScreenSize    []Coordinate
}

type LoadPluginsOptions struct {
	MainPluginDir          int
	PluginsCacheOnly       int
	PluginsPersonal        int
	strCustomPluginsPath   string
	strPersonalPluginsPath string
	SilentLoadPlugin       int
	OEMPluginsSupport      int
	ScanSymlinks           int
}

type FindFileOptions struct {
	FileSearchMode          int
	FindFolders             bool
	FindSymLinks            bool
	CollectFiles            bool
	UseFilter               bool
	FindAlternateStreams    bool
	strSearchInFirstSize    string
	strSearchOutFormat      string
	strSearchOutFormatWidth string
	OutColumnCount          int
	OutColumnTypes          []uint
	OutColumnWidths         []int
	OutColumnWidthType      []int
}

type InfoPanelOptions struct {
	strFolderInfoFiles string
};

type TreeOptions struct {
	LocalDisk        int
	NetDisk          int
	NetPath          int
	RemovableDisk    int
	MinTreeCount     int
	AutoChangeFolder int
	TreeFileAttr     int
};

type CopyMoveOptions struct {
	WriteThrough   int
	CopyXAttr      int
	CopyAccessMode int
	CopyOpened     int
	CopyShowTotal  int
	MultiCopy      int
	CopyTimeRule   int
}

type DeleteOptions struct {
	DelShowTotal int
}

type MacroOptions struct {
	MacroReuseRules      int
	DisableMacro         int
	KeyMacroCtrlDot      int
	KeyMacroCtrlShiftDot int
	CallPluginRules      int
	strMacroCONVFMT      string
	strDateFormat        string
}

type Options struct {
	Clock                      int
	Mouse                      int
	ShowKeyBar                 int
	ScreenSaver                int
	ScreenSaverTime            int
	UseVk_oem_x                int
	InactivityExit             int
	InactivityExitTime         int
	ShowHidden                 int
	Highlight                  int
	strLeftFolder              string
	strRightFolder             string
	strLeftCurFile             string
	strRightCurFile            string
	RightSelectedFirst         int
	LeftSelectedFirst          int
	SelectFolders              int
	ReverseSort                int
	SortFolderExt              int
	DeleteToRecycleBin         int
	DeleteToRecycleBinKillLink int
	WipeSymbol                 int
	SudoEnabled                int
	SudoConfirmModify          int
	SudoPasswordExpiration     int
	CMOpt                      CopyMoveOptions
	DelOpt                     DeleteOptions
	MultiMakeDir               int

	UseRegisteredTypes       int
	ViewerEditorClock        int
	OnlyEditorViewerUsed     int
	SaveViewHistory          int
	ViewHistoryCount         int
	strExternalEditor        string
	EdOpt                    EditorOptions
	strExternalViewer        string
	ViOpt                    ViewerOptions
	strWordDiv               string
	strQuotedSymbols         string
	QuotedName               uint64
	AutoSaveSetup            int
	SetupArgv                int
	ChangeDriveMode          int
	ChangeDriveDisconnetMode int
	SaveHistory              int
	HistoryCount             int
	SaveFoldersHistory       int
	SavePluginFoldersHistory int
	FoldersHistoryCount      int
	DialogsHistoryCount      int
	FindOpt                  FindFileOptions
	LeftHeightDecrement      int
	RightHeightDecrement     int
	WidthDecrement           int
	ShowColumnTitles         int
	ShowPanelStatus          int
	ShowPanelTotals          int
	ShowPanelFree            int
	ShowPanelScrollbar       int
	ShowMenuScrollbar        int
	ShowScreensNumber        int
	ShowSortMode             int
	ShowMenuBar              int
	FormatNumberSeparators   int
	CleanAscii               int
	NoGraphics               int

	ConsolePaintSharp  int
	ExclusiveCtrlLeft  int
	ExclusiveCtrlRight int
	ExclusiveAltLeft   int
	ExclusiveAltRight  int
	ExclusiveWinLeft   int
	ExclusiveWinRight  int

	Confirm       Confirmation
	PluginConfirm PluginConfirmation

	Diz DizOptions

	ShellRightLeftArrowsRule int
	LeftPanel                PanelOptions
	RightPanel               PanelOptions

	AutoComplete AutoCompleteOptions

	AutoUpdateLimit       uint64
	AutoUpdateRemoteDrive int

	strLanguage           string
	SmallIcon             int
	strRegRoot            string
	PanelRightClickRule   int
	PanelCtrlAltShiftRule int
	PanelCtrlFRule        int
	AllCtrlAltShiftRulev  int

	CASRule int

	CmdHistoryRule int

	ExcludeCmdHistory  uint64
	SubstPluginPrefix  int
	MaxPositionCache   int
	SetAttrFolderRules int
	ExceptRules        int
	ExceptCallDebugger int

	ShiftsKeyRules int
	CursorSize     []int

	XLat             CodeXLAT
	ConsoleDetachKey int

	UsePrintManager int
	strHelpLanguage string
	FullScreenHelp  int
	HelpTabSize     int
	HelpURLRules    int

	RememberLogicalDrives int

	MsWheelDelta     int
	MsWheelDeltaView int
	MsWheelDeltaEdit int
	MsWheelDeltaHelp int

	MsHWheelDelta     int
	MsHWheelDeltaView int
	MsHWheelDeltaEdit int

	SubstNameRule           int
	PgUpChangeDisk          int
	ShowCheckingFile        int
	CloseConsoleRule        int
	CloseCDGate             int
	UpdateEnvironment       int
	LCIDSort                uint64
	RestoreCPAfterExecute   int
	ExecuteShowErrorMessage int
	ExecuteUseAppPath       int
	ExecuteFullTitle        int
	ExecuteSilentExternal   int
	strExecuteBatchType     string
	PluginMaxReadData       uint64
	UseNumPad               int
	ScanJunction            int
	OnlyFilesSize           int
	ShowTimeoutDelFiles     uint64
	ShowTimeoutDACLFiles    uint64
	DelThreadPriority       int

	LoadPlug       LoadPluginsOptions
	Dialogs        DialogsOptions
	VMenu          VMenuOptions
	CmdLine        CommandLineOptions
	Policies       PoliciesOptions
	Nowell         NowellOptions
	ScrSize        ScreenSizes
	Macro          MacroOptions
	FindCodePage   int
	Tree           TreeOptions
	InfoPanel      InfoPanelOptions
	CPMenuMode     uint64
	IsUserAdmin    bool
	strTitleAddons string
	WindowMode     bool
}

var Opt Options

func ReadConfig() {

	viper.SetConfigType("json")
	viper.SetConfigName("far2go")        // name of config file (without extension)
	viper.AddConfigPath("$HOME/.far2go") // call multiple times to add many search paths
	//	viper.AddConfigPath(".")             // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file

	if err == nil {
		// unmarshal config
		err = viper.Unmarshal(&Opt)
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal to Unmarshal: %s \n", err))
		}
	} else {
		switch  err.(type) {
		case viper.ConfigFileNotFoundError:
			Opt = *DefaultOptions()
			break
		default:
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
			break

		}
	}

}
func init() {
	ReadConfig()
}

func SaveConfig() {

	hdir, err := homedir.Dir()
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	var dir = filepath.Join(hdir, ".far2go")
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dir, 0755)
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Println(err)
		}
	}

	path := filepath.Join(dir, "far2go.json")
	os.Remove(path)
	b, err := json.MarshalIndent(Opt, "", "    ")
	if err != nil {
		log.Println(err)
	}

	ioutil.WriteFile(path, b, 0644)
}

func DefaultOptions() *Options {
	var defaultOptions = Options{}

	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//{1, REG_BINARY, NKeyColors, L"CurrentPalette",(char*)Palette,(DWORD)SizeArrayPalette,(wchar_t*)DefaultPalette},
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//
	//{1, REG_DWORD,  NKeyScreen, L"Clock", &Opt.Clock, 1, 0},
	defaultOptions.Clock = 1
	//{1, REG_DWORD,  NKeyScreen, L"ViewerEditorClock",&Opt.ViewerEditorClock,0, 0},
	defaultOptions.ViewerEditorClock = 0
	//{1, REG_DWORD,  NKeyScreen, L"KeyBar",&Opt.ShowKeyBar,1, 0},
	defaultOptions.ShowKeyBar = 1
	//{1, REG_DWORD,  NKeyScreen, L"ScreenSaver",&Opt.ScreenSaver, 0, 0},
	defaultOptions.ScreenSaver = 0
	//{1, REG_DWORD,  NKeyScreen, L"ScreenSaverTime",&Opt.ScreenSaverTime,5, 0},
	defaultOptions.ScreenSaverTime = 5
	//{0, REG_DWORD,  NKeyScreen, L"DeltaXY", &Opt.ScrSize.DeltaXY, 0, 0},
	defaultOptions.ScrSize = ScreenSizes{}
	//
	//{1, REG_DWORD,  NKeyCmdline, L"UsePromptFormat", &Opt.CmdLine.UsePromptFormat,0, 0},
	//{1, REG_SZ,     NKeyCmdline, L"PromptFormat",&Opt.CmdLine.strPromptFormat, 0, L"$p$# "},
	//{1, REG_DWORD,  NKeyCmdline, L"DelRemovesBlocks", &Opt.CmdLine.DelRemovesBlocks,1, 0},
	//{1, REG_DWORD,  NKeyCmdline, L"EditBlock", &Opt.CmdLine.EditBlock,0, 0},
	//{1, REG_DWORD,  NKeyCmdline, L"AutoComplete",&Opt.CmdLine.AutoComplete,1, 0},
	defaultOptions.CmdLine = CommandLineOptions{
		EditBlock:        0,
		DelRemovesBlocks: 1,
		AutoComplete:     1,
		UsePromptFormat:  0,
		strPromptFormat:  "$p$# ",
	}

	//
	//
	//{1, REG_DWORD,  NKeyInterface, L"Mouse",&Opt.Mouse,1, 0},
	defaultOptions.Mouse = 1
	//{0, REG_DWORD,  NKeyInterface, L"UseVk_oem_x",&Opt.UseVk_oem_x,1, 0},
	defaultOptions.UseVk_oem_x = 1
	//{1, REG_DWORD,  NKeyInterface, L"ShowMenuBar",&Opt.ShowMenuBar,0, 0},
	defaultOptions.ShowMenuBar = 0
	//{0, REG_DWORD,  NKeyInterface, L"CursorSize1",&Opt.CursorSize[0],15, 0},
	//{0, REG_DWORD,  NKeyInterface, L"CursorSize2",&Opt.CursorSize[1],10, 0},
	//{0, REG_DWORD,  NKeyInterface, L"CursorSize3",&Opt.CursorSize[2],99, 0},
	//{0, REG_DWORD,  NKeyInterface, L"CursorSize4",&Opt.CursorSize[3],99, 0},
	defaultOptions.CursorSize = []int{15, 10, 99, 99}
	//{0, REG_DWORD,  NKeyInterface, L"ShiftsKeyRules",&Opt.ShiftsKeyRules,1, 0},
	defaultOptions.ShiftsKeyRules = 1
	//{1, REG_DWORD,  NKeyInterface, L"CtrlPgUp",&Opt.PgUpChangeDisk, 1, 0},
	defaultOptions.PgUpChangeDisk = 1
	//
	//{1, REG_DWORD,  NKeyInterface, L"ConsolePaintSharp",&Opt.ConsolePaintSharp, 0, 0},
	defaultOptions.ConsolePaintSharp = 0
	//{1, REG_DWORD,  NKeyInterface, L"ExclusiveCtrlLeft",&Opt.ExclusiveCtrlLeft, 0, 0},
	defaultOptions.ExclusiveCtrlLeft = 0
	//{1, REG_DWORD,  NKeyInterface, L"ExclusiveCtrlRight",&Opt.ExclusiveCtrlRight, 0, 0},
	defaultOptions.ExclusiveCtrlRight = 0
	//{1, REG_DWORD,  NKeyInterface, L"ExclusiveAltLeft",&Opt.ExclusiveAltLeft, 0, 0},
	defaultOptions.ExclusiveAltLeft = 0
	//{1, REG_DWORD,  NKeyInterface, L"ExclusiveAltRight",&Opt.ExclusiveAltRight, 0, 0},
	defaultOptions.ExclusiveAltRight = 0
	//{1, REG_DWORD,  NKeyInterface, L"ExclusiveWinLeft",&Opt.ExclusiveWinLeft, 0, 0},
	defaultOptions.ExclusiveWinLeft = 0
	//{1, REG_DWORD,  NKeyInterface, L"ExclusiveWinRight",&Opt.ExclusiveWinRight, 0, 0},
	defaultOptions.ExclusiveWinRight = 0
	//
	//{0, REG_DWORD,  NKeyInterface, L"ShowTimeoutDelFiles",&Opt.ShowTimeoutDelFiles, 50, 0},
	defaultOptions.ShowTimeoutDelFiles = 50
	//{0, REG_DWORD,  NKeyInterface, L"ShowTimeoutDACLFiles",&Opt.ShowTimeoutDACLFiles, 50, 0},
	defaultOptions.ShowTimeoutDACLFiles = 50
	//{0, REG_DWORD,  NKeyInterface, L"FormatNumberSeparators",&Opt.FormatNumberSeparators, 0, 0},
	defaultOptions.FormatNumberSeparators = 0

	//{1, REG_DWORD,  NKeyInterface,L"DelShowTotal",&Opt.DelOpt.DelShowTotal,0, 0},
	defaultOptions.DelOpt = DeleteOptions{0}
	//{1, REG_SZ,     NKeyInterface,L"TitleAddons",&Opt.strTitleAddons, 0, L"%Ver %Build %Platform %Admin"},
	defaultOptions.strTitleAddons = "%Ver %Build %Platform %Admin"
	//{1, REG_SZ,     NKeyInterfaceCompletion,L"Exceptions",&Opt.AutoComplete.Exceptions, 0, L"git*reset*--hard; ftp://*:*@*"},
	//{1, REG_DWORD,  NKeyInterfaceCompletion,L"ShowList",&Opt.AutoComplete.ShowList, 1, 0},
	//{1, REG_DWORD,  NKeyInterfaceCompletion,L"ModalList",&Opt.AutoComplete.ModalList, 0, 0},
	//{1, REG_DWORD,  NKeyInterfaceCompletion,L"Append",&Opt.AutoComplete.AppendCompletion, 0, 0},
	defaultOptions.AutoComplete = AutoCompleteOptions{
		Exceptions:       "git*reset*--hard; ftp://*:*@*",
		ShowList:         1,
		ModalList:        0,
		AppendCompletion: 0,
	}
	//
	//{1, REG_SZ,     NKeyViewer,L"ExternalViewerName",&Opt.strExternalViewer, 0, L""},
	defaultOptions.strExternalViewer = ""
	//{1, REG_DWORD,  NKeyViewer,L"UseExternalViewer",&Opt.ViOpt.UseExternalViewer,0, 0},
	//{1, REG_DWORD,  NKeyViewer,L"SaveViewerPos",&Opt.ViOpt.SavePos,1, 0},
	//{1, REG_DWORD,  NKeyViewer,L"SaveViewerShortPos",&Opt.ViOpt.SaveShortPos,1, 0},
	//{1, REG_DWORD,  NKeyViewer,L"AutoDetectCodePage",&Opt.ViOpt.AutoDetectCodePage,0, 0},
	//{1, REG_DWORD,  NKeyViewer,L"SearchRegexp",&Opt.ViOpt.SearchRegexp,0, 0},
	//
	//{1, REG_DWORD,  NKeyViewer,L"TabSize",&Opt.ViOpt.TabSize,8, 0},
	//{1, REG_DWORD,  NKeyViewer,L"ShowKeyBar",&Opt.ViOpt.ShowKeyBar,1, 0},
	//{0, REG_DWORD,  NKeyViewer,L"ShowTitleBar",&Opt.ViOpt.ShowTitleBar,1, 0},
	//{1, REG_DWORD,  NKeyViewer,L"ShowArrows",&Opt.ViOpt.ShowArrows,1, 0},
	//{1, REG_DWORD,  NKeyViewer,L"ShowScrollbar",&Opt.ViOpt.ShowScrollbar,0, 0},
	//{1, REG_DWORD,  NKeyViewer,L"IsWrap",&Opt.ViOpt.ViewerIsWrap,1, 0},
	//{1, REG_DWORD,  NKeyViewer,L"Wrap",&Opt.ViOpt.ViewerWrap,0, 0},
	//{1, REG_DWORD,  NKeyViewer,L"PersistentBlocks",&Opt.ViOpt.PersistentBlocks,0, 0},
	//{1, REG_DWORD,  NKeyViewer,L"UTF8CodePageAsDefault",&Opt.ViOpt.UTF8CodePageAsDefault,1, 0},
	defaultOptions.ViOpt = ViewerOptions{
		UseExternalViewer:     0,
		SavePos:               1,
		SaveShortPos:          1,
		AutoDetectCodePage:    0,
		SearchRegexp:          0,
		TabSize:               8,
		ShowKeyBar:            1,
		ShowArrows:            1,
		ShowScrollbar:         0,
		ViewerIsWrap:          1,
		ViewerWrap:            0,
		PersistentBlocks:      0,
		UTF8CodePageAsDefault: 1,

	}
	//
	//{1, REG_DWORD,  NKeyDialog, L"EditHistory",&Opt.Dialogs.EditHistory,1, 0},
	//{1, REG_DWORD,  NKeyDialog, L"EditBlock",&Opt.Dialogs.EditBlock,0, 0},
	//{1, REG_DWORD,  NKeyDialog, L"AutoComplete",&Opt.Dialogs.AutoComplete,1, 0},
	//{1, REG_DWORD,  NKeyDialog,L"EULBsClear",&Opt.Dialogs.EULBsClear,0, 0},
	//{0, REG_DWORD,  NKeyDialog,L"SelectFromHistory",&Opt.Dialogs.SelectFromHistory,0, 0},
	//{0, REG_DWORD,  NKeyDialog,L"EditLine",&Opt.Dialogs.EditLine,0, 0},
	//{1, REG_DWORD,  NKeyDialog,L"MouseButton",&Opt.Dialogs.MouseButton,0xFFFF, 0},
	//{1, REG_DWORD,  NKeyDialog,L"DelRemovesBlocks",&Opt.Dialogs.DelRemovesBlocks,1, 0},
	//{0, REG_DWORD,  NKeyDialog,L"CBoxMaxHeight",&Opt.Dialogs.CBoxMaxHeight,24, 0},
	defaultOptions.Dialogs = DialogsOptions{
		EditHistory:       1,
		EditBlock:         0,
		AutoComplete:      1,
		EULBsClear:        0,
		SelectFromHistory: 0,
		EditLine:          0,
		MouseButton:       0xFFFF,
		DelRemovesBlocks:  1,
		CBoxMaxHeight:     24,
	}
	//
	//{1, REG_SZ,     NKeyEditor,L"ExternalEditorName",&Opt.strExternalEditor, 0, L""},
	defaultOptions.strExternalViewer = ""
	//{1, REG_DWORD,  NKeyEditor,L"UseExternalEditor",&Opt.EdOpt.UseExternalEditor,0, 0},
	//{1, REG_DWORD,  NKeyEditor,L"ExpandTabs",&Opt.EdOpt.ExpandTabs,0, 0},
	//{1, REG_DWORD,  NKeyEditor,L"TabSize",&Opt.EdOpt.TabSize,8, 0},
	//{1, REG_DWORD,  NKeyEditor,L"PersistentBlocks",&Opt.EdOpt.PersistentBlocks,0, 0},
	//{1, REG_DWORD,  NKeyEditor,L"DelRemovesBlocks",&Opt.EdOpt.DelRemovesBlocks,1, 0},
	//{1, REG_DWORD,  NKeyEditor,L"AutoIndent",&Opt.EdOpt.AutoIndent,0, 0},
	//{1, REG_DWORD,  NKeyEditor,L"SaveEditorPos",&Opt.EdOpt.SavePos,1, 0},
	//{1, REG_DWORD,  NKeyEditor,L"SaveEditorShortPos",&Opt.EdOpt.SaveShortPos,1, 0},
	//{1, REG_DWORD,  NKeyEditor,L"AutoDetectCodePage",&Opt.EdOpt.AutoDetectCodePage,0, 0},
	//{1, REG_DWORD,  NKeyEditor,L"EditorCursorBeyondEOL",&Opt.EdOpt.CursorBeyondEOL,1, 0},
	//{1, REG_DWORD,  NKeyEditor,L"ReadOnlyLock",&Opt.EdOpt.ReadOnlyLock,0, 0}, // Вернём назад дефолт 1.65 - не предупреждать и не блокировать
	//{0, REG_DWORD,  NKeyEditor,L"EditorUndoSize",&Opt.EdOpt.UndoSize,0, 0}, // $ 03.12.2001 IS размер буфера undo в редакторе
	//{0, REG_DWORD,  NKeyEditor,L"BSLikeDel",&Opt.EdOpt.BSLikeDel,1, 0},
	//{0, REG_DWORD,  NKeyEditor,L"EditorF7Rules",&Opt.EdOpt.F7Rules,1, 0},
	//{0, REG_DWORD,  NKeyEditor,L"FileSizeLimit",&Opt.EdOpt.FileSizeLimitLo,(DWORD)0, 0},
	//{0, REG_DWORD,  NKeyEditor,L"FileSizeLimitHi",&Opt.EdOpt.FileSizeLimitHi,(DWORD)0, 0},
	//{0, REG_DWORD,  NKeyEditor,L"CharCodeBase",&Opt.EdOpt.CharCodeBase,1, 0},
	//{0, REG_DWORD,  NKeyEditor,L"AllowEmptySpaceAfterEof", &Opt.EdOpt.AllowEmptySpaceAfterEof,0,0},//skv
	//{1, REG_DWORD,  NKeyEditor,L"Utf8CodePageForNewFile",&Opt.EdOpt.UTF8CodePageForNewFile,1, 0},
	//{1, REG_DWORD,  NKeyEditor,L"UTF8CodePageAsDefault",&Opt.EdOpt.UTF8CodePageAsDefault,1, 0},
	//{1, REG_DWORD,  NKeyEditor,L"ShowKeyBar",&Opt.EdOpt.ShowKeyBar,1, 0},
	//{0, REG_DWORD,  NKeyEditor,L"ShowTitleBar",&Opt.EdOpt.ShowTitleBar,1, 0},
	//{1, REG_DWORD,  NKeyEditor,L"ShowScrollBar",&Opt.EdOpt.ShowScrollBar,0, 0},
	//{1, REG_DWORD,  NKeyEditor,L"EditOpenedForWrite",&Opt.EdOpt.EditOpenedForWrite,1, 0},
	//{1, REG_DWORD,  NKeyEditor,L"SearchSelFound",&Opt.EdOpt.SearchSelFound,0, 0},
	//{1, REG_DWORD,  NKeyEditor,L"SearchRegexp",&Opt.EdOpt.SearchRegexp,0, 0},
	//{1, REG_DWORD,  NKeyEditor,L"SearchPickUpWord",&Opt.EdOpt.SearchPickUpWord,0, 0},
	//{1, REG_DWORD,  NKeyEditor,L"ShowWhiteSpace",&Opt.EdOpt.ShowWhiteSpace,0, 0},
	defaultOptions.EdOpt = EditorOptions{
		UseExternalEditor:       0,
		ExpandTabs:              0,
		TabSize:                 8,
		PersistentBlocks:        0,
		DelRemovesBlocks:        1,
		AutoIndent:              0,
		SavePos:                 1,
		SaveShortPos:            1,
		AutoDetectCodePage:      0,
		CursorBeyondEOL:         1,
		ReadOnlyLock:            0,
		UndoSize:                0,
		BSLikeDel:               1,
		F7Rules:                 1,
		FileSizeLimitLo:         0,
		FileSizeLimitHi:         0,
		CharCodeBase:            1,
		AllowEmptySpaceAfterEof: 0,
		UTF8CodePageForNewFile:  1,
		UTF8CodePageAsDefault:   1,
		ShowKeyBar:              1,
		ShowTitleBar:            1,
		ShowScrollBar:           0,
		EditOpenedForWrite:      1,
		SearchSelFound:          0,
		SearchRegexp:            0,
		SearchPickUpWord:        0,
		ShowWhiteSpace:          0,

	}

	//{0, REG_SZ,     NKeyEditor,L"WordDiv",&Opt.strWordDiv, 0, WordDiv0},
	defaultOptions.strWordDiv = WordDiv0

	//
	//{0, REG_DWORD,  NKeyXLat,L"Flags",&Opt.XLat.Flags,(DWORD)XLAT_SWITCHKEYBLAYOUT|XLAT_CONVERTALLCMDLINE, 0},
	//{0, REG_SZ,     NKeyXLat,L"Table1",&Opt.XLat.Table[0],0,L""},
	//{0, REG_SZ,     NKeyXLat,L"Table2",&Opt.XLat.Table[1],0,L""},
	//{0, REG_SZ,     NKeyXLat,L"Rules1",&Opt.XLat.Rules[0],0,L""},
	//{0, REG_SZ,     NKeyXLat,L"Rules2",&Opt.XLat.Rules[1],0,L""},
	//{0, REG_SZ,     NKeyXLat,L"Rules3",&Opt.XLat.Rules[2],0,L""},
	//{0, REG_SZ,     NKeyXLat,L"WordDivForXlat",&Opt.XLat.strWordDivForXlat, 0,WordDivForXlat0},
	defaultOptions.XLat = CodeXLAT{
		Flags:             XLAT_SWITCHKEYBLAYOUT | XLAT_CONVERTALLCMDLINE,
		Table:             []string{"", ""},
		Rules:             []string{"", "", ""},
		strWordDivForXlat: WordDivForXlat0,
	}

	//
	//{0, REG_DWORD,  NKeySavedHistory, NParamHistoryCount,&Opt.HistoryCount,512, 0},
	defaultOptions.HistoryCount = 512
	//{0, REG_DWORD,  NKeySavedFolderHistory, NParamHistoryCount,&Opt.FoldersHistoryCount,512, 0},
	defaultOptions.FoldersHistoryCount = 512
	//{0, REG_DWORD,  NKeySavedViewHistory, NParamHistoryCount,&Opt.ViewHistoryCount,512, 0},
	defaultOptions.ViewHistoryCount = 512
	//{0, REG_DWORD,  NKeySavedDialogHistory, NParamHistoryCount,&Opt.DialogsHistoryCount,512, 0},
	defaultOptions.DialogsHistoryCount = 512
	//
	//{1, REG_DWORD,  NKeySystem,L"SaveHistory",&Opt.SaveHistory,1, 0},
	defaultOptions.SaveHistory = 1
	//{1, REG_DWORD,  NKeySystem,L"SaveFoldersHistory",&Opt.SaveFoldersHistory,1, 0},
	defaultOptions.SaveFoldersHistory = 1
	//{0, REG_DWORD,  NKeySystem,L"SavePluginFoldersHistory",&Opt.SavePluginFoldersHistory,0, 0},
	defaultOptions.SavePluginFoldersHistory = 0
	//{1, REG_DWORD,  NKeySystem,L"SaveViewHistory",&Opt.SaveViewHistory,1, 0},
	defaultOptions.SaveViewHistory = 1
	//{1, REG_DWORD,  NKeySystem,L"UseRegisteredTypes",&Opt.UseRegisteredTypes,1, 0},
	defaultOptions.UseRegisteredTypes = 1
	//{1, REG_DWORD,  NKeySystem,L"AutoSaveSetup",&Opt.AutoSaveSetup,0, 0},
	defaultOptions.AutoSaveSetup = 0
	//{1, REG_DWORD,  NKeySystem,L"DeleteToRecycleBin",&Opt.DeleteToRecycleBin,0, 0},
	defaultOptions.DeleteToRecycleBin = 0
	//{1, REG_DWORD,  NKeySystem,L"DeleteToRecycleBinKillLink",&Opt.DeleteToRecycleBinKillLink,1, 0},
	defaultOptions.DeleteToRecycleBinKillLink = 1
	//{0, REG_DWORD,  NKeySystem,L"WipeSymbol",&Opt.WipeSymbol,0, 0},
	defaultOptions.WipeSymbol = 0
	//{1, REG_DWORD,  NKeySystem,L"SudoEnabled",&Opt.SudoEnabled,1, 0},
	defaultOptions.SudoEnabled = 1
	//{1, REG_DWORD,  NKeySystem,L"SudoConfirmModify",&Opt.SudoConfirmModify,1, 0},
	defaultOptions.SudoConfirmModify = 1
	//{1, REG_DWORD,  NKeySystem,L"SudoPasswordExpiration",&Opt.SudoPasswordExpiration,15*60, 0},
	defaultOptions.SudoPasswordExpiration = 15 * 60
	//
	//{1, REG_DWORD,  NKeyInterface, L"CopyShowTotal",&Opt.CMOpt.CopyShowTotal,1, 0},
	//{1, REG_DWORD,  NKeySystem,L"WriteThrough",&Opt.CMOpt.WriteThrough, 0, 0},
	//{1, REG_DWORD,  NKeySystem,L"CopyXAttr",&Opt.CMOpt.CopyXAttr, 0, 0},
	//{0, REG_DWORD,  NKeySystem,L"CopyAccessMode",&Opt.CMOpt.CopyAccessMode,1, 0},
	//{1, REG_DWORD,  NKeySystem, L"MultiCopy",&Opt.CMOpt.MultiCopy,0, 0},
	//{1, REG_DWORD,  NKeySystem,L"CopyTimeRule",  &Opt.CMOpt.CopyTimeRule, 3, 0},
	defaultOptions.CMOpt = CopyMoveOptions{
		CopyShowTotal:  1,
		WriteThrough:   0,
		CopyXAttr:      0,
		CopyAccessMode: 1,
		MultiCopy:      0,
		CopyTimeRule:   3,
	}
	//
	//{1, REG_DWORD,  NKeySystem,L"InactivityExit",&Opt.InactivityExit,0, 0},
	defaultOptions.InactivityExit = 0
	//{1, REG_DWORD,  NKeySystem,L"InactivityExitTime",&Opt.InactivityExitTime,15, 0},
	defaultOptions.InactivityExitTime = 15
	//{1, REG_DWORD,  NKeySystem,L"DriveMenuMode",&Opt.ChangeDriveMode,DRIVE_SHOW_TYPE|DRIVE_SHOW_PLUGINS|DRIVE_SHOW_SIZE_FLOAT|DRIVE_SHOW_CDROM, 0},
	defaultOptions.ChangeDriveMode = DRIVE_SHOW_TYPE | DRIVE_SHOW_PLUGINS | DRIVE_SHOW_SIZE_FLOAT | DRIVE_SHOW_CDROM
	//{1, REG_DWORD,  NKeySystem,L"DriveDisconnetMode",&Opt.ChangeDriveDisconnetMode,1, 0},
	defaultOptions.ChangeDriveDisconnetMode = 1
	//{1, REG_DWORD,  NKeySystem,L"AutoUpdateRemoteDrive",&Opt.AutoUpdateRemoteDrive,1, 0},
	defaultOptions.AutoUpdateRemoteDrive = 1
	//{1, REG_DWORD,  NKeySystem,L"FileSearchMode",&Opt.FindOpt.FileSearchMode,FINDAREA_FROM_CURRENT, 0},
	//{0, REG_DWORD,  NKeySystem,L"CollectFiles",&Opt.FindOpt.CollectFiles, 1, 0},
	//{1, REG_SZ,     NKeySystem,L"SearchInFirstSize",&Opt.FindOpt.strSearchInFirstSize, 0, L""},
	//{1, REG_DWORD,  NKeySystem,L"FindAlternateStreams",&Opt.FindOpt.FindAlternateStreams,0,0},
	//{1, REG_SZ,     NKeySystem,L"SearchOutFormat",&Opt.FindOpt.strSearchOutFormat, 0, L"D,S,A"},
	//{1, REG_SZ,     NKeySystem,L"SearchOutFormatWidth",&Opt.FindOpt.strSearchOutFormatWidth, 0, L"14,13,0"},
	//{1, REG_DWORD,  NKeySystem,L"FindFolders",&Opt.FindOpt.FindFolders, 1, 0},
	//{1, REG_DWORD,  NKeySystem,L"FindSymLinks",&Opt.FindOpt.FindSymLinks, 1, 0},
	//{1, REG_DWORD,  NKeySystem,L"UseFilterInSearch",&Opt.FindOpt.UseFilter,0,0},
	defaultOptions.FindOpt = FindFileOptions{
		FileSearchMode:          FINDAREA_FROM_CURRENT,
		CollectFiles:            true,
		strSearchInFirstSize:    "",
		FindAlternateStreams:    false,
		strSearchOutFormat:      "D,S,A",
		strSearchOutFormatWidth: "14,13,0",
		FindFolders:             true,
		FindSymLinks:            true,
		UseFilter:               false,


	}

	//{1, REG_DWORD,  NKeySystem,L"FindCodePage",&Opt.FindCodePage, CP_AUTODETECT, 0},
	defaultOptions.FindCodePage = CP_AUTODETECT
	//{0, REG_DWORD,  NKeySystem,L"SubstPluginPrefix",&Opt.SubstPluginPrefix, 0, 0},
	defaultOptions.SubstPluginPrefix = 0
	//{0, REG_DWORD,  NKeySystem,L"CmdHistoryRule",&Opt.CmdHistoryRule,0, 0},
	defaultOptions.CmdHistoryRule = 0
	//{0, REG_DWORD,  NKeySystem,L"SetAttrFolderRules",&Opt.SetAttrFolderRules,1, 0},
	defaultOptions.SetAttrFolderRules = 1
	//{0, REG_DWORD,  NKeySystem,L"MaxPositionCache",&Opt.MaxPositionCache,MAX_POSITIONS, 0},
	defaultOptions.MaxPositionCache = MAX_POSITIONS
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//{0, REG_SZ,     NKeySystem,L"ConsoleDetachKey", &strKeyNameConsoleDetachKey, 0, L"CtrlAltTab"},
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	//{0, REG_DWORD,  NKeySystem,L"SilentLoadPlugin",  &Opt.LoadPlug.SilentLoadPlugin, 0, 0},
	//{1, REG_DWORD,  NKeySystem,L"OEMPluginsSupport",  &Opt.LoadPlug.OEMPluginsSupport, 1, 0},
	//{1, REG_DWORD,  NKeySystem,L"ScanSymlinks",  &Opt.LoadPlug.ScanSymlinks, 1, 0},
	defaultOptions.LoadPlug = LoadPluginsOptions{
		SilentLoadPlugin:  0,
		OEMPluginsSupport: 1,
		ScanSymlinks:      1,
	}

	//{1, REG_DWORD,  NKeySystem,L"MultiMakeDir",&Opt.MultiMakeDir,0, 0},
	defaultOptions.MultiMakeDir = 0
	//{0, REG_DWORD,  NKeySystem,L"MsWheelDelta", &Opt.MsWheelDelta, 1, 0},
	defaultOptions.MsWheelDelta = 1
	//{0, REG_DWORD,  NKeySystem,L"MsWheelDeltaView", &Opt.MsWheelDeltaView, 1, 0},
	defaultOptions.MsWheelDeltaView = 1
	//{0, REG_DWORD,  NKeySystem,L"MsWheelDeltaEdit", &Opt.MsWheelDeltaEdit, 1, 0},
	defaultOptions.MsWheelDeltaEdit = 1
	//{0, REG_DWORD,  NKeySystem,L"MsWheelDeltaHelp", &Opt.MsWheelDeltaHelp, 1, 0},
	defaultOptions.MsWheelDeltaHelp = 1
	//{0, REG_DWORD,  NKeySystem,L"MsHWheelDelta", &Opt.MsHWheelDelta, 1, 0},
	defaultOptions.MsHWheelDelta = 1
	//{0, REG_DWORD,  NKeySystem,L"MsHWheelDeltaView", &Opt.MsHWheelDeltaView, 1, 0},
	defaultOptions.MsHWheelDeltaView = 1
	//{0, REG_DWORD,  NKeySystem,L"MsHWheelDeltaEdit", &Opt.MsHWheelDeltaEdit, 1, 0},
	defaultOptions.MsHWheelDeltaEdit = 1
	//{0, REG_DWORD,  NKeySystem,L"SubstNameRule", &Opt.SubstNameRule, 2, 0},
	defaultOptions.SubstNameRule = 2
	//{0, REG_DWORD,  NKeySystem,L"ShowCheckingFile", &Opt.ShowCheckingFile, 0, 0},
	defaultOptions.ShowCheckingFile = 0
	//{0, REG_DWORD,  NKeySystem,L"DelThreadPriority", &Opt.DelThreadPriority, 0, 0},
	defaultOptions.DelThreadPriority = 0
	//{0, REG_SZ,     NKeySystem,L"QuotedSymbols",&Opt.strQuotedSymbols, 0, L" $&()[]{};|*?!'`\"\\\xA0"}, //xA0 => 160 =>oem(0xFF)
	defaultOptions.strQuotedSymbols = " $&()[]{};|*?!'`\"\\\xA0"
	//{0, REG_DWORD,  NKeySystem,L"QuotedName",&Opt.QuotedName,0xFFFFFFFFU, 0},
	defaultOptions.QuotedName = 0xFFFFFFFF
	////{0, REG_DWORD,  NKeySystem,L"CPAJHefuayor",&Opt.strCPAJHefuayor,0, 0},
	//{0, REG_DWORD,  NKeySystem,L"CloseConsoleRule",&Opt.CloseConsoleRule,1, 0},
	defaultOptions.CloseConsoleRule = 1
	//{0, REG_DWORD,  NKeySystem,L"PluginMaxReadData",&Opt.PluginMaxReadData,0x20000, 0},
	defaultOptions.PluginMaxReadData = 0x20000
	//{1, REG_DWORD,  NKeySystem,L"CloseCDGate",&Opt.CloseCDGate,1, 0},
	defaultOptions.CloseCDGate = 1
	//{1, REG_DWORD,  NKeySystem,L"UpdateEnvironment",&Opt.UpdateEnvironment,0,0},
	defaultOptions.UpdateEnvironment = 0
	//{0, REG_DWORD,  NKeySystem,L"UseNumPad",&Opt.UseNumPad,1, 0},
	defaultOptions.UseNumPad = 1
	//{0, REG_DWORD,  NKeySystem,L"CASRule",&Opt.CASRule,0xFFFFFFFFU, 0},
	defaultOptions.CASRule = 0xFFFFFFFF
	//{0, REG_DWORD,  NKeySystem,L"AllCtrlAltShiftRule",&Opt.AllCtrlAltShiftRule,0x0000FFFF, 0},
	defaultOptions.AllCtrlAltShiftRulev = 0x0000FFFF
	//{1, REG_DWORD,  NKeySystem,L"ScanJunction",&Opt.ScanJunction,1, 0},
	defaultOptions.ScanJunction = 1
	//{1, REG_DWORD,  NKeySystem,L"OnlyFilesSize",&Opt.OnlyFilesSize, 0, 0},
	defaultOptions.OnlyFilesSize = 0
	//{0, REG_DWORD,  NKeySystem,L"UsePrintManager",&Opt.UsePrintManager,1, 0},
	defaultOptions.UsePrintManager = 1
	//{0, REG_DWORD,  NKeySystem,L"WindowMode",&Opt.WindowMode, 0, 0},
	defaultOptions.WindowMode = false
	//
	//{0, REG_DWORD,  NKeySystemNowell,L"MoveRO",&Opt.Nowell.MoveRO,1, 0},
	defaultOptions.Nowell = NowellOptions{MoveRO: 1}
	//
	//{0, REG_DWORD,  NKeySystemExecutor,L"RestoreCP",&Opt.RestoreCPAfterExecute,1, 0},
	defaultOptions.RestoreCPAfterExecute = 1

	//{0, REG_DWORD,  NKeySystemExecutor,L"UseAppPath",&Opt.ExecuteUseAppPath,1, 0},
	defaultOptions.ExecuteUseAppPath = 1
	//{0, REG_DWORD,  NKeySystemExecutor,L"ShowErrorMessage",&Opt.ExecuteShowErrorMessage,1, 0},
	defaultOptions.ExecuteShowErrorMessage = 1
	//{0, REG_SZ,     NKeySystemExecutor,L"BatchType",&Opt.strExecuteBatchType,0,constBatchExt},
	defaultOptions.strExternalViewer = constBatchExt
	//{0, REG_DWORD,  NKeySystemExecutor,L"FullTitle",&Opt.ExecuteFullTitle,0, 0},
	defaultOptions.ExecuteFullTitle = 0
	//{0, REG_DWORD,  NKeySystemExecutor,L"SilentExternal",&Opt.ExecuteSilentExternal,0, 0},
	defaultOptions.ExecuteSilentExternal = 0
	//
	//{0, REG_DWORD,  NKeyPanelTree,L"MinTreeCount",&Opt.Tree.MinTreeCount, 4, 0},
	//{0, REG_DWORD,  NKeyPanelTree,L"TreeFileAttr",&Opt.Tree.TreeFileAttr, FILE_ATTRIBUTE_HIDDEN, 0},
	//{0, REG_DWORD,  NKeyPanelTree,L"LocalDisk",&Opt.Tree.LocalDisk, 2, 0},
	//{0, REG_DWORD,  NKeyPanelTree,L"NetDisk",&Opt.Tree.NetDisk, 2, 0},
	//{0, REG_DWORD,  NKeyPanelTree,L"RemovableDisk",&Opt.Tree.RemovableDisk, 2, 0},
	//{0, REG_DWORD,  NKeyPanelTree,L"NetPath",&Opt.Tree.NetPath, 2, 0},
	//{1, REG_DWORD,  NKeyPanelTree,L"AutoChangeFolder",&Opt.Tree.AutoChangeFolder,0, 0}, // ???
	defaultOptions.Tree = TreeOptions{
		MinTreeCount:     4,
		TreeFileAttr:     FILE_ATTRIBUTE_HIDDEN,
		LocalDisk:        2,
		NetDisk:          2,
		RemovableDisk:    2,
		NetPath:          2,
		AutoChangeFolder: 0,
	}

	//
	//{0, REG_DWORD,  NKeyHelp,L"ActivateURL",&Opt.HelpURLRules,1, 0},
	defaultOptions.HelpURLRules = 1
	//
	//{1, REG_SZ,     NKeyLanguage,L"Help",&Opt.strHelpLanguage, 0, L"English"},
	defaultOptions.strHelpLanguage = "English"
	//
	//{1, REG_DWORD,  NKeyConfirmations,L"Copy",&Opt.Confirm.Copy,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"Move",&Opt.Confirm.Move,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"RO",&Opt.Confirm.RO,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"Drag",&Opt.Confirm.Drag,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"Delete",&Opt.Confirm.Delete,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"DeleteFolder",&Opt.Confirm.DeleteFolder,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"Esc",&Opt.Confirm.Esc,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"RemoveConnection",&Opt.Confirm.RemoveConnection,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"RemoveSUBST",&Opt.Confirm.RemoveSUBST,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"DetachVHD",&Opt.Confirm.DetachVHD,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"RemoveHotPlug",&Opt.Confirm.RemoveHotPlug,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"AllowReedit",&Opt.Confirm.AllowReedit,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"HistoryClear",&Opt.Confirm.HistoryClear,1, 0},
	//{1, REG_DWORD,  NKeyConfirmations,L"Exit",&Opt.Confirm.Exit,1, 0},
	//{0, REG_DWORD,  NKeyConfirmations,L"EscTwiceToInterrupt",&Opt.Confirm.EscTwiceToInterrupt,0, 0},
	defaultOptions.Confirm = Confirmation{
		Copy:                1,
		Move:                1,
		RO:                  1,
		Drag:                1,
		Delete:              1,
		DeleteFolder:        1,
		Esc:                 1,
		RemoveConnection:    1,
		RemoveSUBST:         1,
		DetachVHD:           1,
		RemoveHotPlug:       1,
		AllowReedit:         1,
		HistoryClear:        1,
		Exit:                1,
		EscTwiceToInterrupt: 0,
	}

	//
	//{1, REG_DWORD,  NKeyPluginConfirmations, L"OpenFilePlugin", &Opt.PluginConfirm.OpenFilePlugin, 0, 0},
	//{1, REG_DWORD,  NKeyPluginConfirmations, L"StandardAssociation", &Opt.PluginConfirm.StandardAssociation, 0, 0},
	//{1, REG_DWORD,  NKeyPluginConfirmations, L"EvenIfOnlyOnePlugin", &Opt.PluginConfirm.EvenIfOnlyOnePlugin, 0, 0},
	//{1, REG_DWORD,  NKeyPluginConfirmations, L"SetFindList", &Opt.PluginConfirm.SetFindList, 0, 0},
	//{1, REG_DWORD,  NKeyPluginConfirmations, L"Prefix", &Opt.PluginConfirm.Prefix, 0, 0},
	defaultOptions.PluginConfirm = PluginConfirmation{
		OpenFilePlugin:      0,
		StandardAssociation: 0,
		EvenIfOnlyOnePlugin: 0,
		SetFindList:         0,
		Prefix:              0,
	}
	//
	//{0, REG_DWORD,  NKeyPanel,L"ShellRightLeftArrowsRule",&Opt.ShellRightLeftArrowsRule,0, 0},
	defaultOptions.ShellRightLeftArrowsRule = 0
	//{1, REG_DWORD,  NKeyPanel,L"ShowHidden",&Opt.ShowHidden,1, 0},
	defaultOptions.ShowHidden = 1
	//{1, REG_DWORD,  NKeyPanel,L"Highlight",&Opt.Highlight,1, 0},
	defaultOptions.Highlight = 1
	//{1, REG_DWORD,  NKeyPanel,L"SortFolderExt",&Opt.SortFolderExt,0, 0},
	defaultOptions.SortFolderExt = 0
	//{1, REG_DWORD,  NKeyPanel,L"SelectFolders",&Opt.SelectFolders,0, 0},
	defaultOptions.SelectFolders = 0
	//{1, REG_DWORD,  NKeyPanel,L"ReverseSort",&Opt.ReverseSort,1, 0},
	defaultOptions.ReverseSort = 1
	//{0, REG_DWORD,  NKeyPanel,L"RightClickRule",&Opt.PanelRightClickRule,2, 0},
	defaultOptions.PanelRightClickRule = 2
	//{0, REG_DWORD,  NKeyPanel,L"CtrlFRule",&Opt.PanelCtrlFRule,1, 0},
	defaultOptions.PanelCtrlFRule = 1
	//{0, REG_DWORD,  NKeyPanel,L"CtrlAltShiftRule",&Opt.PanelCtrlAltShiftRule,0, 0},
	defaultOptions.PanelCtrlAltShiftRule = 0
	//{0, REG_DWORD,  NKeyPanel,L"RememberLogicalDrives",&Opt.RememberLogicalDrives, 0, 0},
	defaultOptions.RememberLogicalDrives = 0
	//{1, REG_DWORD,  NKeyPanel,L"AutoUpdateLimit",&Opt.AutoUpdateLimit, 0, 0},
	defaultOptions.AutoUpdateLimit = 0
	//
	//{1, REG_DWORD,  NKeyPanelLeft,L"Type",&Opt.LeftPanel.Type,0, 0},
	//{1, REG_DWORD,  NKeyPanelLeft,L"Visible",&Opt.LeftPanel.Visible,1, 0},
	//{1, REG_DWORD,  NKeyPanelLeft,L"Focus",&Opt.LeftPanel.Focus,1, 0},
	//{1, REG_DWORD,  NKeyPanelLeft,L"ViewMode",&Opt.LeftPanel.ViewMode,2, 0},
	//{1, REG_DWORD,  NKeyPanelLeft,L"SortMode",&Opt.LeftPanel.SortMode,1, 0},
	//{1, REG_DWORD,  NKeyPanelLeft,L"SortOrder",&Opt.LeftPanel.SortOrder,1, 0},
	//{1, REG_DWORD,  NKeyPanelLeft,L"SortGroups",&Opt.LeftPanel.SortGroups,0, 0},
	//{1, REG_DWORD,  NKeyPanelLeft,L"NumericSort",&Opt.LeftPanel.NumericSort,0, 0},
	//{1, REG_DWORD,  NKeyPanelLeft,L"CaseSensitiveSort",&Opt.LeftPanel.CaseSensitiveSort,0, 0},
	//{1, REG_DWORD,  NKeyPanelLeft,L"DirectoriesFirst",&Opt.LeftPanel.DirectoriesFirst,1,0},
	defaultOptions.LeftPanel = PanelOptions{
		Type:              0,
		Visible:           1,
		Focus:             1,
		ViewMode:          2,
		SortMode:          1,
		SortOrder:         1,
		SortGroups:        0,
		NumericSort:       0,
		CaseSensitiveSort: 0,
		DirectoriesFirst:  1,

	}

	//{1, REG_SZ,     NKeyPanelLeft,L"Folder",&Opt.strLeftFolder, 0, L""},
	defaultOptions.strLeftFolder = ""
	//{1, REG_SZ,     NKeyPanelLeft,L"CurFile",&Opt.strLeftCurFile, 0, L""},
	defaultOptions.strLeftCurFile = ""
	//{1, REG_DWORD,  NKeyPanelLeft,L"SelectedFirst",&Opt.LeftSelectedFirst,0,0},
	defaultOptions.LeftSelectedFirst = 0

	//
	//{1, REG_DWORD,  NKeyPanelRight,L"Type",&Opt.RightPanel.Type,0, 0},
	//{1, REG_DWORD,  NKeyPanelRight,L"Visible",&Opt.RightPanel.Visible,1, 0},
	//{1, REG_DWORD,  NKeyPanelRight,L"Focus",&Opt.RightPanel.Focus,0, 0},
	//{1, REG_DWORD,  NKeyPanelRight,L"ViewMode",&Opt.RightPanel.ViewMode,2, 0},
	//{1, REG_DWORD,  NKeyPanelRight,L"SortMode",&Opt.RightPanel.SortMode,1, 0},
	//{1, REG_DWORD,  NKeyPanelRight,L"SortOrder",&Opt.RightPanel.SortOrder,1, 0},
	//{1, REG_DWORD,  NKeyPanelRight,L"SortGroups",&Opt.RightPanel.SortGroups,0, 0},
	//{1, REG_DWORD,  NKeyPanelRight,L"NumericSort",&Opt.RightPanel.NumericSort,0, 0},
	//{1, REG_DWORD,  NKeyPanelRight,L"CaseSensitiveSort",&Opt.RightPanel.CaseSensitiveSort,0, 0},
	//{1, REG_DWORD,  NKeyPanelRight,L"DirectoriesFirst",&Opt.RightPanel.DirectoriesFirst,1,0},

	defaultOptions.RightPanel = PanelOptions{
		Type:              0,
		Visible:           1,
		Focus:             0,
		ViewMode:          2,
		SortMode:          1,
		SortOrder:         1,
		SortGroups:        0,
		NumericSort:       0,
		CaseSensitiveSort: 0,
		DirectoriesFirst:  1,

	}

	//{1, REG_SZ,     NKeyPanelRight,L"Folder",&Opt.strRightFolder, 0,L""},
	defaultOptions.strRightFolder = ""
	//{1, REG_SZ,     NKeyPanelRight,L"CurFile",&Opt.strRightCurFile, 0,L""},
	defaultOptions.strRightCurFile = ""
	//{1, REG_DWORD,  NKeyPanelRight,L"SelectedFirst",&Opt.RightSelectedFirst,0, 0},
	defaultOptions.RightSelectedFirst = 0

	//
	//{1, REG_DWORD,  NKeyPanelLayout,L"ColumnTitles",&Opt.ShowColumnTitles,1, 0},
	defaultOptions.ShowColumnTitles = 1
	//{1, REG_DWORD,  NKeyPanelLayout,L"StatusLine",&Opt.ShowPanelStatus,1, 0},
	defaultOptions.ShowPanelStatus = 1
	//{1, REG_DWORD,  NKeyPanelLayout,L"TotalInfo",&Opt.ShowPanelTotals,1, 0},
	defaultOptions.ShowPanelTotals = 1
	//{1, REG_DWORD,  NKeyPanelLayout,L"FreeInfo",&Opt.ShowPanelFree,0, 0},
	defaultOptions.ShowPanelFree = 0
	//{1, REG_DWORD,  NKeyPanelLayout,L"Scrollbar",&Opt.ShowPanelScrollbar,0, 0},
	defaultOptions.ShowPanelScrollbar = 0
	//{0, REG_DWORD,  NKeyPanelLayout,L"ScrollbarMenu",&Opt.ShowMenuScrollbar,1, 0},
	defaultOptions.ShowMenuScrollbar = 1
	//{1, REG_DWORD,  NKeyPanelLayout,L"ScreensNumber",&Opt.ShowScreensNumber,1, 0},
	defaultOptions.ShowScreensNumber = 1
	//{1, REG_DWORD,  NKeyPanelLayout,L"SortMode",&Opt.ShowSortMode,1, 0},
	defaultOptions.ShowSortMode = 1
	//
	//{1, REG_DWORD,  NKeyLayout,L"LeftHeightDecrement",&Opt.LeftHeightDecrement,0, 0},
	defaultOptions.LeftHeightDecrement = 0
	//{1, REG_DWORD,  NKeyLayout,L"RightHeightDecrement",&Opt.RightHeightDecrement,0, 0},
	defaultOptions.RightHeightDecrement = 0
	//{1, REG_DWORD,  NKeyLayout,L"WidthDecrement",&Opt.WidthDecrement,0, 0},
	defaultOptions.WidthDecrement = 0
	//{1, REG_DWORD,  NKeyLayout,L"FullscreenHelp",&Opt.FullScreenHelp,0, 0},
	defaultOptions.FullScreenHelp = 0
	//
	//{1, REG_SZ,     NKeyDescriptions,L"ListNames",&Opt.Diz.strListNames, 0, L"Descript.ion,Files.bbs"},
	//{1, REG_DWORD,  NKeyDescriptions,L"UpdateMode",&Opt.Diz.UpdateMode,DIZ_UPDATE_IF_DISPLAYED, 0},
	//{1, REG_DWORD,  NKeyDescriptions,L"ROUpdate",&Opt.Diz.ROUpdate,0, 0},
	//{1, REG_DWORD,  NKeyDescriptions,L"SetHidden",&Opt.Diz.SetHidden,1, 0},
	//{1, REG_DWORD,  NKeyDescriptions,L"StartPos",&Opt.Diz.StartPos,0, 0},
	//{1, REG_DWORD,  NKeyDescriptions,L"AnsiByDefault",&Opt.Diz.AnsiByDefault,0, 0},
	//{1, REG_DWORD,  NKeyDescriptions,L"SaveInUTF",&Opt.Diz.SaveInUTF,0, 0},
	defaultOptions.Diz = DizOptions{
		strListNames:  "Descript.ion,Files.bbs",
		UpdateMode:    DIZ_UPDATE_IF_DISPLAYED,
		ROUpdate:      0,
		SetHidden:     1,
		StartPos:      0,
		AnsiByDefault: 0,
		SaveInUTF:     0,
	}

	//
	//{0, REG_DWORD,  NKeyKeyMacros,L"MacroReuseRules",&Opt.Macro.MacroReuseRules,0, 0},
	//{0, REG_SZ,     NKeyKeyMacros,L"DateFormat",&Opt.Macro.strDateFormat, 0, L"%a %b %d %H:%M:%S %Z %Y"},
	//{0, REG_SZ,     NKeyKeyMacros,L"CONVFMT",&Opt.Macro.strMacroCONVFMT, 0, L"%.6g"},
	//{0, REG_DWORD,  NKeyKeyMacros,L"CallPluginRules",&Opt.Macro.CallPluginRules,0, 0},
	defaultOptions.Macro = MacroOptions{
		MacroReuseRules: 0,
		strDateFormat:   "%a %b %d %H:%M:%S %Z %Y",
		strMacroCONVFMT: "%.6g",
		CallPluginRules: 0,
	}
	//
	//{0, REG_DWORD,  NKeyPolicies,L"ShowHiddenDrives",&Opt.Policies.ShowHiddenDrives,1, 0},
	//{0, REG_DWORD,  NKeyPolicies,L"DisabledOptions",&Opt.Policies.DisabledOptions,0, 0},
	defaultOptions.Policies = PoliciesOptions{
		ShowHiddenDrives: 1,
		DisabledOptions:  0,
	}
	//
	//
	//{0, REG_DWORD,  NKeySystem,L"ExcludeCmdHistory",&Opt.ExcludeCmdHistory,0, 0}, //AN
	defaultOptions.ExcludeCmdHistory = 0
	//
	//{1, REG_DWORD,  NKeyCodePages,L"CPMenuMode",&Opt.CPMenuMode,0,0},
	defaultOptions.CPMenuMode = 0
	//
	//{1, REG_SZ,     NKeySystem,L"FolderInfo",&Opt.InfoPanel.strFolderInfoFiles, 0, L"DirInfo,File_Id.diz,Descript.ion,ReadMe.*,Read.Me"},
	defaultOptions.InfoPanel = InfoPanelOptions{
		strFolderInfoFiles: "DirInfo,File_Id.diz,Descript.ion,ReadMe.*,Read.Me",
	}
	//
	//{1, REG_DWORD,  NKeyVMenu,L"LBtnClick",&Opt.VMenu.LBtnClick, VMENUCLICK_CANCEL, 0},
	//{1, REG_DWORD,  NKeyVMenu,L"RBtnClick",&Opt.VMenu.RBtnClick, VMENUCLICK_CANCEL, 0},
	//{1, REG_DWORD,  NKeyVMenu,L"MBtnClick",&Opt.VMenu.MBtnClick, VMENUCLICK_APPLY, 0},
	defaultOptions.VMenu = VMenuOptions{
		LBtnClick: VMENUCLICK_CANCEL,
		RBtnClick: VMENUCLICK_CANCEL,
		MBtnClick: VMENUCLICK_APPLY,
	}

	return &defaultOptions
}
