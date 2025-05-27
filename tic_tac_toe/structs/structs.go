package structs

type Field [3][3]rune

type WinInfo struct {
	DidWin    bool
	DidEnd		bool
	Symbol    rune
	Direction int
	StartRow  int
	StartCol  int
}
