package far2go

import "testing"

func TestReplaceStringsEmptyFindString(t *testing.T) {
	result, count := ReplaceStrings("source string", "", "repl", 2, false)

	if result != "source string" {
		t.Error("Source string should be the same")
	}

	if count != 0 {
		t.Error("Replace count should be 0")
	}
}

func TestReplaceStringsZeroCount(t *testing.T) {
	result, count := ReplaceStrings("source string", "source", "repl", 0, false)

	if result != "source string" {
		t.Error("Source string should be the same")
	}

	if count != 0 {
		t.Error("Replace count should be 0")
	}
}
