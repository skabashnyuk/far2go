package far2go

const (
	KBL_MAIN         = iota
	KBL_SHIFT
	KBL_CTRL
	KBL_ALT
	KBL_CTRLSHIFT
	KBL_ALTSHIFT
	KBL_CTRLALT
	KBL_CTRLALTSHIFT

	KBL_GROUP_COUNT
)
const KEY_COUNT int = 12

type KeyBar struct {
	ScreenObject
	owner                           *ScreenObject
	keyTitles                       [KBL_GROUP_COUNT][]string
	keyCounts                       [KBL_GROUP_COUNT]int
	altState, ctrlState, shiftState bool
	disableMask                     int
	regKeyTitles                    [KBL_GROUP_COUNT]string
	regReaded                       bool
	strLanguage                     string
	strRegGroupName                 string
}

func (obj *KeyBar) SetGroup(group int, key []string, keyCount int) {
	if key != nil {
		for i := 0; i < keyCount && i < KEY_COUNT; i++ {
			if key[i] != "" {
				obj.keyTitles[group][i] = key[i]
			}
		}
		obj.keyCounts[group] = keyCount
	}
}
func (obj *KeyBar) Set(key []string, keyCount int) {
	obj.SetGroup(KBL_MAIN, key, keyCount)
}

func (obj *KeyBar) SetShift(key []string, keyCount int) {
	obj.SetGroup(KBL_SHIFT, key, keyCount)
}
func (obj *KeyBar) SetAlt(key []string, keyCount int) {
	obj.SetGroup(KBL_ALT, key, keyCount)
}
func (obj *KeyBar) SetCtrl(key []string, keyCount int) {
	obj.SetGroup(KBL_CTRL, key, keyCount)
}
func (obj *KeyBar) SetCtrlShift(key []string, keyCount int) {
	obj.SetGroup(KBL_CTRLSHIFT, key, keyCount)
}
func (obj *KeyBar) SetAltShift(key []string, keyCount int) {
	obj.SetGroup(KBL_ALTSHIFT, key, keyCount)
}
func (obj *KeyBar) SetCtrlAlt(key []string, keyCount int) {
	obj.SetGroup(KBL_CTRLALT, key, keyCount)
}
func (obj *KeyBar) SetCtrlAltShift(key []string, keyCount int) {
	obj.SetGroup(KBL_CTRLALTSHIFT, key, keyCount)
}

func (obj *KeyBar) RedrawIfChanged() {
	if ShiftPressed != obj.shiftState ||
		CtrlPressed != obj.ctrlState ||
		AltPressed != obj.altState {
		obj.Redraw()
	}
}

func (obj *KeyBar) ProcessKey(key BaseDefKeyboard) bool {
	switch key {
	case KEY_KILLFOCUS:
	case KEY_GOTFOCUS:
		obj.RedrawIfChanged()
		return true

	}
	return false
}

func (obj *KeyBar) ReadRegGroup(regGroup string, language string) {

}
func (obj *KeyBar) SetRegGroup(group int) {

}

func (obj *KeyBar) SetAllGroup(group int, startIndex int, count int) {

}
func (obj *KeyBar) SetAllRegGroup() {

}
func (obj *KeyBar) ClearGroup(group int) {

}
func (obj *KeyBar) SetDisableMask(mask int) {

}
func (obj *KeyBar) Change(newstr []string, pos int) {
	obj.SetGroup(KBL_MAIN, newstr, pos)

}
func (obj *KeyBar) ResizeConsole() {

}
