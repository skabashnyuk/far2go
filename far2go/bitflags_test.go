package far2go

import "testing"

func TestToCheckExecSet(t *testing.T) {
	//given
	var flags BitFlags = BitFlags{2}
	//when
	flags.Set(1)
	//then
	if !flags.Check(3) {
		t.Error("Not able to set or check")
	}

}

func TestOthefCheck(t *testing.T) {
	//given
	var flags BitFlags = BitFlags{2}
	//when
	flags.Set(1)
	//then

	if flags.Check(4) {
		t.Error("Not able to set or check")
	}
}

func TestClearAll(t *testing.T) {
	//given
	var flags BitFlags = BitFlags{19}
	//when
	flags.ClearAll()
	//then

	if flags.Check(19) {
		t.Error("Not able to ClearAll")
	}
}

func TestChangeWithTrueFlag(t *testing.T) {
	//given
	var flags BitFlags = BitFlags{5}
	//when
	flags.Change(1, true)
	//then
	//
	if !flags.Check(4) {
		t.Error("Not able to Change with true flag")
	}
}

func TestChangeWithFalseFlag(t *testing.T) {
	//given
	var flags BitFlags = BitFlags{5}
	//when
	flags.Change(1, false)
	//then
	//
	if !flags.Check(5) {
		t.Error("Not able to Change with false flag")
	}
}

func TestSwapAndSet(t *testing.T) {
	//given
	var flags BitFlags = BitFlags{5}
	//when
	flags.Swap(2)
	//then
	//
	if !flags.Check(7) {
		t.Error("Not able to Swap and Set")
	}
}

func TestSwapAndClear(t *testing.T) {
	//given
	var flags BitFlags = BitFlags{5}
	//when
	flags.Swap(1)
	//then
	//
	if !flags.Check(4) {
		t.Error("Not able to Swap and Set")
	}
}

func TestClear(t *testing.T) {
	//given
	var flags BitFlags = BitFlags{873}
	//when
	flags.Clear(512)
	//then
	//
	if !flags.Check(361) {
		t.Error("Not able to Clear")
	}
}
