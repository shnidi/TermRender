package screen

import (
	"os"

	"golang.org/x/sys/unix"
)

type Screen struct {
}

func NewScreen() (screen *Screen, err error) {
	//get console layout
	ws, err := getWinsize()
	if err != nil {
		return &Screen{}, err
	}
	println(ws.Col)
	println(ws.Row)
	for i := 0; i < int(ws.Row); i++ {
		for j := 0; j < int(ws.Col); j++ {
			print("*")
		}
		if i < int(ws.Row)-1 {
			println("")
		}
	}
	return &Screen{}, nil
}
func getWinsize() (*unix.Winsize, error) {

	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return ws, nil
}
