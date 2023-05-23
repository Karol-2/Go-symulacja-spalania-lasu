package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle = "SDL Window"
)

var (
	bgColor = sdl.Color{R: 0, G: 0, B: 0, A: 255}       // black
	fgColor = sdl.Color{R: 255, G: 255, B: 255, A: 255} // white
)

func run() error {
	// Initialize SDL2
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	// Create window
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

	// Create renderer
	renderer, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return fmt.Errorf("could not create renderer: %v", err)
	}
	defer renderer.Destroy()

	// Set drawing color
	if err := renderer.SetDrawColor(bgColor.R, bgColor.G, bgColor.B, bgColor.A); err != nil {
		return fmt.Errorf("could not set drawing color: %v", err)
	}

	// Clear the window with the background color
	renderer.Clear()

	// Create random number generator
	rand.Seed(time.Now().Unix())

	// Create starting position for the point in the center of the window
	x := int32(1024 / 2)
	y := int32(512 / 2)

	// Draw the initial point
	if err := drawPoint(renderer, x, y, fgColor); err != nil {
		return fmt.Errorf("could not draw point: %v", err)
	}

	// Loop until the point hits the edge of the window or the program is closed
	for {
		// Wait for an event
		event := sdl.WaitEvent()

		// Handle the event
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return nil
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_q {
				return nil
			}
		default:
			// Erase the old point
			if err := drawPoint(renderer, x, y, bgColor); err != nil {
				return fmt.Errorf("could not draw point: %v", err)
			}

			// Move the point
			x += int32(rand.Intn(3) - 1)
			y += int32(rand.Intn(3) - 1)

			// Draw the new point
			if err := drawPoint(renderer, x, y, fgColor); err != nil {
				return fmt.Errorf("could not draw point: %v", err)
			}

			// Check if the point hit the edge of the window
			if x <= 0 || x >= 1024 || y <= 0 || y >= 512 {
				return nil
			}
		}
	}
}

// Funkcja drawPoint rysuje na ekranie punkt o współrzędnych (x, y) i kolorze color.
func drawPoint(renderer *sdl.Renderer, x, y int32, color sdl.Color) error {
	// Ustawiamy kolor rysowania na zadany kolor.
	if err := renderer.SetDrawColor(color.R, color.G, color.B, color.A); err != nil {
		return err
	}

	// Rysujemy punkt o zadanych współrzędnych.
	if err := renderer.DrawPoint(x, y); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
