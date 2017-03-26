package widget

import "strings"

/* Widget messages */
type widget_msg_t uint

const
(
	MSG_INIT           widget_msg_t = iota /* Initialize widget */
	MSG_FOCUS                              /* Draw widget in focused state or widget has got focus */
	MSG_UNFOCUS                            /* Draw widget in unfocused state or widget has been unfocused */
	MSG_CHANGED_FOCUS                      /* Notification to owner about focus state change */
	MSG_ENABLE                             /* Change state to enabled */
	MSG_DISABLE                            /* Change state to disabled */
	MSG_DRAW                               /* Draw widget on screen */
	MSG_KEY                                /* Sent to widgets on key press */
	MSG_HOTKEY                             /* Sent to widget to catch preprocess key */
	MSG_HOTKEY_HANDLED                     /* A widget has got the hotkey */
	MSG_UNHANDLED_KEY                      /* Key that no widget handled */
	MSG_POST_KEY                           /* The key has been handled */
	MSG_ACTION                             /* Send to widget to handle command */
	MSG_NOTIFY                             /* Typically sent to dialog to inform it of state-change
                                    * of listboxes, check- and radiobuttons. */
	MSG_CURSOR    /* Sent to widget to position the cursor */
	MSG_IDLE      /* The idle state is active */
	MSG_RESIZE    /* Screen size has changed */
	MSG_VALIDATE  /* Dialog is to be closed */
	MSG_END       /* Shut down dialog */
	MSG_DESTROY   /* Sent to widget at destruction time */
)

type cb_ret_t uint

const
(
	MSG_NOT_HANDLED cb_ret_t = iota
	MSG_HANDLED
)

/* Widget options */
type widget_options_t uint

const
(
	WOP_DEFAULT     widget_options_t = 0 << 0
	WOP_WANT_HOTKEY                  = 1 << 0
	WOP_WANT_CURSOR                  = 1 << 1
	WOP_WANT_TAB                     = 1 << 2 /* Should the tab key be sent to the dialog? */
	WOP_IS_INPUT                     = 1 << 3
	WOP_SELECTABLE                   = 1 << 4
	WOP_TOP_SELECT                   = 1 << 5
)

/* Widget state */
type widget_state_t uint

const (
	WST_DEFAULT  widget_state_t = 0 << 0
	WST_DISABLED                = 1 << 0 /* Widget cannot be selected */
	WST_IDLE                    = 1 << 1
	WST_MODAL                   = 1 << 2 /* Widget (dialog) is modal */
	WST_FOCUSED                 = 1 << 3

	WST_CONSTRUCT = 1 << 15 /* Dialog has been constructed but not run yet */
	WST_ACTIVE    = 1 << 16 /* Dialog is visible and active */
	WST_SUSPENDED = 1 << 17 /* Dialog is suspended */
	WST_CLOSED    = 1 << 18 /* Dialog is closed */
)

/* Flags for widget repositioning on dialog resize */
type widget_pos_flags_t uint

const
(
	WPOS_FULLSCREEN   widget_pos_flags_t = 1 << 0                              /* widget occupies the whole screen */
	WPOS_CENTER_HORZ                     = 1 << 1                              /* center widget in horizontal */
	WPOS_CENTER_VERT                     = 1 << 2                              /* center widget in vertical */
	WPOS_CENTER                          = WPOS_CENTER_HORZ | WPOS_CENTER_VERT /* center widget */
	WPOS_TRYUP                           = 1 << 3                              /* try to move two lines up the widget */
	WPOS_KEEP_LEFT                       = 1 << 4                              /* keep widget distance to left border of dialog */
	WPOS_KEEP_RIGHT                      = 1 << 5                              /* keep widget distance to right border of dialog */
	WPOS_KEEP_TOP                        = 1 << 6                              /* keep widget distance to top border of dialog */
	WPOS_KEEP_BOTTOM                     = 1 << 7                              /* keep widget distance to bottom border of dialog */
	WPOS_KEEP_HORZ                       = WPOS_KEEP_LEFT | WPOS_KEEP_RIGHT
	WPOS_KEEP_VERT                       = WPOS_KEEP_TOP | WPOS_KEEP_BOTTOM
	WPOS_KEEP_ALL                        = WPOS_KEEP_HORZ | WPOS_KEEP_VERT
	WPOS_KEEP_DEFAULT                    = WPOS_KEEP_LEFT | WPOS_KEEP_TOP
)

/* NOTES:
 * If WPOS_FULLSCREEN is set then all other position flags are ignored.
 * If WPOS_CENTER_HORZ flag is used, other horizontal flags (WPOS_KEEP_LEFT, WPOS_KEEP_RIGHT,
 * and WPOS_KEEP_HORZ) are ignored.
 * If WPOS_CENTER_VERT flag is used, other horizontal flags (WPOS_KEEP_TOP, WPOS_KEEP_BOTTOM,
 * and WPOS_KEEP_VERT) are ignored.
 */

/*** structures declarations (and typedefs of structures)*****************************************/

/* Widget callback */

//typedef cb_ret_t (*widget_cb_fn) (Widget * widget, Widget * sender, widget_msg_t msg, int parm, void *data);
/* Widget mouse callback */
//typedef void (*widget_mouse_cb_fn) (Widget * w, mouse_msg_t msg, mouse_event_t * event);

/* Mouse-related fields. */
type mouse struct {
	/* Public members: */
	forced_capture bool /* Overrides the 'capture' member. Set explicitly by the programmer. */

	/* Implementation details: */
	capture           bool        /* Whether the widget "owns" the mouse. */
	last_msg          mouse_msg_t /* The previous event type processed. */
	last_buttons_down int
}

/* Every Widget must have this as its first element */
type Widget struct {
	x, y        int
	cols, lines int
	pos_flags   widget_pos_flags_t /* repositioning flags */
	options     widget_options_t   //   options;
	state       widget_state_t     //     ;
	id          uint               /* Number of the widget, starting with 0 */
	//callback    widget_cb_fn       //       ;
	//widget_mouse_cb_fn mouse_callback;
	owner *WDialog
}

/* structure for label (caption) with hotkey, if original text does not contain
 * hotkey, only start is valid and is equal to original text
 * hotkey is defined as char*, but mc support only singlebyte hotkey
 */
type hotkey_t struct {
	start  string
	hotkey string
	end    string
}

/*** global variables defined in .c file *********************************************************/

/*** declarations of public functions ************************************************************/

/* create hotkey from text */
func parse_hotkey(text string) hotkey_t {
	var result hotkey_t

	/* search for '&', that is not on the of text */
	cp := strings.Index(text, "&")
	if cp > 0 { //&& text[cp] != '0'
		result.start = text[:cp]
		result.hotkey = text[cp+1:cp+2]
		result.end = text[cp+2:]
	} else {
		result.start = text
		result.hotkey = ""
		result.end = ""
	}

	return result
}

/* --------------------------------------------------------------------------------------------- */
/**
  * Check whether one or several option flags are set or not.
  * @param w widget
  * @param options widget option flags
  *
  * @return TRUE if all requested option flags are set, FALSE otherwise.
  */
func widget_get_options(w *Widget, options widget_options_t) bool {
	return w.options&options == options
}

/* --------------------------------------------------------------------------------------------- */

/**
  * Check whether one or several state flags are set or not.
  * @param w widget
  * @param state widget state flags
  *
  * @return TRUE if all requested state flags are set, FALSE otherwise.
  */
func widget_get_state(w *Widget, state widget_state_t) bool {
	return w.state&state == state
}
