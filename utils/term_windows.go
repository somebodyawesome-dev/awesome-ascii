//go:build windows

package utils

func getTerminalSize() TermSize {
	ws := &TermSize{}

	return *ws
}
