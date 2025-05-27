package structs

type WinInfo struct {
	DidWin    bool
	DidEnd		bool
	Symbol    rune
	Direction int
	StartRow  int
	StartCol  int
}
