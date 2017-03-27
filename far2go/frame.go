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
	frameToBack          *Frame
	prevModal, nextModal *Frame
	dynamicallyBorn      int
	canLoseFocus         int
	exitCode             int
	keyBarVisible        int
	titleBarVisible      int
	macroMode            int

	//FrameKeyBar *KeyBar ;
}
