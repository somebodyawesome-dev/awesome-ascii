//go:build windows

package utils

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

func getTerminalSize() TermSize {
	var csbi windows.ConsoleScreenBufferInfo
	handle := windows.Handle(os.Stdout.Fd())
	err := windows.GetConsoleScreenBufferInfo(handle, &csbi)
	if err != nil {
		fmt.Println("Couldn't get tty size")
		fmt.Println("Falling back to use default value for tty")
		return TermSize{Col: 50, Row: 25}
	}

	return TermSize{
		Row: uint16(csbi.Window.Bottom) - uint16(csbi.Window.Top) + 1,
		Col: uint16(csbi.Window.Right) - uint16(csbi.Window.Left) + 1,
	}
}
