package screen

import (
	"bytes"
	"os"
	"time"

	"golang.org/x/sys/unix"
)

type Dimension struct {
	Width  int
	Height int
}

type Screen struct {
	Dimension
	Scale    int
	OnUpdate func(screen *Screen)
	FPS      int
	InPixel  *[][]uint8
	ASCIISTR string
}

func (screen *Screen) Start() {
	for {
		screen.Loop()
	}
}
func (screen *Screen) Loop() {
	screen.render()
	time.Sleep(time.Second / time.Duration(screen.FPS))
}
func (screen *Screen) render() {
	screen.OnUpdate(screen)
	print(string(screen.Convert2Ascii()))
}
func (screen *Screen) Convert2Ascii() []byte {
	table := []byte(screen.ASCIISTR)
	buf := new(bytes.Buffer)
	for i := 0; i < screen.Height; i++ {
		for j := 0; j < screen.Width; j++ {
			E := uint8(0)
			for s := 0; s < screen.Scale; s++ {
				E = E + (*screen.InPixel)[i*screen.Scale+s][j*screen.Scale+s]/uint8(screen.Scale)
			}

			E = E / uint8(screen.Scale)
			pos := 16 - int(E)*16/255
			_ = buf.WriteByte(table[pos])
		}
	}
	return buf.Bytes()
}

func NewScreen(scale int) (screen *Screen, err error) {
	//get console layout
	ws, err := getWinsize()
	if err != nil {
		return &Screen{}, err
	}

	s := make([][]uint8, ws.Row*uint16(scale))
	for i := uint16(0); i < ws.Row*uint16(scale); i++ {
		s[i] = make([]uint8, ws.Col*uint16(scale))
	}
	return &Screen{
		Dimension: Dimension{
			Width:  scale * int(ws.Col),
			Height: scale * int(ws.Row),
		},
		ASCIISTR: "MND8OZ$7I?+=~:,. ",
		Scale:    scale,
		InPixel:  &s,
		FPS:      60,
	}, nil
}
func getWinsize() (*unix.Winsize, error) {

	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return ws, nil
}
