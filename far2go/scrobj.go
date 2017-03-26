package far2go

type SaveScreen int

const
(
	FSCROBJ_VISIBLE             SaveScreen = 1 << iota
	FSCROBJ_ENABLERESTORESCREEN
	FSCROBJ_SETPOSITIONDONE
	FSCROBJ_ISREDRAWING
)



type ScreenObject struct {
	Flags               BitFlags
	ShadowSaveScr       *SaveScreen
	X1, Y1, X2, Y2      int
	ObjWidth, ObjHeight int
	nLockCount          int

	SaveScr            *SaveScreen
	CaptureMouseObject *ScreenObject
}
