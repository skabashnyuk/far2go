package widget

type  mouse_msg_t uint
/* Mouse messages */
const
(
	/*
	 * Notes:
	 * (1) "anywhere" means "inside or outside the widget".
	 * (2) the mouse wheel is not considered "mouse button".
	 */
	MSG_MOUSE_NONE mouse_msg_t = iota
	MSG_MOUSE_DOWN             /* When mouse button is pressed down inside the widget. */
	MSG_MOUSE_UP               /* When mouse button, previously pressed inside the widget, is released anywhere. */
	MSG_MOUSE_CLICK            /* When mouse button, previously pressed inside the widget, is released inside the widget. */
	MSG_MOUSE_DRAG             /* When a drag, initiated by button press inside the widget, occurs anywhere. */
	MSG_MOUSE_MOVE             /* (Not currently implemented in MC.) */
	MSG_MOUSE_SCROLL_UP        /* When mouse wheel is rotated away from the user. */
	MSG_MOUSE_SCROLL_DOWN      /* When mouse wheel is rotated towards the user. */
)

/*** structures declarations (and typedefs of structures)*****************************************/

type result struct {
	abort  bool;
	repeat bool;
}

/* Mouse event structure. */
type  mouse_event_t struct {
	msg     mouse_msg_t;
	x, y    int; /* Local to the widget. */
	buttons int; /* Bitwise-or of: GPM_B_LEFT, GPM_B_MIDDLE, GPM_B_RIGHT */
	count   int; /* One of: GPM_SINGLE, GPM_DOUBLE, GPM_TRIPLE */
	result;      /* A mechanism for the callback to report back: */
}




/*** typedefs(not structures) and defined constants **********************************************/

/*** global variables defined in .c file *********************************************************/

/*** declarations of public functions ************************************************************/

/* Translate GPM event to high-level event */
//mouse_event_t mouse_translate_event (Widget * w, Gpm_Event * event);
/* Process high-level mouse event */
//int mouse_process_event (Widget * w, mouse_event_t * event);
