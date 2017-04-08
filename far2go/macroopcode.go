package far2go

const (
	MCODE_OP_EXIT = KEY_MACRO_OP_BASE // принудительно закончить выполнение макропоследовательности

	MCODE_OP_JMP  // Jumps..
	MCODE_OP_JZ
	MCODE_OP_JNZ
	MCODE_OP_JLT
	MCODE_OP_JLE
	MCODE_OP_JGT
	MCODE_OP_JGE

	MCODE_OP_NOP  // нет операции

	MCODE_OP_SAVE          // Присваивание переменной. Имя переменной следующие DWORD (как в $Text).
	MCODE_OP_SAVEREPCOUNT
	MCODE_OP_PUSHUNKNOWN   // неиницализированное значение (опускаемые параметры функций)
	MCODE_OP_PUSHINT       // Положить значение на стек. Само
	MCODE_OP_PUSHFLOAT     // Положить значение на стек. double
	MCODE_OP_PUSHSTR       // значение - следующий DWORD
	MCODE_OP_PUSHVAR       // или несколько таковых (как в $Text)
	MCODE_OP_PUSHCONST     // в стек положить константу

	MCODE_OP_REP  // $rep - признак начала цикла
	MCODE_OP_END  // $end - признак конца цикла/условия

	// Одноместные операции
	// ++a --a

	MCODE_OP_NEGATE  // -a
	MCODE_OP_NOT     // !a
	MCODE_OP_BITNOT  // ~a

	// Двуместные операции
	MCODE_OP_MUL  // a *  b
	MCODE_OP_DIV  // a /  b

	MCODE_OP_ADD  // a +  b
	MCODE_OP_SUB  // a -  b

	MCODE_OP_BITSHR  // a >> b
	MCODE_OP_BITSHL  // a << b

	MCODE_OP_LT  // a <  b
	MCODE_OP_LE  // a <= b
	MCODE_OP_GT  // a >  b
	MCODE_OP_GE  // a >= b

	MCODE_OP_EQ  // a == b
	MCODE_OP_NE  // a != b

	MCODE_OP_BITAND  // a &  b

	MCODE_OP_BITXOR  // a ^  b

	MCODE_OP_BITOR  // a |  b

	MCODE_OP_AND  // a && b

	MCODE_OP_XOR  // a ^^ b

	MCODE_OP_OR  // a || b

	MCODE_OP_ADDEQ     // a +=  b
	MCODE_OP_SUBEQ     // a -=  b
	MCODE_OP_MULEQ     // a *=  b
	MCODE_OP_DIVEQ     // a /=  b
	MCODE_OP_BITSHREQ  // a >>= b
	MCODE_OP_BITSHLEQ  // a <<= b
	MCODE_OP_BITANDEQ  // a &=  b
	MCODE_OP_BITXOREQ  // a ^=  b
	MCODE_OP_BITOREQ   // a |=  b

	// a++ a--

	MCODE_OP_DISCARD  // убрать значение с вершины стека
	MCODE_OP_DUP      // продублировать верхнее значение в стеке
	MCODE_OP_SWAP     // обменять местами два значения в вершине стека
	MCODE_OP_POP      // присвоить значение переменной и убрать из вершины стека
	MCODE_OP_COPY     // %a=%d стек не используется

	MCODE_OP_KEYS     // за этим кодом следуют ФАРовы коды клавиш
	MCODE_OP_ENDKEYS  // ФАРовы коды закончились.

	/* ************************************************************************* */
	MCODE_OP_IF     // Вообще-то эта группа в байткод
	MCODE_OP_ELSE   // не попадет никогда :)
	MCODE_OP_WHILE
	/* ************************************************************************* */
	MCODE_OP_CONTINUE  // $continue

	MCODE_OP_XLAT
	MCODE_OP_PLAINTEXT

	MCODE_OP_AKEY     // $AKey - клавиша которой вызвали макрос
	MCODE_OP_SELWORD  // $SelWord - выделить "слово"

	/* ************************************************************************* */
	// функции
	MCODE_F_NOFUNC = KEY_MACRO_F_BASE

	MCODE_F_ABS               // N=abs(N)
	MCODE_F_AKEY              // V=akey(Mode[Type])
	MCODE_F_ASC               // N=asc(S)
	MCODE_F_ATOI              // N=atoi(S[radix])
	MCODE_F_CLIP              // V=clip(N[V])
	MCODE_F_CHR               // S=chr(N)
	MCODE_F_DATE              // S=date([S])
	MCODE_F_DLG_GETVALUE      // V=Dlg.GetValue(IDN)
	MCODE_F_EDITOR_SEL        // V=Editor.Sel(Action[Opt])
	MCODE_F_EDITOR_SET        // N=Editor.Set(NVar)
	MCODE_F_EDITOR_UNDO       // V=Editor.Undo(N)
	MCODE_F_EDITOR_POS        // N=Editor.Pos(OpWhat[Where])
	MCODE_F_ENVIRON           // S=env(S)
	MCODE_F_FATTR             // N=fattr(S)
	MCODE_F_FEXIST            // S=fexist(S)
	MCODE_F_FSPLIT            // S=fsplit(SN)
	MCODE_F_IIF               // V=iif(CV1V2)
	MCODE_F_INDEX             // S=index(S1S2[Mode])
	MCODE_F_INT               // N=int(V)
	MCODE_F_ITOA              // S=itoa(N[radix])
	MCODE_F_KEY               // S=key(V)
	MCODE_F_LCASE             // S=lcase(S1)
	MCODE_F_LEN               // N=len(S)
	MCODE_F_MAX               // N=max(N1N2)
	MCODE_F_MENU_CHECKHOTKEY  // N=checkhotkey(S[N])
	MCODE_F_MENU_GETHOTKEY    // S=gethotkey([N])
	MCODE_F_MENU_SELECT       // N=Menu.Select(S[N[Dir]])
	MCODE_F_MIN               // N=min(N1N2)
	MCODE_F_MOD               // N=mod(ab) == a %  b
	MCODE_F_MLOAD             // B=mload(var)
	MCODE_F_MSAVE             // B=msave(var)
	MCODE_F_MSGBOX            // N=msgbox(["Title"["Text"[flags]]])
	MCODE_F_PANEL_FATTR       // N=Panel.FAttr(panelTypefileMask)
	MCODE_F_PANEL_SETPATH     // N=panel.SetPath(panelTypepathName[fileName])
	MCODE_F_PANEL_FEXIST      // N=Panel.FExist(panelTypefileMask)
	MCODE_F_PANEL_SETPOS      // N=Panel.SetPos(panelTypefileName)
	MCODE_F_PANEL_SETPOSIDX   // N=Panel.SetPosIdx(panelTypeIdx[InSelection])
	MCODE_F_PANEL_SELECT      // V=Panel.Select(panelTypeAction[Mode[Items]])
	MCODE_F_PANELITEM         // V=PanelItem(PanelIndexTypeInfo)
	MCODE_F_EVAL              // N=eval(S[N])
	MCODE_F_RINDEX            // S=rindex(S1S2[Mode])
	MCODE_F_SLEEP             // Sleep(N)
	MCODE_F_STRING            // S=string(V)
	MCODE_F_SUBSTR            // S=substr(Sstart[length])
	MCODE_F_UCASE             // S=ucase(S1)
	MCODE_F_WAITKEY           // V=waitkey([N[T]])
	MCODE_F_XLAT              // S=xlat(S)
	MCODE_F_FLOCK             // N=FLock(NN)
	MCODE_F_CALLPLUGIN        // V=callplugin(SysID[param])
	MCODE_F_REPLACE           // S=replace(sSsFsR[Count[Mode]])
	MCODE_F_PROMPT            // S=prompt("Title"["Prompt"[flags[ "Src"[ "History"]]]])
	MCODE_F_BM_ADD            // N=BM.Add()  - добавить текущие координаты и обрезать хвост
	MCODE_F_BM_CLEAR          // N=BM.Clear() - очистить все закладки
	MCODE_F_BM_DEL            // N=BM.Del([Idx]) - удаляет закладку с указанным индексом (x=1...) 0 - удаляет текущую закладку
	MCODE_F_BM_GET            // N=BM.Get(IdxM) - возвращает координаты строки (M==0) или колонки (M==1) закладки с индексом (Idx=1...)
	MCODE_F_BM_GOTO           // N=BM.Goto([n]) - переход на закладку с указанным индексом (0 --> текущую)
	MCODE_F_BM_NEXT           // N=BM.Next() - перейти на следующую закладку
	MCODE_F_BM_POP            // N=BM.Pop() - восстановить текущую позицию из закладки в конце стека и удалить закладку
	MCODE_F_BM_PREV           // N=BM.Prev() - перейти на предыдущую закладку
	MCODE_F_BM_BACK           // N=BM.Back() - перейти на предыдущую закладку с возможным сохранением текущей позиции
	MCODE_F_BM_PUSH           // N=BM.Push() - сохранить текущую позицию в виде закладки в конце стека
	MCODE_F_BM_STAT           // N=BM.Stat([M]) - возвращает информацию о закладках N=0 - текущее количество закладок	MCODE_F_TRIM                     // S=trim(S[N])
	MCODE_F_TRIM              // S=trim(S[N])
	MCODE_F_FLOAT             // N=float(V)
	MCODE_F_TESTFOLDER        // N=testfolder(S)
	MCODE_F_PRINT             // N=Print(Str)
	MCODE_F_MMODE             // N=MMode(Action[Value])
	MCODE_F_EDITOR_SETTITLE   // N=Editor.SetTitle([Title])
	MCODE_F_MENU_GETVALUE     // S=Menu.GetValue([N])
	MCODE_F_MENU_ITEMSTATUS   // N=Menu.ItemStatus([N])
	MCODE_F_BEEP              // N=beep([N])
	MCODE_F_KBDLAYOUT         // N=kbdLayout([N])
	MCODE_F_WINDOW_SCROLL     // N=Window.Scroll(Lines[Axis])

	/* ************************************************************************* */
	// булевые переменные - различные состояния
	MCODE_C_AREA_OTHER          = KEY_MACRO_C_BASE // Режим копирования текста с экрана вертикальные меню
	MCODE_C_AREA_SHELL                             // Файловые панели
	MCODE_C_AREA_VIEWER                            // Внутренняя программа просмотра
	MCODE_C_AREA_EDITOR                            // Редактор
	MCODE_C_AREA_DIALOG                            // Диалоги
	MCODE_C_AREA_SEARCH                            // Быстрый поиск в панелях
	MCODE_C_AREA_DISKS                             // Меню выбора дисков
	MCODE_C_AREA_MAINMENU                          // Основное меню
	MCODE_C_AREA_MENU                              // Прочие меню
	MCODE_C_AREA_HELP                              // Система помощи
	MCODE_C_AREA_INFOPANEL                         // Информационная панель
	MCODE_C_AREA_QVIEWPANEL                        // Панель быстрого просмотра
	MCODE_C_AREA_TREEPANEL                         // Панель дерева папок
	MCODE_C_AREA_FINDFOLDER                        // Поиск папок
	MCODE_C_AREA_USERMENU                          // Меню пользователя
	MCODE_C_AREA_AUTOCOMPLETION                    // Список автодополнения

	MCODE_C_FULLSCREENMODE  // полноэкранный режим?
	MCODE_C_ISUSERADMIN     // Administrator status
	MCODE_C_BOF             // начало файла/активного каталога?
	MCODE_C_EOF             // конец файла/активного каталога?
	MCODE_C_EMPTY           // ком.строка пуста?
	MCODE_C_SELECTED        // выделенный блок есть?
	MCODE_C_ROOTFOLDER      // аналог MCODE_C_APANEL_ROOT для активной панели

	MCODE_C_APANEL_BOF        // начало активного  каталога?
	MCODE_C_PPANEL_BOF        // начало пассивного каталога?
	MCODE_C_APANEL_EOF        // конец активного  каталога?
	MCODE_C_PPANEL_EOF        // конец пассивного каталога?
	MCODE_C_APANEL_ISEMPTY    // активная панель:  пуста?
	MCODE_C_PPANEL_ISEMPTY    // пассивная панель: пуста?
	MCODE_C_APANEL_SELECTED   // активная панель:  выделенные элементы есть?
	MCODE_C_PPANEL_SELECTED   // пассивная панель: выделенные элементы есть?
	MCODE_C_APANEL_ROOT       // это корневой каталог активной панели?
	MCODE_C_PPANEL_ROOT       // это корневой каталог пассивной панели?
	MCODE_C_APANEL_VISIBLE    // активная панель:  видима?
	MCODE_C_PPANEL_VISIBLE    // пассивная панель: видима?
	MCODE_C_APANEL_PLUGIN     // активная панель:  плагиновая?
	MCODE_C_PPANEL_PLUGIN     // пассивная панель: плагиновая?
	MCODE_C_APANEL_FILEPANEL  // активная панель:  файловая?
	MCODE_C_PPANEL_FILEPANEL  // пассивная панель: файловая?
	MCODE_C_APANEL_FOLDER     // активная панель:  текущий элемент каталог?
	MCODE_C_PPANEL_FOLDER     // пассивная панель: текущий элемент каталог?
	MCODE_C_APANEL_LEFT       // активная панель левая?
	MCODE_C_PPANEL_LEFT       // пассивная панель левая?
	MCODE_C_APANEL_LFN        // на активной панели длинные имена?
	MCODE_C_PPANEL_LFN        // на пассивной панели длинные имена?
	MCODE_C_APANEL_FILTER     // на активной панели включен фильтр?
	MCODE_C_PPANEL_FILTER     // на пассивной панели включен фильтр?

	MCODE_C_CMDLINE_BOF       // курсор в начале cmd-строки редактирования?
	MCODE_C_CMDLINE_EOF       // курсор в конце cmd-строки редактирования?
	MCODE_C_CMDLINE_EMPTY     // ком.строка пуста?
	MCODE_C_CMDLINE_SELECTED  // в ком.строке есть выделение блока?

	/* ************************************************************************* */
	// не булевые переменные
	MCODE_V_FAR_WIDTH  = KEY_MACRO_V_BASE // Far.Width - ширина консольного окна
	MCODE_V_FAR_HEIGHT                    // Far.Height - высота консольного окна
	MCODE_V_FAR_TITLE                     // Far.Title - текущий заголовок консольного окна
	MCODE_V_MACROAREA                     // MacroArea - имя текущей макрос области

	MCODE_V_APANEL_CURRENT      // APanel.Current - имя файла на активной панели
	MCODE_V_PPANEL_CURRENT      // PPanel.Current - имя файла на пассивной панели
	MCODE_V_APANEL_SELCOUNT     // APanel.SelCount - активная панель:  число выделенных элементов
	MCODE_V_PPANEL_SELCOUNT     // PPanel.SelCount - пассивная панель: число выделенных элементов
	MCODE_V_APANEL_PATH         // APanel.Path - активная панель:  путь на панели
	MCODE_V_PPANEL_PATH         // PPanel.Path - пассивная панель: путь на панели
	MCODE_V_APANEL_PATH0        // APanel.Path0 - активная панель:  путь на панели до вызова плагинов
	MCODE_V_PPANEL_PATH0        // PPanel.Path0 - пассивная панель: путь на панели до вызова плагинов
	MCODE_V_APANEL_UNCPATH      // APanel.UNCPath - активная панель:  UNC-путь на панели
	MCODE_V_PPANEL_UNCPATH      // PPanel.UNCPath - пассивная панель: UNC-путь на панели
	MCODE_V_APANEL_WIDTH        // APanel.Width - активная панель:  ширина панели
	MCODE_V_PPANEL_WIDTH        // PPanel.Width - пассивная панель: ширина панели
	MCODE_V_APANEL_TYPE         // APanel.Type - тип активной панели
	MCODE_V_PPANEL_TYPE         // PPanel.Type - тип пассивной панели
	MCODE_V_APANEL_ITEMCOUNT    // APanel.ItemCount - активная панель:  число элементов
	MCODE_V_PPANEL_ITEMCOUNT    // PPanel.ItemCount - пассивная панель: число элементов
	MCODE_V_APANEL_CURPOS       // APanel.CurPos - активная панель:  текущий индекс
	MCODE_V_PPANEL_CURPOS       // PPanel.CurPos - пассивная панель: текущий индекс
	MCODE_V_APANEL_OPIFLAGS     // APanel.OPIFlags - активная панель: флаги открытого плагина
	MCODE_V_PPANEL_OPIFLAGS     // PPanel.OPIFlags - пассивная панель: флаги открытого плагина
	MCODE_V_APANEL_DRIVETYPE    // APanel.DriveType - активная панель: тип привода
	MCODE_V_PPANEL_DRIVETYPE    // PPanel.DriveType - пассивная панель: тип привода
	MCODE_V_APANEL_HEIGHT       // APanel.Height - активная панель:  высота панели
	MCODE_V_PPANEL_HEIGHT       // PPanel.Height - пассивная панель: высота панели
	MCODE_V_APANEL_COLUMNCOUNT  // APanel.ColumnCount - активная панель:  количество колонок
	MCODE_V_PPANEL_COLUMNCOUNT  // PPanel.ColumnCount - пассивная панель: количество колонок
	MCODE_V_APANEL_HOSTFILE     // APanel.HostFile - активная панель:  имя Host-файла
	MCODE_V_PPANEL_HOSTFILE     // PPanel.HostFile - пассивная панель: имя Host-файла
	MCODE_V_APANEL_PREFIX       // APanel.Prefix
	MCODE_V_PPANEL_PREFIX       // PPanel.Prefix

	MCODE_V_ITEMCOUNT  // ItemCount - число элементов в текущем объекте
	MCODE_V_CURPOS     // CurPos - текущий индекс в текущем объекте
	MCODE_V_TITLE      // Title - заголовок текущего объекта
	MCODE_V_HEIGHT     // Height - высота текущего объекта
	MCODE_V_WIDTH      // Width - ширина текущего объекта

	MCODE_V_EDITORFILENAME  // Editor.FileName - имя редактируемого файла
	MCODE_V_EDITORLINES     // Editor.Lines - количество строк в редакторе
	MCODE_V_EDITORCURLINE   // Editor.CurLine - текущая линия в редакторе (в дополнении к Count)
	MCODE_V_EDITORCURPOS    // Editor.CurPos - текущая поз. в редакторе
	MCODE_V_EDITORREALPOS   // Editor.RealPos - текущая поз. в редакторе без привязки к размеру табуляции
	MCODE_V_EDITORSTATE     // Editor.State
	MCODE_V_EDITORVALUE     // Editor.Value - содержимое текущей строки
	MCODE_V_EDITORSELVALUE  // Editor.SelValue - содержит содержимое выделенного блока

	MCODE_V_DLGITEMTYPE   // Dlg.ItemType
	MCODE_V_DLGITEMCOUNT  // Dlg.ItemCount
	MCODE_V_DLGCURPOS     // Dlg.CurPos
	MCODE_V_DLGINFOID     // Dlg.Info.Id

	MCODE_V_VIEWERFILENAME  // Viewer.FileName - имя просматриваемого файла
	MCODE_V_VIEWERSTATE     // Viewer.State

	MCODE_V_CMDLINE_ITEMCOUNT  // CmdLine.ItemCount
	MCODE_V_CMDLINE_CURPOS     // CmdLine.CurPos
	MCODE_V_CMDLINE_VALUE      // CmdLine.Value

	MCODE_V_DRVSHOWPOS   // Drv.ShowPos - меню выбора дисков отображено: 1=слева (Alt-F1) 2=справа (Alt-F2) 0="нету его"
	MCODE_V_DRVSHOWMODE  // Drv.ShowMode - режимы отображения меню выбора дисков

	MCODE_V_HELPFILENAME  // Help.FileName
	MCODE_V_HELPTOPIC     // Help.Topic
	MCODE_V_HELPSELTOPIC  // Help.SelTopic

	MCODE_V_MENU_VALUE  // Menu.Value
)
