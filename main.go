package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	height = 600.0
	width  = 600.0
	scale  = 20.0
)

func clearLastDrawing(imd *imdraw.IMDraw, win *pixelgl.Window) {
	imd.Clear()
	win.Clear(colornames.Darkgrey)
}

func run() {
	var eat bool
	var snakeEnd int
	lastPressed := 'r' // default going right
	speed := scale     //20.0
	player := snake{}
	player.Init(4)

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, width, height),
		//VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Darkgrey)

	player.Set()

	// init food and draw it
	imdFood := imdraw.New(nil)
	imdFood.Color = colornames.Crimson
	imdFood.EndShape = imdraw.SharpEndShape
	food := pickLocation()
	imdFood.Push(food.Min, food.Max)
	imdFood.Rectangle(0)
	imdFood.Draw(win)

	var (
		frames = 0
		second = time.Tick(time.Second)
	)
	startTime := time.Now()
	exeTime := 1 * time.Second // just to start the loop
	for !win.Closed() {
		// check if the fps is over 20
		if exeTime < 50*time.Millisecond {
			exeTime = time.Since(startTime)
			continue
		}
		startTime = time.Now()
		win.Clear(colornames.Black)

		eat = player.EatFood(food)
		snakeEnd = len(player.Snake)
		if eat {
			food = pickLocation()
		}

		if win.JustPressed(pixelgl.KeyTab) {
			player.Grow()
		}

		//rotate the positions
		for i := 0; i < snakeEnd-1; i++ {
			player.Snake[i] = player.Snake[i+1]
		}

		player.Imd.Clear()
		imdFood.Clear()
		imdFood.Push(food.Min, food.Max)
		imdFood.Rectangle(0)
		imdFood.Draw(win)

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}

		// only turn the snake head
		if win.Pressed(pixelgl.KeyLeft) && lastPressed != 'r' {
			lastPressed = 'l'
			player.Snake[snakeEnd-1] = player.Snake[snakeEnd-1].Moved(pixel.V(-speed, 0))
		} else if win.Pressed(pixelgl.KeyRight) && lastPressed != 'l' {
			lastPressed = 'r'
			player.Snake[snakeEnd-1] = player.Snake[snakeEnd-1].Moved(pixel.V(speed, 0))
		} else if win.Pressed(pixelgl.KeyDown) && lastPressed != 'u' {
			lastPressed = 'd'
			player.Snake[snakeEnd-1] = player.Snake[snakeEnd-1].Moved(pixel.V(0, -speed))
		} else if win.Pressed(pixelgl.KeyUp) && lastPressed != 'd' {
			lastPressed = 'u'
			player.Snake[snakeEnd-1] = player.Snake[snakeEnd-1].Moved(pixel.V(0, speed))
		} else {
			switch lastPressed {
			case 'l':
				player.Snake[snakeEnd-1] = player.Snake[snakeEnd-1].Moved(pixel.V(-speed, 0))
			case 'r':
				player.Snake[snakeEnd-1] = player.Snake[snakeEnd-1].Moved(pixel.V(speed, 0))
			case 'd':
				player.Snake[snakeEnd-1] = player.Snake[snakeEnd-1].Moved(pixel.V(0, -speed))
			case 'u':
				player.Snake[snakeEnd-1] = player.Snake[snakeEnd-1].Moved(pixel.V(0, speed))
			default:
				player.Snake[snakeEnd-1] = player.Snake[snakeEnd-1].Moved(pixel.V(speed, 0))
			}

		}
		// check if the player reached the walls
		// reached the top
		if player.Snake[snakeEnd-1].Max.Y >= height+scale { //20 {
			player.Snake[snakeEnd-1].Max.Y = scale //20.0
			player.Snake[snakeEnd-1].Min.Y = 0.0
			// reached the right bounder
		} else if player.Snake[snakeEnd-1].Min.X >= width+scale { //20 {
			player.Snake[snakeEnd-1].Min.X = scale //20.0
			player.Snake[snakeEnd-1].Max.X = 0.0
			// reached the left bounder
		} else if player.Snake[snakeEnd-1].Max.X <= -scale { //-20.0 {
			player.Snake[snakeEnd-1].Max.X = width - scale //580.0
			player.Snake[snakeEnd-1].Min.X = width         //600.0
			// reached the buttom
		} else if player.Snake[snakeEnd-1].Min.Y <= -scale { //-20.0 {
			player.Snake[snakeEnd-1].Min.Y = height - scale //580.0
			player.Snake[snakeEnd-1].Max.Y = height         //600.0
		}

		for i, r := range player.Snake {
			numberOfColors := len(player.Colors)
			player.Imd.Color = player.Colors[i-(i/numberOfColors*numberOfColors)]
			player.Imd.Push(r.Min, r.Max)
			player.Imd.Rectangle(10.0)
		}
		player.Imd.Draw(win)
		win.Update()

		for i := 0; i < len(player.Snake)-2; i++ {
			if player.Interact(player.Snake[snakeEnd-1], player.Snake[i]) {
				player.Reset()
				lastPressed = 'r'
				continue
			}
		}
		// reset button
		if win.JustReleased(pixelgl.KeySpace) {
			player.Reset()
			lastPressed = 'r'
			continue
		}
		exeTime = time.Since(startTime)
	}
}

func main() {
	pixelgl.Run(run)
}

func pickLocation() pixel.Rect {
	var (
		MinX, MinY, MaxX, MaxY float64
	)
	cols := width / scale
	rows := height / scale
	randCol := float64(rand.Intn(int(cols)) * scale)
	randRow := float64(rand.Intn(int(rows)) * scale)

	MaxX = randCol         //scale + randRow
	MaxY = scale + randRow //randCol

	MinY = randRow
	MinX = scale + randCol

	food := pixel.R(MinX, MinY, MaxX, MaxY)

	return food
}
