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
	var (
		FMaxX, FMaxY, FMinX, FMinY float64
	)
	var eat bool
	var snakeEnd int
	lastPressed := 'r'
	speed := scale //20.0
	player := snake{}
	player.Init(4)
	fmt.Println(player)

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 600, 600),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Darkgrey)

	player.Set()

	// init food and draw it
	imdFood := imdraw.New(nil)
	imdFood.Color = colornames.Green
	imdFood.EndShape = imdraw.SharpEndShape
	FMaxX, FMaxY, FMinX, FMinY = pickLocation()
	fmt.Println(FMaxX, FMaxY, FMinX, FMinY)
	imdFood.Push(pixel.V(FMaxX, FMaxY))
	imdFood.Push(pixel.V(FMinX, FMinY))
	imdFood.Rectangle(0)
	imdFood.Draw(win)

	for !win.Closed() {
		snakeEnd = len(player.Snake)
		win.Clear(colornames.Darkgrey)

		eat = player.EatFood(FMaxX, FMaxY, FMinX, FMinY)

		if eat {
			FMaxX, FMaxY, FMinX, FMinY = pickLocation()
		}

		//rotate the positions
		for i := 0; i < snakeEnd-1; i++ {
			player.Snake[i] = player.Snake[i+1]
		}

		imdFood.Push(pixel.V(FMaxX, FMaxY))
		imdFood.Push(pixel.V(FMinX, FMinY))
		imdFood.Rectangle(0)
		imdFood.Draw(win)
		win.Update()
		player.Imd.Clear()
		imdFood.Clear()

		// only turn the snake head
		if win.Pressed(pixelgl.KeyLeft) && lastPressed != 'r' {
			lastPressed = 'l'

			player.Snake[snakeEnd-1].MaxX -= speed
			player.Snake[snakeEnd-1].MinX -= speed
		} else if win.Pressed(pixelgl.KeyRight) && lastPressed != 'l' {
			lastPressed = 'r'

			player.Snake[snakeEnd-1].MaxX += speed
			player.Snake[snakeEnd-1].MinX += speed
		} else if win.Pressed(pixelgl.KeyDown) && lastPressed != 'u' {
			lastPressed = 'd'

			player.Snake[snakeEnd-1].MaxY -= speed
			player.Snake[snakeEnd-1].MinY -= speed
		} else if win.Pressed(pixelgl.KeyUp) && lastPressed != 'd' {
			lastPressed = 'u'

			player.Snake[snakeEnd-1].MaxY += speed
			player.Snake[snakeEnd-1].MinY += speed
		} else {
			switch lastPressed {
			case 'l':
				player.Snake[snakeEnd-1].MaxX -= speed
				player.Snake[snakeEnd-1].MinX -= speed
			case 'r':
				player.Snake[snakeEnd-1].MaxX += speed
				player.Snake[snakeEnd-1].MinX += speed
			case 'd':
				player.Snake[snakeEnd-1].MaxY -= speed
				player.Snake[snakeEnd-1].MinY -= speed
			case 'u':
				player.Snake[snakeEnd-1].MaxY += speed
				player.Snake[snakeEnd-1].MinY += speed
			default:
				player.Snake[snakeEnd-1].MaxX += speed
				player.Snake[snakeEnd-1].MinX += speed
			}

		}
		// check if the player reached the walls
		// reached the top
		if player.Snake[snakeEnd-1].MaxY >= height+20 {
			player.Snake[snakeEnd-1].MaxY = scale //20.0
			player.Snake[snakeEnd-1].MinY = 0.0
			// reached the right bounder
		} else if player.Snake[snakeEnd-1].MinX >= width+20 {
			player.Snake[snakeEnd-1].MinX = scale //20.0
			player.Snake[snakeEnd-1].MaxX = 0.0
			// reached the left bounder
		} else if player.Snake[snakeEnd-1].MaxX <= -20.0 {
			player.Snake[snakeEnd-1].MaxX = width - scale //580.0
			player.Snake[snakeEnd-1].MinX = width         //600.0
			// reached the buttom
		} else if player.Snake[snakeEnd-1].MinY <= -20.0 {
			player.Snake[snakeEnd-1].MinY = height - scale //580.0
			player.Snake[snakeEnd-1].MaxY = height         //600.0
		}
		for _, p := range player.Snake {
			player.Imd.Push(pixel.V(p.MaxX, p.MaxY))
			player.Imd.Push(pixel.V(p.MinX, p.MinY))
			player.Imd.Rectangle(0)
			player.Imd.Draw(win)
			win.Update()
		}
		time.Sleep(75 * time.Millisecond)
	}
	defer func() { fmt.Println(lastPressed) }()
}

func main() {
	pixelgl.Run(run)
}

func pickLocation() (MaxX, MaxY, MinX, MinY float64) {

	cols := width / scale
	rows := height / scale
	randCol := float64(rand.Intn(int(cols)) * scale)
	randRow := float64(rand.Intn(int(rows)) * scale)

	MaxX = randCol         //scale + randRow
	MaxY = scale + randRow //randCol

	MinY = randRow
	MinX = scale + randCol

	return
}
