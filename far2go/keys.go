package far2go

type VirtualKey uint
const (
	/*
	* Virtual Keys, Standard Set
	*/
	VK_LBUTTON VirtualKey = 0x01
	VK_RBUTTON VirtualKey = 0x02
	VK_CANCEL  VirtualKey = 0x03
	VK_MBUTTON VirtualKey = 0x04 /* NOT contiguous with L & RBUTTON */

	VK_XBUTTON1 VirtualKey = 0x05 /* NOT contiguous with L & RBUTTON */
	VK_XBUTTON2 VirtualKey = 0x06 /* NOT contiguous with L & RBUTTON */

	/*
	 * 0x07 : unassigned
	 */

	VK_BACK VirtualKey = 0x08
	VK_TAB  VirtualKey = 0x09

	/*
	 * 0x0A - 0x0B : reserved
	 */

	VK_CLEAR  VirtualKey = 0x0C
	VK_RETURN VirtualKey = 0x0D

	VK_SHIFT   VirtualKey = 0x10
	VK_CONTROL VirtualKey = 0x11
	VK_MENU    VirtualKey = 0x12
	VK_PAUSE   VirtualKey = 0x13
	VK_CAPITAL VirtualKey = 0x14

	VK_KANA    VirtualKey = 0x15
	VK_HANGEUL VirtualKey = 0x15 /* old name - should be here for compatibility */
	VK_HANGUL  VirtualKey = 0x15
	VK_JUNJA   VirtualKey = 0x17
	VK_FINAL   VirtualKey = 0x18
	VK_HANJA   VirtualKey = 0x19
	VK_KANJI   VirtualKey = 0x19

	VK_ESCAPE VirtualKey = 0x1B

	VK_CONVERT    VirtualKey = 0x1C
	VK_NONCONVERT VirtualKey = 0x1D
	VK_ACCEPT     VirtualKey = 0x1E
	VK_MODECHANGE VirtualKey = 0x1F

	VK_SPACE    VirtualKey = 0x20
	VK_PRIOR    VirtualKey = 0x21
	VK_NEXT     VirtualKey = 0x22
	VK_END      VirtualKey = 0x23
	VK_HOME     VirtualKey = 0x24
	VK_LEFT     VirtualKey = 0x25
	VK_UP       VirtualKey = 0x26
	VK_RIGHT    VirtualKey = 0x27
	VK_DOWN     VirtualKey = 0x28
	VK_SELECT   VirtualKey = 0x29
	VK_PRINT    VirtualKey = 0x2A
	VK_EXECUTE  VirtualKey = 0x2B
	VK_SNAPSHOT VirtualKey = 0x2C
	VK_INSERT   VirtualKey = 0x2D
	VK_DELETE   VirtualKey = 0x2E
	VK_HELP     VirtualKey = 0x2F

	/*
	 * VK_0 - VK_9 are the same as ASCII '0' - '9' (0x30 - 0x39)
	 * 0x40 : unassigned
	 * VK_A - VK_Z are the same as ASCII 'A' - 'Z' (0x41 - 0x5A)
	 */

	VK_LWIN VirtualKey = 0x5B
	VK_RWIN VirtualKey = 0x5C
	VK_APPS VirtualKey = 0x5D

	/*
	 * 0x5E : reserved
	 */

	VK_SLEEP VirtualKey = 0x5F

	VK_NUMPAD0   VirtualKey = 0x60
	VK_NUMPAD1   VirtualKey = 0x61
	VK_NUMPAD2   VirtualKey = 0x62
	VK_NUMPAD3   VirtualKey = 0x63
	VK_NUMPAD4   VirtualKey = 0x64
	VK_NUMPAD5   VirtualKey = 0x65
	VK_NUMPAD6   VirtualKey = 0x66
	VK_NUMPAD7   VirtualKey = 0x67
	VK_NUMPAD8   VirtualKey = 0x68
	VK_NUMPAD9   VirtualKey = 0x69
	VK_MULTIPLY  VirtualKey = 0x6A
	VK_ADD       VirtualKey = 0x6B
	VK_SEPARATOR VirtualKey = 0x6C
	VK_SUBTRACT  VirtualKey = 0x6D
	VK_DECIMAL   VirtualKey = 0x6E
	VK_DIVIDE    VirtualKey = 0x6F
	VK_F1        VirtualKey = 0x70
	VK_F2        VirtualKey = 0x71
	VK_F3        VirtualKey = 0x72
	VK_F4        VirtualKey = 0x73
	VK_F5        VirtualKey = 0x74
	VK_F6        VirtualKey = 0x75
	VK_F7        VirtualKey = 0x76
	VK_F8        VirtualKey = 0x77
	VK_F9        VirtualKey = 0x78
	VK_F10       VirtualKey = 0x79
	VK_F11       VirtualKey = 0x7A
	VK_F12       VirtualKey = 0x7B
	VK_F13       VirtualKey = 0x7C
	VK_F14       VirtualKey = 0x7D
	VK_F15       VirtualKey = 0x7E
	VK_F16       VirtualKey = 0x7F
	VK_F17       VirtualKey = 0x80
	VK_F18       VirtualKey = 0x81
	VK_F19       VirtualKey = 0x82
	VK_F20       VirtualKey = 0x83
	VK_F21       VirtualKey = 0x84
	VK_F22       VirtualKey = 0x85
	VK_F23       VirtualKey = 0x86
	VK_F24       VirtualKey = 0x87

	/*
	 * 0x88 - 0x8F : unassigned
	 */

	VK_NUMLOCK VirtualKey = 0x90
	VK_SCROLL  VirtualKey = 0x91

	/*
	 * VK_L* & VK_R* - left and right Alt, Ctrl and Shift virtual keys.
	 * Used only as parameters to GetAsyncKeyState() and GetKeyState().
	 * No other API or message will distinguish left and right keys in this way.
	 */
	VK_LSHIFT   VirtualKey = 0xA0
	VK_RSHIFT   VirtualKey = 0xA1
	VK_LCONTROL VirtualKey = 0xA2
	VK_RCONTROL VirtualKey = 0xA3
	VK_LMENU    VirtualKey = 0xA4
	VK_RMENU    VirtualKey = 0xA5

	VK_BROWSER_BACK      VirtualKey = 0xA6
	VK_BROWSER_FORWARD   VirtualKey = 0xA7
	VK_BROWSER_REFRESH   VirtualKey = 0xA8
	VK_BROWSER_STOP      VirtualKey = 0xA9
	VK_BROWSER_SEARCH    VirtualKey = 0xAA
	VK_BROWSER_FAVORITES VirtualKey = 0xAB
	VK_BROWSER_HOME      VirtualKey = 0xAC

	VK_VOLUME_MUTE         VirtualKey = 0xAD
	VK_VOLUME_DOWN         VirtualKey = 0xAE
	VK_VOLUME_UP           VirtualKey = 0xAF
	VK_MEDIA_NEXT_TRACK    VirtualKey = 0xB0
	VK_MEDIA_PREV_TRACK    VirtualKey = 0xB1
	VK_MEDIA_STOP          VirtualKey = 0xB2
	VK_MEDIA_PLAY_PAUSE    VirtualKey = 0xB3
	VK_LAUNCH_MAIL         VirtualKey = 0xB4
	VK_LAUNCH_MEDIA_SELECT VirtualKey = 0xB5
	VK_LAUNCH_APP1         VirtualKey = 0xB6
	VK_LAUNCH_APP2         VirtualKey = 0xB7

	EXTENDED_KEY_BASE   VirtualKey = 0x00010000
	INTERNAL_KEY_BASE   VirtualKey = 0x00020000
	INTERNAL_KEY_BASE_2 VirtualKey = 0x00030000
	INTERNAL_MACRO_BASE VirtualKey = 0x00080000
)

type BaseDefKeyboard uint
const (
	KEY_CTRLMASK BaseDefKeyboard = 0xFFF00000

	KEY_M_OEM  BaseDefKeyboard = 0x00100000
	KEY_M_SPEC BaseDefKeyboard = 0x00200000

	KEY_CTRL  BaseDefKeyboard = 0x01000000
	KEY_ALT   BaseDefKeyboard = 0x02000000
	KEY_SHIFT BaseDefKeyboard = 0x04000000
	KEY_RCTRL BaseDefKeyboard = 0x10000000
	KEY_RALT  BaseDefKeyboard = 0x20000000

	KEY_ALTDIGIT BaseDefKeyboard = 0x40000000
	KEY_RSHIFT   BaseDefKeyboard = 0x80000000

	KEY_BRACKET     BaseDefKeyboard = '['
	KEY_BACKBRACKET BaseDefKeyboard = ']'
	KEY_COMMA       BaseDefKeyboard = ','
	KEY_QUOTE       BaseDefKeyboard = '"'
	KEY_DOT         BaseDefKeyboard = '.'
	KEY_SLASH       BaseDefKeyboard = '/'
	KEY_COLON       BaseDefKeyboard = ':'
	KEY_SEMICOLON   BaseDefKeyboard = ';'
	KEY_BACKSLASH   BaseDefKeyboard = '\\'

	KEY_BS    BaseDefKeyboard = 0x00000008
	KEY_TAB   BaseDefKeyboard = 0x00000009
	KEY_ENTER BaseDefKeyboard = 0x0000000D
	KEY_ESC   BaseDefKeyboard = 0x0000001B
	KEY_SPACE BaseDefKeyboard = 0x00000020

	KEY_MASKF BaseDefKeyboard = 0x0001FFFF

	KEY_FKEY_BEGIN BaseDefKeyboard = EXTENDED_KEY_BASE

	KEY_BREAK BaseDefKeyboard = EXTENDED_KEY_BASE + VK_CANCEL

	KEY_PAUSE    BaseDefKeyboard = EXTENDED_KEY_BASE + VK_PAUSE
	KEY_CAPSLOCK BaseDefKeyboard = EXTENDED_KEY_BASE + VK_CAPITAL

	KEY_PGUP     BaseDefKeyboard = EXTENDED_KEY_BASE + VK_PRIOR
	KEY_PGDN     BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NEXT
	KEY_END      BaseDefKeyboard = EXTENDED_KEY_BASE + VK_END
	KEY_HOME     BaseDefKeyboard = EXTENDED_KEY_BASE + VK_HOME
	KEY_LEFT     BaseDefKeyboard = EXTENDED_KEY_BASE + VK_LEFT
	KEY_UP       BaseDefKeyboard = EXTENDED_KEY_BASE + VK_UP
	KEY_RIGHT    BaseDefKeyboard = EXTENDED_KEY_BASE + VK_RIGHT
	KEY_DOWN     BaseDefKeyboard = EXTENDED_KEY_BASE + VK_DOWN
	KEY_PRNTSCRN BaseDefKeyboard = EXTENDED_KEY_BASE + VK_SNAPSHOT
	KEY_INS      BaseDefKeyboard = EXTENDED_KEY_BASE + VK_INSERT
	KEY_DEL      BaseDefKeyboard = EXTENDED_KEY_BASE + VK_DELETE

	KEY_LWIN    BaseDefKeyboard = EXTENDED_KEY_BASE + VK_LWIN
	KEY_RWIN    BaseDefKeyboard = EXTENDED_KEY_BASE + VK_RWIN
	KEY_APPS    BaseDefKeyboard = EXTENDED_KEY_BASE + VK_APPS
	KEY_STANDBY BaseDefKeyboard = EXTENDED_KEY_BASE + VK_SLEEP

	KEY_NUMPAD0 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMPAD0
	KEY_NUMPAD1 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMPAD1
	KEY_NUMPAD2 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMPAD2
	KEY_NUMPAD3 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMPAD3
	KEY_NUMPAD4 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMPAD4
	KEY_NUMPAD5 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMPAD5
	KEY_CLEAR   BaseDefKeyboard = KEY_NUMPAD5
	KEY_NUMPAD6 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMPAD6
	KEY_NUMPAD7 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMPAD7
	KEY_NUMPAD8 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMPAD8
	KEY_NUMPAD9 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMPAD9

	KEY_MULTIPLY BaseDefKeyboard = EXTENDED_KEY_BASE + VK_MULTIPLY
	KEY_ADD      BaseDefKeyboard = EXTENDED_KEY_BASE + VK_ADD
	KEY_SUBTRACT BaseDefKeyboard = EXTENDED_KEY_BASE + VK_SUBTRACT
	KEY_DECIMAL  BaseDefKeyboard = EXTENDED_KEY_BASE + VK_DECIMAL
	KEY_DIVIDE   BaseDefKeyboard = EXTENDED_KEY_BASE + VK_DIVIDE

	KEY_F1  BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F1
	KEY_F2  BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F2
	KEY_F3  BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F3
	KEY_F4  BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F4
	KEY_F5  BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F5
	KEY_F6  BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F6
	KEY_F7  BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F7
	KEY_F8  BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F8
	KEY_F9  BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F9
	KEY_F10 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F10
	KEY_F11 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F11
	KEY_F12 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F12

	KEY_F13 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F13
	KEY_F14 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F14
	KEY_F15 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F15
	KEY_F16 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F16
	KEY_F17 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F17
	KEY_F18 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F18
	KEY_F19 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F19
	KEY_F20 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F20
	KEY_F21 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F21
	KEY_F22 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F22
	KEY_F23 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F23
	KEY_F24 BaseDefKeyboard = EXTENDED_KEY_BASE + VK_F24

	KEY_NUMLOCK    BaseDefKeyboard = EXTENDED_KEY_BASE + VK_NUMLOCK
	KEY_SCROLLLOCK BaseDefKeyboard = EXTENDED_KEY_BASE + VK_SCROLL

	KEY_BROWSER_BACK        BaseDefKeyboard = EXTENDED_KEY_BASE + VK_BROWSER_BACK
	KEY_BROWSER_FORWARD     BaseDefKeyboard = EXTENDED_KEY_BASE + VK_BROWSER_FORWARD
	KEY_BROWSER_REFRESH     BaseDefKeyboard = EXTENDED_KEY_BASE + VK_BROWSER_REFRESH
	KEY_BROWSER_STOP        BaseDefKeyboard = EXTENDED_KEY_BASE + VK_BROWSER_STOP
	KEY_BROWSER_SEARCH      BaseDefKeyboard = EXTENDED_KEY_BASE + VK_BROWSER_SEARCH
	KEY_BROWSER_FAVORITES   BaseDefKeyboard = EXTENDED_KEY_BASE + VK_BROWSER_FAVORITES
	KEY_BROWSER_HOME        BaseDefKeyboard = EXTENDED_KEY_BASE + VK_BROWSER_HOME
	KEY_VOLUME_MUTE         BaseDefKeyboard = EXTENDED_KEY_BASE + VK_VOLUME_MUTE
	KEY_VOLUME_DOWN         BaseDefKeyboard = EXTENDED_KEY_BASE + VK_VOLUME_DOWN
	KEY_VOLUME_UP           BaseDefKeyboard = EXTENDED_KEY_BASE + VK_VOLUME_UP
	KEY_MEDIA_NEXT_TRACK    BaseDefKeyboard = EXTENDED_KEY_BASE + VK_MEDIA_NEXT_TRACK
	KEY_MEDIA_PREV_TRACK    BaseDefKeyboard = EXTENDED_KEY_BASE + VK_MEDIA_PREV_TRACK
	KEY_MEDIA_STOP          BaseDefKeyboard = EXTENDED_KEY_BASE + VK_MEDIA_STOP
	KEY_MEDIA_PLAY_PAUSE    BaseDefKeyboard = EXTENDED_KEY_BASE + VK_MEDIA_PLAY_PAUSE
	KEY_LAUNCH_MAIL         BaseDefKeyboard = EXTENDED_KEY_BASE + VK_LAUNCH_MAIL
	KEY_LAUNCH_MEDIA_SELECT BaseDefKeyboard = EXTENDED_KEY_BASE + VK_LAUNCH_MEDIA_SELECT
	KEY_LAUNCH_APP1         BaseDefKeyboard = EXTENDED_KEY_BASE + VK_LAUNCH_APP1
	KEY_LAUNCH_APP2         BaseDefKeyboard = EXTENDED_KEY_BASE + VK_LAUNCH_APP2

	KEY_CTRLALTSHIFTPRESS   BaseDefKeyboard = INTERNAL_KEY_BASE + 1
	KEY_CTRLALTSHIFTRELEASE BaseDefKeyboard = INTERNAL_KEY_BASE + 2

	KEY_MSWHEEL_UP   BaseDefKeyboard = INTERNAL_KEY_BASE + 3
	KEY_MSWHEEL_DOWN BaseDefKeyboard = INTERNAL_KEY_BASE + 4

	KEY_RCTRLALTSHIFTPRESS   BaseDefKeyboard = INTERNAL_KEY_BASE + 7
	KEY_RCTRLALTSHIFTRELEASE BaseDefKeyboard = INTERNAL_KEY_BASE + 8
	KEY_NUMDEL               BaseDefKeyboard = INTERNAL_KEY_BASE + 9
	KEY_NUMENTER             BaseDefKeyboard = INTERNAL_KEY_BASE + 0xB

	KEY_MSWHEEL_LEFT  BaseDefKeyboard = INTERNAL_KEY_BASE + 0xC
	KEY_MSWHEEL_RIGHT BaseDefKeyboard = INTERNAL_KEY_BASE + 0xD

	KEY_MSLCLICK BaseDefKeyboard = INTERNAL_KEY_BASE + 0xF
	KEY_MSRCLICK BaseDefKeyboard = INTERNAL_KEY_BASE + 0x10

	KEY_MSM1CLICK BaseDefKeyboard = INTERNAL_KEY_BASE + 0x11
	KEY_MSM2CLICK BaseDefKeyboard = INTERNAL_KEY_BASE + 0x12
	KEY_MSM3CLICK BaseDefKeyboard = INTERNAL_KEY_BASE + 0x13

	KEY_MSLDBLCLICK BaseDefKeyboard = INTERNAL_KEY_BASE + 0x14
	KEY_MSRDBLCLICK BaseDefKeyboard = INTERNAL_KEY_BASE + 0x15

	KEY_MSM1DBLCLICK  BaseDefKeyboard = INTERNAL_KEY_BASE + 0x16
	KEY_MSM2DBLCLICK  BaseDefKeyboard = INTERNAL_KEY_BASE + 0x17
	KEY_MSM3DBLCLICK  BaseDefKeyboard = INTERNAL_KEY_BASE + 0x18
	KEY_VK_0xFF_BEGIN BaseDefKeyboard = EXTENDED_KEY_BASE + 0x00000100

	KEY_VK_0xFF_END BaseDefKeyboard = EXTENDED_KEY_BASE + 0x000001FF

	KEY_END_FKEY BaseDefKeyboard = 0x0001FFFF

	KEY_NONE     BaseDefKeyboard = INTERNAL_KEY_BASE_2 + 1
	KEY_IDLE     BaseDefKeyboard = INTERNAL_KEY_BASE_2 + 2
	KEY_DRAGCOPY BaseDefKeyboard = INTERNAL_KEY_BASE_2 + 3
	KEY_DRAGMOVE BaseDefKeyboard = INTERNAL_KEY_BASE_2 + 4

	KEY_KILLFOCUS                            BaseDefKeyboard = INTERNAL_KEY_BASE_2 + 6
	KEY_GOTFOCUS                             BaseDefKeyboard = INTERNAL_KEY_BASE_2 + 7
	KEY_CONSOLE_BUFFER_RESIZEBaseDefKeyboard                 = INTERNAL_KEY_BASE_2 + 8

	KEY_OP_BASE      BaseDefKeyboard = INTERNAL_KEY_BASE_2 + 0x100
	KEY_OP_XLAT      BaseDefKeyboard = KEY_OP_BASE + 0
	KEY_OP_PLAINTEXT BaseDefKeyboard = KEY_OP_BASE + 2
	KEY_OP_SELWORD   BaseDefKeyboard = KEY_OP_BASE + 3
	KEY_OP_ENDBASE   BaseDefKeyboard = INTERNAL_KEY_BASE_2 + 0x1FF

	KEY_END_SKEY  BaseDefKeyboard = 0x0003FFFF
	KEY_LAST_BASE BaseDefKeyboard = KEY_END_SKEY

	KEY_MACRO_BASE    BaseDefKeyboard = INTERNAL_MACRO_BASE
	KEY_MACRO_OP_BASE BaseDefKeyboard = INTERNAL_MACRO_BASE + 0x0000
	KEY_MACRO_C_BASE  BaseDefKeyboard = INTERNAL_MACRO_BASE + 0x0400
	KEY_MACRO_V_BASE  BaseDefKeyboard = INTERNAL_MACRO_BASE + 0x0800
	KEY_MACRO_F_BASE  BaseDefKeyboard = INTERNAL_MACRO_BASE + 0x0C00

	KEY_MACRO_U_BASE BaseDefKeyboard = INTERNAL_MACRO_BASE + 0x8000

	KEY_MACRO_ENDBASE BaseDefKeyboard = 0x000FFFFF
)
