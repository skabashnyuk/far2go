package far2go

//Plugin API

type FarListItem struct {
	Flags    int
	Text     string
	Reserved [3]int
}

type FarList struct {
	ItemsNumber int
	Items       *FarListItem
}
