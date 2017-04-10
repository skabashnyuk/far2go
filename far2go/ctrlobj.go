package far2go

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
	Macro          KeyMacro
	//Plugins        PluginManager
}

func NewControlObject() (*ControlObject) {
	return nil
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
