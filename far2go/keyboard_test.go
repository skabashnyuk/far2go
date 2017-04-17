package far2go

import "testing"

func TestKeyNameToKeyOnEmptyString(t *testing.T) {
	if KeyNameToKey("") != -1 {
		t.Error("Not able to test empty string")
	}
}

func TestKeyNameOnMacroDollar(t *testing.T) {
	if KeyNameToKey("$sd") != -1 {
		t.Error("Not able to detect macro")
	}
}

func TestKeyNameOnMacroPercent(t *testing.T) {
	if KeyNameToKey("%sd") != -1 {
		t.Error("Not able to detect macro")
	}
}

func TestKeyNameOnMacroOpenBracket(t *testing.T) {
	if KeyNameToKey("(sd") != -1 {
		t.Error("Not able to detect macro")
	}
}

func TestKeyNameOnMacroCloseBracket(t *testing.T) {
	if KeyNameToKey(")sd") != -1 {
		t.Error("Not able to detect macro")
	}
}
