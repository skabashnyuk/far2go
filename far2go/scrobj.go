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
	ObjWidth, ObjHeight uint
	nLockCount          uint
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

func (obj *ScreenObject) Redraw() {
	if obj.flags.Check(FSCROBJ_VISIBLE) {
		obj.Show()
	}
}

func (obj *ScreenObject) IsVisible() bool {
	return obj.flags.Check(FSCROBJ_VISIBLE)
}

func (obj *ScreenObject) SetVisible(visible bool) {
	obj.flags.Change(FSCROBJ_VISIBLE, visible)
}

func (obj *ScreenObject) SetRestoreScreenMode(mode bool) {
	obj.flags.Change(FSCROBJ_ENABLERESTORESCREEN, mode)
}

func (obj *ScreenObject) Shadow(full bool) {

	if obj.flags.Check(FSCROBJ_VISIBLE) {
		if full {
			//if (!ShadowSaveScr)
			//ShadowSaveScr=new SaveScreen(0,0,ScrX,ScrY);
			//
			//MakeShadow(0,0,ScrX,ScrY);

		} else {
			//if (!ShadowSaveScr)
			//ShadowSaveScr=new SaveScreen(X1,Y1,X2+2,Y2+1);
			//
			//MakeShadow(X1+2,Y2+1,X2+1,Y2+1);
			//MakeShadow(X2+1,Y1+1,X2+2,Y2+1);

		}
	}
}

func (obj *ScreenObject) SetPosition(c1, c2 Coordinate) {
	//if (SaveScr)
	//{
	//	delete SaveScr;
	//	SaveScr = nullptr;
	//}

	obj.c1 = c1
	obj.c2 = c2
	obj.ObjWidth = c2.X - c1.X + 1
	obj.ObjHeight = c2.Y - c1.Y + 1
	obj.flags.Set(FSCROBJ_SETPOSITIONDONE)

}

func (obj *ScreenObject) SetScreenPosition() {
	obj.flags.Clear(FSCROBJ_SETPOSITIONDONE)
}

func (obj *ScreenObject) GetPosition() (c1, c2 Coordinate) {
	return c1, c2
}

func (obj *ScreenObject) Hide() {
	if obj.flags.Check(FSCROBJ_VISIBLE) {
		obj.flags.Clear(FSCROBJ_VISIBLE)

		//if (ShadowSaveScr)
		//{
		//	delete ShadowSaveScr;
		//	ShadowSaveScr=nullptr;
		//}
		//
		//if (SaveScr)
		//{
		//	delete SaveScr;
		//	SaveScr=nullptr;
		//}
	}
}

func (obj *ScreenObject) Show() {
	if !obj.Locked() {
		if obj.flags.Check(FSCROBJ_SETPOSITIONDONE) {
			obj.SavePrevScreen()
			obj.DisplayObject()
		}
	}
}
func (obj *ScreenObject) DisplayObject() {

}
