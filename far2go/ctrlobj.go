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
	controlObject := &ControlObject{
		fPanels: new(FilePanels),
		CmdLine: new(CommandLine),
		CmdHistory: NewHistory(HISTORYTYPE_CMD,
			512,
			"SavedHistory",
			true,
			false),
		FolderHistory: NewHistory(HISTORYTYPE_FOLDER,
			512,
			"SavedFolderHistory",
			true,
			true),
		ViewHistory: NewHistory(HISTORYTYPE_VIEW,
			512,
			"SavedFolderHistory",
			true,
			true),


	}
	//controlObject.
	logrus.Debugf("[%p] ControlObject::ControlObject()", controlObject)
	return controlObject
}

func (obj *ControlObject) Init() {

}

func (obj *ControlObject) Cp() (*FilePanels) {
	return nil
}

func (obj *ControlObject) CreateFilePanels() {

}

func (obj *ControlObject) ShowCopyright(flags uint) {

}
