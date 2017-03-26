package widget

import (
	"testing"
	"strings"
)

func TestShouldParseHotkeyIfItExist(t *testing.T) {
	hotkey := parse_hotkey("Vie&w file...")
	if strings.Compare(hotkey.start, "Vie") != 0 {
		t.Errorf("start '%q' is not expected", hotkey.start)
	}
	if strings.Compare(hotkey.hotkey, "w") != 0 {
		t.Errorf(" Hotkey '%q' is not expected", hotkey.hotkey)
	}
	if strings.Compare(hotkey.end, " file...") != 0 {
		t.Errorf(" end '%q' is not expected.", hotkey.end)
	}
}


func TestShouldParseHotkeyIfItNotExist(t *testing.T) {
	hotkey := parse_hotkey("View file...")
	if strings.Compare(hotkey.start, "View file...") != 0 {
		t.Errorf("start '%q' is not expected", hotkey.start)
	}
	if strings.Compare(hotkey.hotkey, "") != 0 {
		t.Errorf(" Hotkey '%q' is not expected", hotkey.hotkey)
	}
	if strings.Compare(hotkey.end, "") != 0 {
		t.Errorf(" end '%q' is not expected.", hotkey.end)
	}
}
