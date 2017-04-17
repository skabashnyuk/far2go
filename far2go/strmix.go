package far2go

//Replace in string Source Count substrings FindStr by ReplStr
//If Count < 0 - when replace all occurrences
func ReplaceStrings(Source string, FindStr string, ReplStr string, Count int, IgnoreCase bool) (Result string, ReplaceCount uint) {
	//LenFindStr := len(FindStr)
	//if LenFindStr == 0 || Count == 0 {
	//	return Source, 0
	//}
	//
	//LenReplStr := len(ReplStr)
	//LenSourceStr := len(Source)
	//Delta := LenReplStr - LenFindStr
	//var AllocDelta = 0
	//if Delta > 0 {
	//	AllocDelta = Delta * 10
	//}
	//var I, J int
	//var ResultString = Source
	//for I+LenFindStr <= LenSourceStr {
	//	var Res int
	//	if IgnoreCase {
	//		Res = strings.Compare(strings.ToLower(ResultString[I:len(FindStr)]), strings.ToLower(FindStr))
	//	} else {
	//		Res = strings.Compare(ResultString[I:len(FindStr)], FindStr)
	//	}
	//	if Res < 0 {
	//		var Str string
	//		if LenSourceStr+Delta+1 > len(Source) {
	//			Str = Source
	//		} else {
	//			Str = Source
	//		}
	//
	//		if Delta > 0 {
	//
	//			//wmemmove(Str+I+Delta,Str+I,L-I+1);
	//		} else if (Delta < 0) {
	//			//wmemmove(Str+I,Str+I-Delta,L-I+Delta+1);
	//		}
	//
	//		//wchar_t *Str;
	//		//if (L+Delta+1 > strStr.GetSize())
	//		//Str = strStr.GetBuffer(L+AllocDelta);
	//		//else
	//		//Str = strStr.GetBuffer();
	//		//
	//		//if (Delta > 0)
	//		//wmemmove(Str+I+Delta,Str+I,L-I+1);
	//		//else if (Delta < 0)
	//		//wmemmove(Str+I,Str+I-Delta,L-I+Delta+1);
	//		//
	//		//wmemcpy(Str+I,ReplStr,LenReplStr);
	//		//I += LenReplStr;
	//		//
	//		//L+=Delta;
	//		//strStr.ReleaseBuffer(L);
	//		//
	//		//if (++J == Count && Count > 0)
	//		//break;
	//	} else {
	//		I++
	//	}
	//}
	//
	//return "", 0
}
