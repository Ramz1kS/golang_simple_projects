package consts

const (
	SYM_CROSS rune = 'X'
	SYM_ZERO rune = 'O'
	SYM_EMPTY = '.'
	SYM_CHECK_HORIZ = '-'
	SYM_CHECK_VERT = '|'
	SYM_CHECK_ANGL_UP = '/'
	SYM_CHECK_ANGL_DWN = '\\'
)

const (
	DIR_HORIZ int = iota 
	DIR_VERT 
	DIR_ANGL_UP
	DIR_ANGL_DWN
)

const (
	ACTION_PLAY int = iota + 1
	ACTION_RENAME
	ACTION_SWITCH_SYM
	ACTION_EXIT
)