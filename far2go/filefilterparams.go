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

//void SetTitle(const wchar_t *Title);
//void SetMask(bool Used, const wchar_t *Mask);
//void SetDate(bool Used, DWORD DateType, FILETIME DateAfter, FILETIME DateBefore, bool bRelative);
//void SetSize(bool Used, const wchar_t *SizeAbove, const wchar_t *SizeBelow);
//void SetAttr(bool Used, DWORD AttrSet, DWORD AttrClear);
//void SetColors(HighlightDataColor *Colors);
//void SetSortGroup(int SortGroup) { FHighlight.SortGroup = SortGroup; }
//void SetContinueProcessing(bool bContinueProcessing) { FHighlight.bContinueProcessing = bContinueProcessing; }
//void SetFlags(enumFileFilterFlagsType FType, DWORD Flags) { FFlags[FType] = Flags; }
//void ClearAllFlags() { memset(FFlags,0,sizeof(FFlags)); }
//
//const wchar_t *GetTitle() const;
//bool  GetMask(const wchar_t **Mask) const;
//bool  GetDate(DWORD *DateType, FILETIME *DateAfter, FILETIME *DateBefore, bool *bRelative) const;
//bool  GetSize(const wchar_t **SizeAbove, const wchar_t **SizeBelow) const;
//bool  GetAttr(DWORD *AttrSet, DWORD *AttrClear) const;
//void  GetColors(HighlightDataColor *Colors) const;
//int   GetMarkChar() const;
//int   GetSortGroup() const { return FHighlight.SortGroup; }
//bool  GetContinueProcessing() const { return FHighlight.bContinueProcessing; }
//DWORD GetFlags(enumFileFilterFlagsType FType) const { return FFlags[FType]; }

//bool FileInFilter(const FAR_FIND_DATA_EX& fde, uint64_t CurrentTime);
//bool FileInFilter(const FAR_FIND_DATA& fd, uint64_t CurrentTime);
//bool FileFilterConfig(FileFilterParams *FF, bool ColorConfig=false);
//void MenuString(FARString &strDest, FileFilterParams *FF, bool bHighlightType=false, int Hotkey=0, bool bPanelType=false, const wchar_t *FMask=nullptr, const wchar_t *Title=nullptr);

