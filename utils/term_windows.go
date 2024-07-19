//go:build windows

package utils

func getTerminalSize()  (TermSize,error) {
	ws := &TermSize{}

	return *ws,nil
}
