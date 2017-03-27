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
	altState, ctrlState, shiftState int
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

//void Set(const wchar_t * const *Key,int KeyCount)            { SetGroup(KBL_MAIN, Key, KeyCount); }
//void SetShift(const wchar_t * const *Key,int KeyCount)       { SetGroup(KBL_SHIFT, Key, KeyCount); }
//void SetAlt(const wchar_t * const *Key,int KeyCount)         { SetGroup(KBL_ALT, Key, KeyCount); }
//void SetCtrl(const wchar_t * const *Key,int KeyCount)        { SetGroup(KBL_CTRL, Key, KeyCount); }
//void SetCtrlShift(const wchar_t * const *Key,int KeyCount)   { SetGroup(KBL_CTRLSHIFT, Key, KeyCount); }
//void SetAltShift(const wchar_t * const *Key,int KeyCount)    { SetGroup(KBL_ALTSHIFT, Key, KeyCount); }
//void SetCtrlAlt(const wchar_t **Key,int KeyCount)            { SetGroup(KBL_CTRLALT, Key, KeyCount); }
//void SetCtrlAltShift(const wchar_t **Key,int KeyCount)       { SetGroup(KBL_CTRLALTSHIFT, Key, KeyCount); }
