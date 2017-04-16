package far2go

//Plugin API
const
(
	XLAT_SWITCHKEYBLAYOUT  = 0x00000001
	XLAT_SWITCHKEYBBEEP    = 0x00000002
	XLAT_USEKEYBLAYOUTNAME = 0x00000004
	XLAT_CONVERTALLCMDLINE = 0x00010000
)

type FarListItem struct {
	Flags    int
	Text     string
	Reserved [3]int
}

type FarList struct {
	ItemsNumber int
	Items       *FarListItem
}
