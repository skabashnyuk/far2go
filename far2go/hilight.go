package far2go

import "time"

type HighlightDataColor struct {
	Color    [2][4]uint // [0=file, 1=mark][0=normal,1=selected,2=undercursor,3=selectedundercursor]; if HIBYTE == 0xFF then transparent
	MarkChar uint64
}
type HighlightFiles struct {
	HiData                                        []FileFilterParams
	FirstCount, UpperCount, LowerCount, LastCount int
	CurrentTime                                   time.Time
}

//
//void InitHighlightFiles();
//void ClearData();
//
//int  MenuPosToRealPos(int MenuPos, int **Count, bool Insert=false);
//void FillMenu(VMenu *HiMenu,int MenuPos);
//void ProcessGroups();
//void UpdateCurrentTime();
//void GetHiColor(FileListItem **FileItem,int FileCount,bool UseAttrHighlighting=false);
//int  GetGroup(const FileListItem *fli);
//void HiEdit(int MenuPos);
//
//void SaveHiData();
//void SetHighlighting();
