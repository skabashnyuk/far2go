package far2go

import "errors"

type PaletteType map[PaletteColor]uint

var DefaultPalette = PaletteType{
	COL_MENUTEXT:              F_WHITE | B_CYAN,
	COL_MENUSELECTEDTEXT:      F_WHITE | B_BLACK,
	COL_MENUHIGHLIGHT:         F_YELLOW | B_CYAN,
	COL_MENUSELECTEDHIGHLIGHT: F_YELLOW | B_BLACK,
	COL_MENUBOX:               F_WHITE | B_CYAN,
	COL_MENUTITLE:             F_WHITE | B_CYAN,

	COL_HMENUTEXT:              F_BLACK | B_CYAN,
	COL_HMENUSELECTEDTEXT:      F_WHITE | B_BLACK,
	COL_HMENUHIGHLIGHT:         F_YELLOW | B_CYAN,
	COL_HMENUSELECTEDHIGHLIGHT: F_YELLOW | B_BLACK,

	COL_PANELTEXT:           F_LIGHTCYAN | B_BLUE,
	COL_PANELSELECTEDTEXT:   F_YELLOW | B_BLUE,
	COL_PANELHIGHLIGHTTEXT:  F_LIGHTGRAY | B_BLUE,
	COL_PANELINFOTEXT:       F_YELLOW | B_BLUE,
	COL_PANELCURSOR:         F_BLACK | B_CYAN,
	COL_PANELSELECTEDCURSOR: F_YELLOW | B_CYAN,
	COL_PANELTITLE:          F_LIGHTCYAN | B_BLUE,
	COL_PANELSELECTEDTITLE:  F_BLACK | B_CYAN,
	COL_PANELCOLUMNTITLE:    F_YELLOW | B_BLUE,
	COL_PANELTOTALINFO:      F_LIGHTCYAN | B_BLUE,
	COL_PANELSELECTEDINFO:   F_YELLOW | B_CYAN,

	COL_DIALOGTEXT:                    F_BLACK | B_LIGHTGRAY,
	COL_DIALOGHIGHLIGHTTEXT:           F_YELLOW | B_LIGHTGRAY,
	COL_DIALOGBOX:                     F_BLACK | B_LIGHTGRAY,
	COL_DIALOGBOXTITLE:                F_BLACK | B_LIGHTGRAY,
	COL_DIALOGHIGHLIGHTBOXTITLE:       F_YELLOW | B_LIGHTGRAY,
	COL_DIALOGEDIT:                    F_BLACK | B_CYAN,
	COL_DIALOGBUTTON:                  F_BLACK | B_LIGHTGRAY,
	COL_DIALOGSELECTEDBUTTON:          F_BLACK | B_CYAN,
	COL_DIALOGHIGHLIGHTBUTTON:         F_YELLOW | B_LIGHTGRAY,
	COL_DIALOGHIGHLIGHTSELECTEDBUTTON: F_YELLOW | B_CYAN,

	COL_DIALOGLISTTEXT:              F_BLACK | B_LIGHTGRAY,
	COL_DIALOGLISTSELECTEDTEXT:      F_WHITE | B_BLACK,
	COL_DIALOGLISTHIGHLIGHT:         F_YELLOW | B_LIGHTGRAY,
	COL_DIALOGLISTSELECTEDHIGHLIGHT: F_YELLOW | B_BLACK,

	COL_WARNDIALOGTEXT:                    F_WHITE | B_RED,
	COL_WARNDIALOGHIGHLIGHTTEXT:           F_YELLOW | B_RED,
	COL_WARNDIALOGBOX:                     F_WHITE | B_RED,
	COL_WARNDIALOGBOXTITLE:                F_WHITE | B_RED,
	COL_WARNDIALOGHIGHLIGHTBOXTITLE:       F_YELLOW | B_RED,
	COL_WARNDIALOGEDIT:                    F_BLACK | B_CYAN,
	COL_WARNDIALOGBUTTON:                  F_WHITE | B_RED,
	COL_WARNDIALOGSELECTEDBUTTON:          F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGHIGHLIGHTBUTTON:         F_YELLOW | B_RED,
	COL_WARNDIALOGHIGHLIGHTSELECTEDBUTTON: F_YELLOW | B_LIGHTGRAY,

	COL_KEYBARNUM:        F_LIGHTGRAY | B_BLACK,
	COL_KEYBARTEXT:       F_BLACK | B_CYAN,
	COL_KEYBARBACKGROUND: F_LIGHTGRAY | B_BLACK,

	COL_COMMANDLINE: F_LIGHTGRAY | B_BLACK,

	COL_CLOCK: F_BLACK | B_CYAN,

	COL_VIEWERTEXT:         F_LIGHTCYAN | B_BLUE,
	COL_VIEWERSELECTEDTEXT: F_BLACK | B_CYAN,
	COL_VIEWERSTATUS:       F_BLACK | B_CYAN,

	COL_EDITORTEXT:         F_LIGHTCYAN | B_BLUE,
	COL_EDITORSELECTEDTEXT: F_BLACK | B_CYAN,
	COL_EDITORSTATUS:       F_BLACK | B_CYAN,

	COL_HELPTEXT:          F_BLACK | B_CYAN,
	COL_HELPHIGHLIGHTTEXT: F_WHITE | B_CYAN,
	COL_HELPTOPIC:         F_YELLOW | B_CYAN,
	COL_HELPSELECTEDTOPIC: F_WHITE | B_BLACK,
	COL_HELPBOX:           F_BLACK | B_CYAN,
	COL_HELPBOXTITLE:      F_BLACK | B_CYAN,

	COL_PANELDRAGTEXT:       F_YELLOW | B_CYAN,
	COL_DIALOGEDITUNCHANGED: F_LIGHTGRAY | B_CYAN,
	COL_PANELSCROLLBAR:      F_LIGHTCYAN | B_BLUE,
	COL_HELPSCROLLBAR:       F_BLACK | B_CYAN,
	COL_PANELBOX:            F_LIGHTCYAN | B_BLUE,
	COL_PANELSCREENSNUMBER:  F_LIGHTCYAN | B_BLACK,
	COL_DIALOGEDITSELECTED:  F_WHITE | B_BLACK,
	COL_COMMANDLINESELECTED: F_BLACK | B_CYAN,
	COL_VIEWERARROWS:        F_YELLOW | B_BLUE,

	COL_RESERVED0: 0,

	COL_DIALOGLISTSCROLLBAR: F_BLACK | B_LIGHTGRAY,
	COL_MENUSCROLLBAR:       F_WHITE | B_CYAN,
	COL_VIEWERSCROLLBAR:     F_LIGHTCYAN | B_BLUE,

	COL_COMMANDLINEPREFIX: F_LIGHTGRAY | B_BLACK,

	COL_DIALOGDISABLED:     F_DARKGRAY | B_LIGHTGRAY,
	COL_DIALOGEDITDISABLED: F_DARKGRAY | B_CYAN,
	COL_DIALOGLISTDISABLED: F_DARKGRAY | B_LIGHTGRAY,

	COL_WARNDIALOGDISABLED:     F_DARKGRAY | B_RED,
	COL_WARNDIALOGEDITDISABLED: F_DARKGRAY | B_CYAN,
	COL_WARNDIALOGLISTDISABLED: F_DARKGRAY | B_RED,

	COL_MENUDISABLEDTEXT: F_DARKGRAY | B_CYAN,

	COL_EDITORCLOCK: F_BLACK | B_CYAN,
	COL_VIEWERCLOCK: F_BLACK | B_CYAN,

	COL_DIALOGLISTTITLE: F_BLACK | B_LIGHTGRAY,
	COL_DIALOGLISTBOX:   F_BLACK | B_LIGHTGRAY,

	COL_WARNDIALOGEDITSELECTED:  F_WHITE | B_BLACK,
	COL_WARNDIALOGEDITUNCHANGED: F_LIGHTGRAY | B_CYAN,

	//COL_DIALOGCBOXTEXT:F_WHITE | B_CYAN,
	//COL_DIALOGCBOXSELECTEDTEXT:F_WHITE | B_BLACK,
	//COL_DIALOGCBOXHIGHLIGHT:F_YELLOW | B_CYAN,
	//COL_DIALOGCBOXSELECTEDHIGHLIGHT:F_YELLOW | B_BLACK,
	//COL_DIALOGCBOXBOX:F_WHITE | B_CYAN,
	//COL_DIALOGCBOXTITLE:F_WHITE | B_CYAN,
	//COL_DIALOGCBOXDISABLED:F_DARKGRAY | B_CYAN,
	//COL_DIALOGCBOXSCROLLBAR:F_WHITE | B_CYAN,

	COL_WARNDIALOGLISTTEXT:              F_WHITE | B_RED,
	COL_WARNDIALOGLISTSELECTEDTEXT:      F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGLISTHIGHLIGHT:         F_YELLOW | B_RED,
	COL_WARNDIALOGLISTSELECTEDHIGHLIGHT: F_YELLOW | B_LIGHTGRAY,
	COL_WARNDIALOGLISTBOX:               F_WHITE | B_RED,
	COL_WARNDIALOGLISTTITLE:             F_WHITE | B_RED,
	COL_WARNDIALOGLISTSCROLLBAR:         F_WHITE | B_RED,

	//COL_WARNDIALOGCBOXTEXT:F_WHITE | B_CYAN,
	//COL_WARNDIALOGCBOXSELECTEDTEXT:F_WHITE | B_BLACK,
	//COL_WARNDIALOGCBOXHIGHLIGHT:F_YELLOW | B_CYAN,
	//COL_WARNDIALOGCBOXSELECTEDHIGHLIGHT:F_YELLOW | B_BLACK,
	//COL_WARNDIALOGCBOXBOX:F_WHITE | B_CYAN,
	//COL_WARNDIALOGCBOXTITLE:F_WHITE | B_CYAN,
	//COL_WARNDIALOGCBOXDISABLED:F_DARKGRAY | B_CYAN,
	//COL_WARNDIALOGCBOXSCROLLBAR:F_WHITE | B_CYAN,

	COL_DIALOGLISTARROWS:              F_YELLOW | B_LIGHTGRAY,
	COL_DIALOGLISTARROWSDISABLED:      F_DARKGRAY | B_LIGHTGRAY,
	COL_DIALOGLISTARROWSSELECTED:      F_YELLOW | B_BLACK,
	COL_DIALOGCOMBOARROWS:             F_YELLOW | B_CYAN,
	COL_DIALOGCOMBOARROWSDISABLED:     F_DARKGRAY | B_CYAN,
	COL_DIALOGCOMBOARROWSSELECTED:     F_YELLOW | B_BLACK,
	COL_WARNDIALOGLISTARROWS:          F_YELLOW | B_RED,
	COL_WARNDIALOGLISTARROWSDISABLED:  F_DARKGRAY | B_RED,
	COL_WARNDIALOGLISTARROWSSELECTED:  F_YELLOW | B_LIGHTGRAY,
	COL_WARNDIALOGCOMBOARROWS:         F_YELLOW | B_CYAN,
	COL_WARNDIALOGCOMBOARROWSDISABLED: F_DARKGRAY | B_CYAN,
	COL_WARNDIALOGCOMBOARROWSSELECTED: F_YELLOW | B_BLACK,
	COL_MENUARROWS:                    F_YELLOW | B_CYAN,
	COL_MENUARROWSDISABLED:            F_DARKGRAY | B_CYAN,
	COL_MENUARROWSSELECTED:            F_YELLOW | B_BLACK,
	COL_COMMANDLINEUSERSCREEN:         F_LIGHTGRAY | B_BLACK,
	COL_EDITORSCROLLBAR:               F_LIGHTCYAN | B_BLUE,

	COL_MENUGRAYTEXT:                    F_DARKGRAY | B_CYAN,
	COL_MENUSELECTEDGRAYTEXT:            F_LIGHTGRAY | B_BLACK,
	COL_DIALOGCOMBOGRAY:                 F_DARKGRAY | B_CYAN,
	COL_DIALOGCOMBOSELECTEDGRAYTEXT:     F_LIGHTGRAY | B_BLACK,
	COL_DIALOGLISTGRAY:                  F_DARKGRAY | B_LIGHTGRAY,
	COL_DIALOGLISTSELECTEDGRAYTEXT:      F_LIGHTGRAY | B_BLACK,
	COL_WARNDIALOGCOMBOGRAY:             F_DARKGRAY | B_CYAN,
	COL_WARNDIALOGCOMBOSELECTEDGRAYTEXT: F_LIGHTGRAY | B_BLACK,
	COL_WARNDIALOGLISTGRAY:              F_DARKGRAY | B_RED,
	COL_WARNDIALOGLISTSELECTEDGRAYTEXT:  F_BLACK | B_LIGHTGRAY,

	COL_DIALOGDEFAULTBUTTON:                      F_BLACK | B_LIGHTGRAY,
	COL_DIALOGSELECTEDDEFAULTBUTTON:              F_BLACK | B_CYAN,
	COL_DIALOGHIGHLIGHTDEFAULTBUTTON:             F_YELLOW | B_LIGHTGRAY,
	COL_DIALOGHIGHLIGHTSELECTEDDEFAULTBUTTON:     F_YELLOW | B_CYAN,
	COL_WARNDIALOGDEFAULTBUTTON:                  F_WHITE | B_RED,
	COL_WARNDIALOGSELECTEDDEFAULTBUTTON:          F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGHIGHLIGHTDEFAULTBUTTON:         F_YELLOW | B_RED,
	COL_WARNDIALOGHIGHLIGHTSELECTEDDEFAULTBUTTON: F_YELLOW | B_LIGHTGRAY,
}

var BlackPalette = PaletteType{
	COL_MENUTEXT:              F_BLACK | B_LIGHTGRAY,
	COL_MENUSELECTEDTEXT:      F_LIGHTGRAY | B_BLACK,
	COL_MENUHIGHLIGHT:         F_WHITE | B_LIGHTGRAY,
	COL_MENUSELECTEDHIGHLIGHT: F_WHITE | B_BLACK,
	COL_MENUBOX:               F_BLACK | B_LIGHTGRAY,
	COL_MENUTITLE:             F_BLACK | B_LIGHTGRAY,

	COL_HMENUTEXT:              F_BLACK | B_LIGHTGRAY,
	COL_HMENUSELECTEDTEXT:      F_LIGHTGRAY | B_BLACK,
	COL_HMENUHIGHLIGHT:         F_WHITE | B_LIGHTGRAY,
	COL_HMENUSELECTEDHIGHLIGHT: F_WHITE | B_BLACK,

	COL_PANELTEXT:           F_LIGHTGRAY | B_BLACK,
	COL_PANELSELECTEDTEXT:   F_WHITE | B_BLACK,
	COL_PANELHIGHLIGHTTEXT:  F_LIGHTGRAY | B_BLACK,
	COL_PANELINFOTEXT:       F_WHITE | B_BLACK,
	COL_PANELCURSOR:         F_BLACK | B_LIGHTGRAY,
	COL_PANELSELECTEDCURSOR: F_WHITE | B_LIGHTGRAY,
	COL_PANELTITLE:          F_LIGHTGRAY | B_BLACK,
	COL_PANELSELECTEDTITLE:  F_BLACK | B_LIGHTGRAY,
	COL_PANELCOLUMNTITLE:    F_WHITE | B_BLACK,
	COL_PANELTOTALINFO:      F_LIGHTGRAY | B_BLACK,
	COL_PANELSELECTEDINFO:   F_BLACK | B_LIGHTGRAY,

	COL_DIALOGTEXT:                    F_BLACK | B_LIGHTGRAY,
	COL_DIALOGHIGHLIGHTTEXT:           F_WHITE | B_LIGHTGRAY,
	COL_DIALOGBOX:                     F_BLACK | B_LIGHTGRAY,
	COL_DIALOGBOXTITLE:                F_BLACK | B_LIGHTGRAY,
	COL_DIALOGHIGHLIGHTBOXTITLE:       F_WHITE | B_BLACK,
	COL_DIALOGEDIT:                    F_LIGHTGRAY | B_BLACK,
	COL_DIALOGBUTTON:                  F_BLACK | B_LIGHTGRAY,
	COL_DIALOGSELECTEDBUTTON:          F_LIGHTGRAY | B_BLACK,
	COL_DIALOGHIGHLIGHTBUTTON:         F_WHITE | B_LIGHTGRAY,
	COL_DIALOGHIGHLIGHTSELECTEDBUTTON: F_WHITE | B_BLACK,

	COL_DIALOGLISTTEXT:              F_BLACK | B_LIGHTGRAY,
	COL_DIALOGLISTSELECTEDTEXT:      F_WHITE | B_BLACK,
	COL_DIALOGLISTHIGHLIGHT:         F_WHITE | B_LIGHTGRAY,
	COL_DIALOGLISTSELECTEDHIGHLIGHT: F_WHITE | B_BLACK,

	COL_WARNDIALOGTEXT:                    F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGHIGHLIGHTTEXT:           F_WHITE | B_LIGHTGRAY,
	COL_WARNDIALOGBOX:                     F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGBOXTITLE:                F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGHIGHLIGHTBOXTITLE:       F_WHITE | B_BLACK,
	COL_WARNDIALOGEDIT:                    F_LIGHTGRAY | B_BLACK,
	COL_WARNDIALOGBUTTON:                  F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGSELECTEDBUTTON:          F_LIGHTGRAY | B_BLACK,
	COL_WARNDIALOGHIGHLIGHTBUTTON:         F_WHITE | B_LIGHTGRAY,
	COL_WARNDIALOGHIGHLIGHTSELECTEDBUTTON: F_WHITE | B_BLACK,

	COL_KEYBARNUM:        F_LIGHTGRAY | B_BLACK,
	COL_KEYBARTEXT:       F_BLACK | B_LIGHTGRAY,
	COL_KEYBARBACKGROUND: F_LIGHTGRAY | B_BLACK,

	COL_COMMANDLINE: F_LIGHTGRAY | B_BLACK,

	COL_CLOCK: F_BLACK | B_LIGHTGRAY,

	COL_VIEWERTEXT:         F_LIGHTGRAY | B_BLACK,
	COL_VIEWERSELECTEDTEXT: F_BLACK | B_LIGHTGRAY,
	COL_VIEWERSTATUS:       F_BLACK | B_LIGHTGRAY,

	COL_EDITORTEXT:         F_LIGHTGRAY | B_BLACK,
	COL_EDITORSELECTEDTEXT: F_BLACK | B_LIGHTGRAY,
	COL_EDITORSTATUS:       F_BLACK | B_LIGHTGRAY,

	COL_HELPTEXT:          F_LIGHTGRAY | B_BLACK,
	COL_HELPHIGHLIGHTTEXT: F_LIGHTGRAY | B_BLACK,
	COL_HELPTOPIC:         F_WHITE | B_BLACK,
	COL_HELPSELECTEDTOPIC: F_BLACK | B_LIGHTGRAY,
	COL_HELPBOX:           F_LIGHTGRAY | B_BLACK,
	COL_HELPBOXTITLE:      F_LIGHTGRAY | B_BLACK,

	COL_PANELDRAGTEXT:       F_BLACK | B_LIGHTGRAY,
	COL_DIALOGEDITUNCHANGED: F_WHITE | B_BLACK,
	COL_PANELSCROLLBAR:      F_LIGHTGRAY | B_BLACK,
	COL_HELPSCROLLBAR:       F_LIGHTGRAY | B_BLACK,
	COL_PANELBOX:            F_LIGHTGRAY | B_BLACK,
	COL_PANELSCREENSNUMBER:  F_WHITE | B_BLACK,
	COL_DIALOGEDITSELECTED:  F_BLACK | B_LIGHTGRAY,
	COL_COMMANDLINESELECTED: F_BLACK | B_LIGHTGRAY,
	COL_VIEWERARROWS:        F_WHITE | B_BLACK,

	COL_RESERVED0: 1,

	COL_DIALOGLISTSCROLLBAR: F_BLACK | B_LIGHTGRAY,
	COL_MENUSCROLLBAR:       F_BLACK | B_LIGHTGRAY,
	COL_VIEWERSCROLLBAR:     F_LIGHTGRAY | B_BLACK,

	COL_COMMANDLINEPREFIX: F_LIGHTGRAY | B_BLACK,

	COL_DIALOGDISABLED:     F_LIGHTGRAY | B_BLACK,
	COL_DIALOGEDITDISABLED: F_DARKGRAY | B_LIGHTGRAY,
	COL_DIALOGLISTDISABLED: F_DARKGRAY | B_LIGHTGRAY,

	COL_WARNDIALOGDISABLED:     F_DARKGRAY | B_LIGHTGRAY,
	COL_WARNDIALOGEDITDISABLED: F_DARKGRAY | B_LIGHTGRAY,
	COL_WARNDIALOGLISTDISABLED: F_DARKGRAY | B_LIGHTGRAY,

	COL_MENUDISABLEDTEXT: F_DARKGRAY | B_LIGHTGRAY,

	COL_EDITORCLOCK: F_BLACK | B_LIGHTGRAY,
	COL_VIEWERCLOCK: F_BLACK | B_LIGHTGRAY,

	COL_DIALOGLISTTITLE: F_BLACK | B_LIGHTGRAY,
	COL_DIALOGLISTBOX:   F_BLACK | B_LIGHTGRAY,

	COL_WARNDIALOGEDITSELECTED:  F_BLACK | B_WHITE,
	COL_WARNDIALOGEDITUNCHANGED: F_WHITE | B_BLACK,

	//COL_DIALOGCBOXTEXT:F_BLACK | B_LIGHTGRAY,
	//COL_DIALOGCBOXSELECTEDTEXT:F_LIGHTGRAY | B_BLACK,
	//COL_DIALOGCBOXHIGHLIGHT:F_WHITE | B_LIGHTGRAY,
	//COL_DIALOGCBOXSELECTEDHIGHLIGHT:F_WHITE | B_BLACK,
	//COL_DIALOGCBOXBOX:F_BLACK | B_LIGHTGRAY,
	//COL_DIALOGCBOXTITLE:F_BLACK | B_LIGHTGRAY,
	//COL_DIALOGCBOXDISABLED:F_DARKGRAY | B_LIGHTGRAY,
	//COL_DIALOGCBOXSCROLLBAR:F_BLACK | B_LIGHTGRAY,

	COL_WARNDIALOGLISTTEXT:              F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGLISTSELECTEDTEXT:      F_WHITE | B_BLACK,
	COL_WARNDIALOGLISTHIGHLIGHT:         F_WHITE | B_LIGHTGRAY,
	COL_WARNDIALOGLISTSELECTEDHIGHLIGHT: F_WHITE | B_BLACK,
	COL_WARNDIALOGLISTBOX:               F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGLISTTITLE:             F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGLISTSCROLLBAR:         F_BLACK | B_LIGHTGRAY,

	//COL_WARNDIALOGCBOXTEXT:F_BLACK | B_LIGHTGRAY,
	//COL_WARNDIALOGCBOXSELECTEDTEXT:F_WHITE | B_BLACK,
	//COL_WARNDIALOGCBOXHIGHLIGHT:F_WHITE | B_LIGHTGRAY,
	//COL_WARNDIALOGCBOXSELECTEDHIGHLIGHT:F_WHITE | B_BLACK,
	//COL_WARNDIALOGCBOXBOX:F_BLACK | B_LIGHTGRAY,
	//COL_WARNDIALOGCBOXTITLE:F_BLACK | B_LIGHTGRAY,
	//COL_WARNDIALOGCBOXDISABLED:F_DARKGRAY | B_LIGHTGRAY,
	//COL_WARNDIALOGCBOXSCROLLBAR:F_BLACK | B_LIGHTGRAY,

	COL_DIALOGLISTARROWS:              F_WHITE | B_LIGHTGRAY,
	COL_DIALOGLISTARROWSDISABLED:      F_DARKGRAY | B_LIGHTGRAY,
	COL_DIALOGLISTARROWSSELECTED:      F_WHITE | B_BLACK,
	COL_DIALOGCOMBOARROWS:             F_WHITE | B_LIGHTGRAY,
	COL_DIALOGCOMBOARROWSDISABLED:     F_DARKGRAY | B_LIGHTGRAY,
	COL_DIALOGCOMBOARROWSSELECTED:     F_WHITE | B_BLACK,
	COL_WARNDIALOGLISTARROWS:          F_WHITE | B_LIGHTGRAY,
	COL_WARNDIALOGLISTARROWSDISABLED:  F_DARKGRAY | B_LIGHTGRAY,
	COL_WARNDIALOGLISTARROWSSELECTED:  F_WHITE | B_BLACK,
	COL_WARNDIALOGCOMBOARROWS:         F_WHITE | B_LIGHTGRAY,
	COL_WARNDIALOGCOMBOARROWSDISABLED: F_DARKGRAY | B_LIGHTGRAY,
	COL_WARNDIALOGCOMBOARROWSSELECTED: F_WHITE | B_BLACK,
	COL_MENUARROWS:                    F_WHITE | B_LIGHTGRAY,
	COL_MENUARROWSDISABLED:            F_DARKGRAY | B_LIGHTGRAY,
	COL_MENUARROWSSELECTED:            F_WHITE | B_BLACK,
	COL_COMMANDLINEUSERSCREEN:         F_LIGHTGRAY | B_BLACK,
	COL_EDITORSCROLLBAR:               F_LIGHTGRAY | B_BLACK,

	COL_MENUGRAYTEXT:                    F_DARKGRAY | B_LIGHTGRAY,
	COL_MENUSELECTEDGRAYTEXT:            F_LIGHTGRAY | B_BLACK,
	COL_DIALOGCOMBOGRAY:                 F_DARKGRAY | B_LIGHTGRAY,
	COL_DIALOGCOMBOSELECTEDGRAYTEXT:     F_LIGHTGRAY | B_BLACK,
	COL_DIALOGLISTGRAY:                  F_DARKGRAY | B_LIGHTGRAY,
	COL_DIALOGLISTSELECTEDGRAYTEXT:      F_WHITE | B_BLACK,
	COL_WARNDIALOGCOMBOGRAY:             F_DARKGRAY | B_LIGHTGRAY,
	COL_WARNDIALOGCOMBOSELECTEDGRAYTEXT: F_WHITE | B_BLACK,
	COL_WARNDIALOGLISTGRAY:              F_DARKGRAY | B_LIGHTGRAY,
	COL_WARNDIALOGLISTSELECTEDGRAYTEXT:  F_WHITE | B_BLACK,

	COL_DIALOGDEFAULTBUTTON:                      F_BLACK | B_LIGHTGRAY,
	COL_DIALOGSELECTEDDEFAULTBUTTON:              F_LIGHTGRAY | B_BLACK,
	COL_DIALOGHIGHLIGHTDEFAULTBUTTON:             F_WHITE | B_LIGHTGRAY,
	COL_DIALOGHIGHLIGHTSELECTEDDEFAULTBUTTON:     F_WHITE | B_BLACK,
	COL_WARNDIALOGDEFAULTBUTTON:                  F_BLACK | B_LIGHTGRAY,
	COL_WARNDIALOGSELECTEDDEFAULTBUTTON:          F_LIGHTGRAY | B_BLACK,
	COL_WARNDIALOGHIGHLIGHTDEFAULTBUTTON:         F_WHITE | B_LIGHTGRAY,
	COL_WARNDIALOGHIGHLIGHTSELECTEDDEFAULTBUTTON: F_WHITE | B_BLACK,
}

func (Obj *PaletteType) FarColorToReal(FarColor PaletteColor) (uint, bool) {
	i, ok := (*Obj)[FarColor]
	return i, ok
}

func PaletteForName(PaletteName string) (PaletteType, error) {
	switch PaletteName {
	case "default":
		fallthrough
	case "":
		return DefaultPalette, nil
	case "black":
		return BlackPalette, nil
	default:
		return nil, errors.New("Unknown palette " + PaletteName)

	}
}