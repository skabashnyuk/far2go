package far2go

import (
	"os"
	"log"
	"bufio"
	"strings"
	"fmt"
	"unicode/utf8"
)

type LngErrors uint8

const (
	LERROR_SUCCESS        LngErrors = iota
	LERROR_FILE_NOT_FOUND
	LERROR_BAD_FILE
)

type Language struct {
	lastError      LngErrors
	languageLoaded bool
	msgAddr        string
	msgList        []string
	msgAddrA       string
	msgListA       []string
	msgSize        int
	msgCount       int
	strMessageFile string
	m_bUnicode     bool
}

func (obj *Language) Init(path string) bool {
	if len(obj.msgList) > 0 || len(obj.msgListA) > 0 {
		return true
	}
	scanner := bufio.NewScanner(obj.OpenLangFile(path))
	//scanner.Scan()
	//fmt.Println(/*string(newScanner_a.Bytes())*/scanner.Text(), scanner.Bytes())
	for scanner.Scan() {
		readStr := strings.TrimSpace(scanner.Text())
		r, _ := utf8.DecodeRuneInString(readStr)
		if r != '"' {
			continue
		}
		readStr = readStr[1:len(readStr)-1]
		readStr = obj.ConvertString(readStr)
		obj.msgList = append(obj.msgList, readStr)
		fmt.Println(readStr)
		obj.msgCount++
	}
	return true

}
func (obj *Language) Close() {

}
func (obj *Language) IsLanguageLoaded() {

}
func (obj *Language) GetLastError() {

}
func (obj *Language) ConvertString(src string) string {
	return src
}
func (obj *Language) CheckMsgId() {

}

func (obj *Language) Free() {

}

func (obj *Language) OpenLangFile(path string) *bufio.Reader {
	// Open file and create a buffered reader on top
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	bufio.NewScanner(os.Stdin)
	return bufio.NewReader(file)

}

func (obj *Language) GetLangParam() {

}
func (obj *Language) GetOptionsParam() {

}

func (obj *Language) Select() {

}
