package far2go

import (
	"time"
	"container/list"
)

type HistoryType uint

const (
	HISTORYTYPE_CMD    HistoryType = iota
	HISTORYTYPE_FOLDER
	HISTORYTYPE_VIEW
	HISTORYTYPE_DIALOG
)

type HistoryRecord struct {
	Lock      bool
	Type      int
	strName   string
	Timestamp time.Time
}

type History struct {
	RegKey                               string
	EnableAdd, KeepSelectedPos, SaveType bool
	RemoveDups                           int
	TypeHistory                          HistoryType
	HistoryCount                         uint
	EnableSave                           int
	HistoryList                          list.List
	CurrentItem                          HistoryRecord
}

//void AddToHistory(const wchar_t *Str, int Type=0, const wchar_t *Prefix=nullptr, bool SaveForbid=false);
//bool ReadHistory(bool bOnlyLines=false);
//bool SaveHistory();
//static bool ReadLastItem(const wchar_t *RegKey, FARString &strStr);
//int  Select(const wchar_t *Title, const wchar_t *HelpTopic, FARString &strStr, int &Type);
//int  Select(VMenu &HistoryMenu, int Height, Dialog *Dlg, FARString &strStr);
//void GetPrev(FARString &strStr);
//void GetNext(FARString &strStr);
//bool GetSimilar(FARString &strStr, int LastCmdPartLength, bool bAppend=false);
//bool GetAllSimilar(VMenu &HistoryMenu,const wchar_t *Str);
//void SetAddMode(bool EnableAdd, int RemoveDups, bool KeepSelectedPos);
//void ResetPosition() { CurrentItem = nullptr; }
