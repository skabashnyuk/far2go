package far2go

import (
	"github.com/sirupsen/logrus"
	"strings"
)

var (
	AltPressed        bool
	CtrlPressed       bool
	ShiftPressed      bool
	RightAltPressed   bool
	RightCtrlPressed  bool
	RightShiftPressed bool
)

type TFKey3 struct {
	Key         KeyCode
	Len         int
	Name, Uname string
}

var ModifKeyName []TFKey3 = []TFKey3{
	TFKey3{KEY_RCTRL, 5, "RCtrl", "RCTRL"},
	TFKey3{KEY_SHIFT, 5, "Shift", "SHIFT"},
	TFKey3{KEY_CTRL, 4, "Ctrl", "CTRL"},
	TFKey3{KEY_RALT, 4, "RAlt", "RALT"},
	TFKey3{KEY_ALT, 3, "Alt", "ALT"},
	TFKey3{KEY_M_SPEC, 4, "Spec", "SPEC"},
	TFKey3{KEY_M_OEM, 3, "Oem", "OEM"},
}

func KeyNameToKey(Name string) int {
	if len(Name) < 1 {
		return -1
	}
	logrus.Debugf("KeyNameToKey('%ls')", Name)

	//is this a macro key
	if len(Name) > 2 {
		if firstChar := string([]rune(Name)[0]); firstChar == "$" || firstChar == "%" || firstChar == "(" || firstChar == ")" {
			return -1
		}

	}

	//if (Name[1] && wcspbrk(Name,L"()")) // если не один символ и встречаются '(' или ')', то это явно не клавиша!
	//return -1;
	//
	////   if((Key=KeyNameMacroToKey(Name)) != (DWORD)-1)
	////     return Key;
	//int I, Pos;
	var Key KeyCode
	//var I, Pos int
	var strTmpName = Name
	strTmpName = strings.ToUpper(strTmpName)
	//static FARString strTmpName;
	//strTmpName = Name;
	//strTmpName.Upper();
	//int Len=(int)strTmpName.GetLength();
	//var Len = len(strTmpName)

	for _, element := range ModifKeyName {
		//
		if strings.HasPrefix(strTmpName, element.Uname) && Key&element.Key != 0 {

		}
	}

	//
	//// пройдемся по всем модификаторам
	//for (Pos=I=0; I < int(ARRAYSIZE(ModifKeyName)); ++I)
	//{
	//if (wcsstr(strTmpName,ModifKeyName[I].UName) && !(Key&ModifKeyName[I].Key))
	//{
	//int CntReplace=ReplaceStrings(strTmpName,ModifKeyName[I].UName,L"",-1,TRUE);
	//Key|=ModifKeyName[I].Key;
	//Pos+=ModifKeyName[I].Len*CntReplace;
	//}
	//}
	//// _SVS(SysLog(L"[%d] Name=%ls",__LINE__,Name));
	//
	////Pos=strlen(TmpName);
	//
	//// если что-то осталось - преобразуем.
	//if (Pos < Len)
	//{
	//	// сначала - FKeys1 - Вариант (1)
	//	const wchar_t* Ptr=Name+Pos;
	//	int PtrLen = Len-Pos;
	//
	//	for (I=(int)ARRAYSIZE(FKeys1)-1; I>=0; I--)
	//	{
	//	if (PtrLen == FKeys1[I].Len && !StrCmpI(Ptr,FKeys1[I].Name))
	//	{
	//	Key|=FKeys1[I].Key;
	//	Pos+=FKeys1[I].Len;
	//	break;
	//	}
	//	}
	//
	//	if (I == -1) // F-клавиш нет?
	//	{
	//		/*
	//			здесь только 5 оставшихся вариантов:
	//			2) Опциональные модификаторы (Alt/RAlt/Ctrl/RCtrl/Shift) и 1 символ, например, AltD или CtrlC
	//			3) "Alt" (или RAlt) и 5 десятичных цифр (с ведущими нулями)
	//			4) "Spec" и 5 десятичных цифр (с ведущими нулями)
	//			5) "Oem" и 5 десятичных цифр (с ведущими нулями)
	//			6) только модификаторы (Alt/RAlt/Ctrl/RCtrl/Shift)
	//		*/
	//
	//		if (Len == 1 || Pos == Len-1) // Вариант (2)
	//		{
	//			int Chr=Name[Pos];
	//
	//			// если были модификаторы Alt/Ctrl, то преобразуем в "физичекую клавишу" (независимо от языка)
	//			if (Key&(KEY_ALT|KEY_RCTRL|KEY_CTRL|KEY_RALT))
	//			{
	//				if (Chr > 0x7F)
	//				Chr=KeyToKeyLayout(Chr);
	//
	//				Chr=Upper(Chr);
	//			}
	//
	//			Key|=Chr;
	//
	//			if (Chr)
	//			Pos++;
	//		}
	//else if (Key == KEY_ALT || Key == KEY_RALT || Key == KEY_M_SPEC || Key == KEY_M_OEM) // Варианты (3), (4) и (5)
	//{
	//wchar_t *endptr=nullptr;
	//int K=(int)wcstol(Ptr, &endptr, 10);
	//
	//if (Ptr+5 == endptr)
	//{
	//if (Key == KEY_ALT || Key == KEY_RALT) // Вариант (3) - Alt-Num
	//Key=(Key|K|KEY_ALTDIGIT)&(~(KEY_ALT|KEY_RALT));
	//else if (Key == KEY_M_SPEC) // Вариант (4)
	//Key=(Key|(K+KEY_VK_0xFF_BEGIN))&(~(KEY_M_SPEC|KEY_M_OEM));
	//else if (Key == KEY_M_OEM) // Вариант (5)
	//Key=(Key|(K+KEY_FKEY_BEGIN))&(~(KEY_M_SPEC|KEY_M_OEM));
	//
	//Pos=Len;
	//}
	//}
	//// Вариант (6). Уже "собран".
	//}
	//}
	//
	///*
	//if(!(Key&(KEY_ALT|KEY_RCTRL|KEY_CTRL|KEY_RALT|KEY_ALTDIGIT)) && (Key&KEY_SHIFT) && LocalIsalpha(Key&(~KEY_CTRLMASK)))
	//{
	//	Key&=~KEY_SHIFT;
	//	Key=LocalUpper(Key);
	//}
	//*/
	//// _SVS(SysLog(L"Key=0x%08X (%c) => '%ls'",Key,(Key?Key:' '),Name));
	//return (!Key || Pos < Len)? -1: (int)Key;
	return 0
}
