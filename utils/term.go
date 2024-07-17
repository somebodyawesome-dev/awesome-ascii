package utils

type TermSize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func GetTerminalSize() (TermSize, error) {
	return getTerminalSize()
}
