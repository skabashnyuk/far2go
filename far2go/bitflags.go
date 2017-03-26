package far2go

type BitFlags struct {
	flags uint32
}

func (flags *BitFlags) Set(newFlags uint32) uint32 {
	flags.flags |= newFlags
	return flags.flags
}

func (flags *BitFlags) Clear(newFlags uint32) uint32 {
	flags.flags &^= newFlags
	return flags.flags
}

func (flags *BitFlags) Check(newFlags uint32) bool {
	return flags.flags&newFlags != 0
}

func (flags *BitFlags) Change(newFlags uint32, status bool) uint32 {
	if status {
		flags.Set(newFlags)
	} else {
		flags.Clear(newFlags)
	}

	return flags.flags
}

func (flags *BitFlags) Swap(swappedFlags uint32) uint32 {
	if flags.Check(swappedFlags) {
		flags.Clear(swappedFlags)
	} else {
		flags.Set(swappedFlags)
	}
	return flags.flags
}

func (flags *BitFlags) ClearAll() {
	flags.flags = 0
}
