package far2go

import (
	"testing"
)

func TestPaletteType_FarColorToReal(t *testing.T) {
	color, ok := DefaultPalette.FarColorToReal(COL_MENUBOX)
	if !ok {
		t.Error("Not able to get FarColor")
	}
	if color&F_WHITE == 0 {
		t.Error("default color for COL_MENUBOX is not F_WHITE ")
	}
	if color&B_CYAN == 0 {
		t.Error("default color for COL_MENUBOX is not B_CYAN ")
	}
}

func TestPaletteType_FarColorToRealFake(t *testing.T) {
	_, ok := DefaultPalette.FarColorToReal(COL_LASTPALETTECOLOR)
	if ok {
		t.Error("Should not be able to find FarColor for COL_LASTPALETTECOLOR")
	}

}
