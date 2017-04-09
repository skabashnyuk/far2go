package far2go

import "time"

const FILEFILTER_SIZE_SIZE = 32
const DEFAULT_SORT_GROUP = 10000

const (
	FFFT_FIRST = iota

	FFFT_LEFTPANEL
	FFFT_RIGHTPANEL
	FFFT_FINDFILE
	FFFT_COPY
	FFFT_SELECT
	FFFT_CUSTOM

	FFFT_COUNT
)
const (
	FFF_NONE    = 0x00000000
	FFF_INCLUDE = 0x00000001
	FFF_EXCLUDE = 0x00000002
	FFF_STRONG  = 0x10000000
)

type DateTypeConst uint

const (
	FDATE_MODIFIED DateTypeConst = iota
	FDATE_CREATED
	FDATE_OPENED
	FDATE_CHANGED
	FDATE_COUNT
)

type MaskType struct {
	Used       bool
	Mask       string
	FilterMask FilterMask
}

type DateType struct {
	Used       bool
	DateType   DateTypeConst
	DateAfter  time.Time
	DateBefore time.Time
	Relative   bool
}
type SizeType struct {
	Used                         bool
	SizeAboveReal, SizeBelowReal uint
	SizeAbove, SizeBelow         string
}

type AttrType struct {
	Used               bool
	AttrSet, AttrClear uint64
}
type HighlightType struct {
	Colors             HighlightDataColor
	SortGroup          int
	ContinueProcessing bool
}
type FileFilterParams struct {
	title     string
	flags     [FFFT_COUNT]int
	mask      MaskType
	date      DateType
	size      SizeType
	attr      AttrType
	highlight HighlightType
}
