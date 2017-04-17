package far2go

import (
	"strings"
	"github.com/sirupsen/logrus"
)

//Replace in string Source Count substrings FindStr by ReplStr
//If Count < 0 - when replace all occurrences
func ReplaceStrings(Source string, FindStr string, ReplStr string, Count int, IgnoreCase bool) (Result string, ReplaceCount int) {
	var LenFindStr = len(FindStr)
	if LenFindStr == 0 || Count == 0 {
		return Source, 0
	}
	var ResultString string
	if IgnoreCase {
		ResultString = strings.Replace(Source, FindStr, ReplStr, Count)
	} else {
		ResultString = strings.Replace(Source, FindStr, ReplStr, Count)
	}
	Occurrences := strings.Count(Source, FindStr)
	ReplaceDelta := len(FindStr) - len(ReplStr)
	ResultDelta := len(Source) - len(ResultString)

	var ActualReplaceCount int
	if Count < 0 {
		ActualReplaceCount = Occurrences
	} else {
		if Count < Occurrences {
			ActualReplaceCount = Count
		} else {
			ActualReplaceCount = Occurrences
		}
	}
	logrus.WithFields(logrus.Fields{
		"Source":             Source,
		"FindStr":            FindStr,
		"ReplStr":            ReplStr,
		"Count":              Count,
		"IgnoreCase":         IgnoreCase,
		"Occurrences":        Occurrences,
		"ReplaceDelta":       ReplaceDelta,
		"ResultDelta":        ResultDelta,
		"ResultString":       ResultString,
		"ActualReplaceCount": ActualReplaceCount,

	}).Info("Replace")
	return ResultString, ActualReplaceCount

}
