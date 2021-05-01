package main

import (
	"time"

	"github.com/shnidi/TermRender/pkg/screen"
)

func main() {
	s, _ := screen.NewScreen(1)
	var l *int
	x := 0
	l = &x
	s.OnUpdate = func(screen *screen.Screen) {
		if (*l) < screen.Width-1 {
			(*l) = (*l) + 1

			(*screen.InPixel)[5][*l] = 255
		}
	}
	s.Start()
	time.Sleep(5 * time.Second)
}
