package widget

import "container/list"


type  WDialog struct {
	widget        Widget

					/* Set by the user */
	compact       bool              /* Suppress spaces around the frame */
	help_ctx      string            /* Name of the help entry */
	color         int;              /* Color set. Unused in viewer and editor */
	title         *string;          /* Title of the dialog */

					/* Set and received by the user */
	ret_value     int;              /* Result of dlg_run() */

					/* Internal flags */
	winch_pending bool;             /* SIGWINCH signal has been got. Resize dialog after rise */
	mouse_status  int;              /* For the autorepeat status of the mouse */

					/* Internal variables */
	widgets       *list.List        /* widgets list */
	current       *list.List        /* Currently active widget */
	widget_id     uint64;           /* maximum id of all widgets */
					//void             *data;        /* Data can be passed to dialog */
	event_group   *string;          /* Name of event group for this dialog */

	get_shortcut  dlg_shortcut_str; /* Shortcut string */
					//dlg_title_str    get_title;    /* useless for modal dialogs */
};

type dlg_shortcut_str func(int) string


