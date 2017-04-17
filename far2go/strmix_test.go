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

func TestReplaceStrings(t *testing.T) {
	UseCases := []ReplaceStringsTestCase{
		{
			"CTRLALTTAB", "CTRL", "", -1, true, "ALTTAB", 1,
		},
		{
			"CTRL.", "CTRL", "", -1, true, ".", 1,
		},
		{
			"CTRLSHIFT.", "SHIFT", "", -1, true, "CTRL.", 1,
		},
		{
			" - Far %Ver %Build %Platform %Admin", "%Ver", "2.1", -1, true, " - Far 2.1 %Build %Platform %Admin", 1,
		},

	}
	for _, Usecase := range UseCases {

		result, count := ReplaceStrings(Usecase.Source, Usecase.FindStr, Usecase.ReplStr, Usecase.Count, Usecase.IgnoreCase)

		if result != Usecase.ReplacedString {
			t.Errorf("Replace string should be %v but found %v", Usecase.ReplacedString, result)
		}

		if count != Usecase.ReplaceOccurrences {
			t.Errorf("Replace string occurrences should be %v but found %v", Usecase.ReplaceOccurrences, count)
		}
	}

}

type ReplaceStringsTestCase struct {
	Source             string
	FindStr            string
	ReplStr            string
	Count              int
	IgnoreCase         bool
	ReplacedString     string
	ReplaceOccurrences int
}
