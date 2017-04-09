package far2go

const MAX_POSITIONS = 512
const BOOKMARK_COUNT = 10

type PosCache struct {
	Param    [5]uint64
	Position *[4]uint64
}

type FilePositionCache struct {
	IsMemory bool
	CurPos   int
	Names    string
	Param    *byte
	Position *byte
}

//
//int FindPosition(const wchar_t *FullName);
//void AddPosition(const wchar_t *Name,PosCache& poscache);
//bool GetPosition(const wchar_t *Name,PosCache& poscache);
//
//bool Read(const wchar_t *Key);
//bool Save(const wchar_t *Key);
