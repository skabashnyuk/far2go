package far2go

import (
	"testing"
	"github.com/nicksnyder/go-i18n/i18n"
	"fmt"

)

func TestInit(t *testing.T) {

	i18n.MustLoadTranslationFile("/Users/sj/dev/src/go/src/github.com/skabashnyuk/far2go/far2go/Far.en-US.json")
	i18n.MustLoadTranslationFile("/Users/sj/dev/src/go/src/github.com/skabashnyuk/far2go/far2go/Far.ru-RU.json")

	T, _ := i18n.Tfunc("en-US")
	T2, _ := i18n.Tfunc("ru-RU")
	fmt.Println(T(MConfigSaveFoldersHistory))
	fmt.Println(T2(MConfigSaveFoldersHistory))

}
