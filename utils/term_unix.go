//go:build linux || darwin || dragonfly || freebsd || nacl || netbsd || openbsd || solaris

package utils

import (
	"fmt"
	"syscall"
	"unsafe"
)

func getTerminalSize() TermSize {
	ws := &TermSize{Col: 50, Row: 25}
	retCode, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		fmt.Println("Couldn't get tty size")
		fmt.Println("Falling back to use default value for tty")
		return *ws
	}
	return *ws
}
