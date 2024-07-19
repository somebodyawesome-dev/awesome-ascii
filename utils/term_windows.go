//go:build windows

package utils

import (
	"os"

	"golang.org/x/sys/windows"
)

func getTerminalSize() (TermSize, error) {
	var csbi windows.ConsoleScreenBufferInfo
	handle := windows.Handle(os.Stdout.Fd())
	err := windows.GetConsoleScreenBufferInfo(handle, &csbi)
	if err != nil {
		return TermSize{}, err
	}

	return TermSize{
		Row: uint16(csbi.Window.Bottom) - uint16(csbi.Window.Top) + 1,
		Col: uint16(csbi.Window.Right) - uint16(csbi.Window.Left) + 1,
	}, nil
}
