package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

type winsize struct {
    Row    uint16
    Col    uint16
    Xpixel uint16
    Ypixel uint16
}

func GetWidth() uint {
    ws := &winsize{}
    retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(ws)))

    if int(retCode) == -1 {
        panic(errno)
    }
    return uint(ws.Col)
}
func GetWidth2() {
  cmd := exec.Command("stty", "size")
  cmd.Stdin = os.Stdin
  out, err := cmd.Output()
  fmt.Printf("out: %#v\n", string(out))
  fmt.Printf("err: %#v\n", err)
  if err != nil {
    log.Fatal(err)
  }
}
