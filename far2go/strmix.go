package far2go

//Replace in string Source Count substrings FindStr by ReplStr
//If Count < 0 - when replace all occurrences
func ReplaceStrings(Source string, FindStr string, ReplStr string, Count int, IgnoreCase bool) (Result string, ReplaceCount uint) {
	var LenFindStr = len(FindStr)
	if LenFindStr == 0 || Count == 0 {
		return Source, 0
	}
	return "", 0
}
