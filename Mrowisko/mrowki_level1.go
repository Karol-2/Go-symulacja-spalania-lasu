package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

const (
	winTitle = "Mrowki"
)

func run() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	win, err := sdl.CreateWindow(
		winTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		1024,
		512,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer win.Destroy()

	for {
		event := sdl.WaitEvent()
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return nil
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_q {
				return nil
			}
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
