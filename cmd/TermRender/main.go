package main

import (
	"time"

	"github.com/shnidi/TermRender/pkg/screen"
)

func main() {
	screen.NewScreen()
	time.Sleep(5 * time.Second)
}
