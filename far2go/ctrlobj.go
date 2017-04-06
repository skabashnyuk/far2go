package far2go

type ControlObject struct {
	fPanels     *FilePanels
	CmdLine     *CommandLine
	CmdHistory  *History
	ViewHistory *FolderHistory

	MainKeyBar     *KeyBar
	TopMenuBar     *MenuBar
	HiFiles        *HighlightFiles
	ViewerPosCache *FilePositionCache
	EditorPosCache *FilePositionCache
	Macro          KeyMacro
	Plugins        PluginManager
}

func (obj *ControlObject) Init() {

}

func (obj *ControlObject) Cp() (*FilePanels) {

}

func (obj *ControlObject) CreateFilePanels() {

}

func (obj *ControlObject) ShowCopyright(flags uint) {

}
