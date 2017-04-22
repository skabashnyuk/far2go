package far2go

import (
	"github.com/sirupsen/logrus"
)

type ControlObject struct {
	fPanels       *FilePanels
	CmdLine       *CommandLine
	CmdHistory    *History
	ViewHistory   *History
	FolderHistory *History

	MainKeyBar     *KeyBar
	TopMenuBar     *MenuBar
	HiFiles        *HighlightFiles
	ViewerPosCache *FilePositionCache
	EditorPosCache *FilePositionCache
	Macro          *KeyMacro
	//Plugins        PluginManager
}

func NewControlObject() (*ControlObject) {
	ReadConfig()
	controlObject := &ControlObject{
		fPanels: new(FilePanels),
		CmdLine: new(CommandLine),
		CmdHistory: NewHistory(HISTORYTYPE_CMD,
			Opt.HistoryCount,
			"SavedHistory",
			Opt.SaveHistory,
			false),
		FolderHistory: NewHistory(HISTORYTYPE_FOLDER,
			Opt.FoldersHistoryCount,
			"SavedFolderHistory",
			Opt.SaveFoldersHistory,
			true),
		ViewHistory: NewHistory(HISTORYTYPE_VIEW,
			Opt.ViewHistoryCount,
			"SavedFolderHistory",
			Opt.SaveViewHistory,
			true),


	}
	controlObject.FolderHistory.SetAddMode(true, 2, true)
	controlObject.FolderHistory.SetAddMode(true, 1, true)

	if Opt.SaveHistory {
		controlObject.CmdHistory.ReadHistory(false)
	}

	if Opt.SaveFoldersHistory {
		controlObject.FolderHistory.ReadHistory(false)
	}

	if Opt.SaveViewHistory {
		controlObject.ViewHistory.ReadHistory(false)
	}
	logrus.Debugf("[%p] ControlObject::ControlObject()", controlObject)
	return controlObject
}

func (obj *ControlObject) Init() {
	logrus.Debugf("[%p] ControlObject::Init()", obj)

	//TreeList::ClearCache(0);
	//FileFilter::InitFilter();
	//SetColor(COL_COMMANDLINEUSERSCREEN);
	//GotoXY(0, ScrY-3);
	//ShowCopyright();
	//GotoXY(0, ScrY-2);
	///MoveCursor(0, ScrY-1);
	obj.fPanels = newFilePanels()
	obj.CmdLine = newCommandLine()
	//obj.CmdLine.SaveBackground(0, 0, ScrX, ScrY)
	//CmdLine- > SaveBackground(0, 0, ScrX, ScrY);

	//this- > MainKeyBar = &(FPanels- > MainKeyBar);
	//this- > TopMenuBar = &(FPanels- > TopMenuBar);
	//FPanels- > Init();
	//FPanels- > SetScreenPosition();
}

func (obj *ControlObject) Cp() (*FilePanels) {
	return nil
}

func (obj *ControlObject) CreateFilePanels() {

}

func (obj *ControlObject) ShowCopyright(flags uint) {

}
