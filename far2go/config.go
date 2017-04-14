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
	UpdateMode    int
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
	MreeCount        int
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
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Opt = Options{}
	// unmarshal config
	err = viper.Unmarshal(&Opt)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal to Unmarshal: %s \n", err))
	}

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
