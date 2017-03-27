package far2go

type SaveScreen struct {
}

const
(
	FSCROBJ_VISIBLE             uint32 = 1 << iota
	FSCROBJ_ENABLERESTORESCREEN
	FSCROBJ_SETPOSITIONDONE
	FSCROBJ_ISREDRAWING
)

type ScreenObject struct {
	flags BitFlags
	//shadowSaveScr       *SaveScreen
	c1, c2              Coordinate
	ObjWidth, ObjHeight int
	nLockCount          int
	pOwner              *ScreenObject
	//SaveScr             *SaveScreen
	CaptureMouseObject *ScreenObject
}

func (obj *ScreenObject) Lock() {
	obj.nLockCount++
}

func (obj *ScreenObject) Unlock() {
	if obj.nLockCount > 0 {
		obj.nLockCount--
	} else {
		obj.nLockCount = 0
	}
}

func (obj *ScreenObject) Locked() bool {
	a := func() bool {
		if obj.pOwner != nil {
			return obj.pOwner.Locked()
		} else {
			return false
		}
	}()

	return obj.nLockCount > 0 || a
}

func (obj *ScreenObject) SetOwner(pOwner *ScreenObject) {
	obj.pOwner = pOwner
}

func (obj *ScreenObject) GetOwner() (pOwner *ScreenObject) {
	return obj.pOwner
}

func (obj *ScreenObject) SavePrevScreen() {

	if obj.flags.Check(FSCROBJ_SETPOSITIONDONE) {
		if !obj.flags.Check(FSCROBJ_VISIBLE) {
			obj.flags.Set(FSCROBJ_VISIBLE)
			if obj.flags.Check(FSCROBJ_ENABLERESTORESCREEN) /*&& obj.SaveScr != nil */ {
				//SaveScr=new SaveScreen(X1,Y1,X2,Y2)
			}

		}
	}
	//if (!)
	//return;
	//if (!Flags.Check(FSCROBJ_VISIBLE)) {
	//	Flags.Set(FSCROBJ_VISIBLE);
	//	if (Flags.Check(FSCROBJ_ENABLERESTORESCREEN) && !SaveScr)
	//	SaveScr = new
	//	SaveScreen(X1, Y1, X2, Y2);
	//}
}

//void SavePrevScreen();
//void Redraw();
//int  IsVisible() { return Flags.Check(FSCROBJ_VISIBLE); };
//void SetVisible(int Visible) {Flags.Change(FSCROBJ_VISIBLE,Visible);};
//void SetRestoreScreenMode(int Mode) {Flags.Change(FSCROBJ_ENABLERESTORESCREEN,Mode);};
//void Shadow(bool Full=false);
