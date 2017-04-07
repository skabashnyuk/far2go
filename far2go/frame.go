package far2go

const (
	MODALTYPE_VIRTUAL    = iota
	MODALTYPE_PANELS
	MODALTYPE_VIEWER
	MODALTYPE_EDITOR
	MODALTYPE_DIALOG
	MODALTYPE_VMENU
	MODALTYPE_HELP
	MODALTYPE_COMBOBOX
	MODALTYPE_FINDFOLDER
	MODALTYPE_USER
)

type Frame struct {
	ScreenObject
	frameToBack          *Frame
	prevModal, nextModal *Frame
	dynamicallyBorn      uint
	canLoseFocus         uint
	exitCode             uint
	keyBarVisible        uint
	titleBarVisible      uint
	macroMode            uint

	FrameKeyBar *KeyBar
}

func (obj *Frame) GetCanLoseFocus(DynamicMode uint) (uint) {
	return obj.canLoseFocus
}

func (obj *Frame) SetCanLoseFocus(mode uint) {
	obj.canLoseFocus = mode
}

func (obj *Frame) GetExitCode() (uint) {
	return obj.exitCode
}

func (obj *Frame) SetExitCode(Code uint) {
	obj.exitCode = Code
}

func (obj *Frame) IsFileModified() (bool) {
	return false
}

func (obj *Frame) GetTypeName() (string) {
	return "[FarModal]"
}

func (obj *Frame) GetTypeAndName(strType string, strName string) (uint) {
	return MODALTYPE_VIRTUAL
}

func (obj *Frame) GetType() (uint) {
	return MODALTYPE_VIRTUAL
}

func (obj *Frame) OnDestroy() {

}

func (obj *Frame) OnCreate() {

}

func (obj *Frame) OnChangeFocus(focus uint) {

}

func (obj *Frame) Refresh() {
	obj.OnChangeFocus(1)
}

func (obj *Frame) InitKeyBar() {
}

func (obj *Frame) SetKeyBar(frameKeyBar *KeyBar) {

}

func (obj *Frame) UpdateKeyBar() {

}

func (obj *Frame) RedrawKeyBar() {
	obj.UpdateKeyBar()
}

func (obj *Frame) IsTitleBarVisible() (uint) {
	return obj.titleBarVisible
}

func (obj *Frame) IsTopFrame() (uint) {
	return obj.titleBarVisible
}

func (obj *Frame) GetMacroMode() (uint) {
	return obj.macroMode
}

func (obj *Frame) Push(Modalized *Frame) {

}

func (obj *Frame) GetTopModal() (*Frame) {
	return obj.nextModal
}

func (obj *Frame) DestroyAllModal() {
}

func (obj *Frame) SetDynamicallyBorn(born uint) {
	obj.dynamicallyBorn = born
}

func (obj *Frame) GetDynamicallyBorn() (uint) {
	return obj.dynamicallyBorn
}

func (obj *Frame) FastHide() (uint) {
	return 0
}

func (obj *Frame) RemoveModal(aFrame *Frame) (bool) {
	return true
}

func (obj *Frame) ResizeConsole() {

}

func (obj *Frame) HasSaveScreen() {

}

func (obj *Frame) GetTitle(title string) (string) {
	return title
}

func (obj *Frame) ProcessEvents() (bool) {
	return true
}
