package screen

import (
	"os"

	"golang.org/x/sys/unix"
)

type Dimension struct {
	Width  int
	Height int
}
type Screen struct {
	Dimension
	OnUpdate func(screen *Screen)
	SwapSlice
	FPS int
	spf float64
}
type SwapSlice struct {
	A    *[][]byte
	B    *[][]byte
	Swap chan bool
}

func (screen *Screen) render() {
	screen.OnUpdate(screen)
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
	a := make([][]byte, ws.Col)
	b := make([][]byte, ws.Col)
	for i := uint16(0); i < ws.Row; i++ {
		a[i] = make([]byte, ws.Row)
		b[i] = make([]byte, ws.Row)
	}
	return &Screen{
		SwapSlice: SwapSlice{
			A:    &a,
			B:    &b,
			Swap: make(chan bool),
		},
		Dimension: Dimension{
			Width:  int(ws.Col),
			Height: int(ws.Row),
		},
	}, nil
}
func getWinsize() (*unix.Winsize, error) {

	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return ws, nil
}
