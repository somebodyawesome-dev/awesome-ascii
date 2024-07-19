//go:build linux || darwin || dragonfly || freebsd || nacl || netbsd || openbsd || solaris

package utils

import (
	"syscall"
	"unsafe"
)

func getTerminalSize() (TermSize,error) {
	ws := &TermSize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		return *ws,errno
	}
	return *ws,nil
}

