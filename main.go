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
		snakeEnd = len(player.ImdSnake)
		win.Clear(colornames.Darkgrey)

		eat = player.EatFood(FMaxX, FMaxY, FMinX, FMinY)

		if eat {
			FMaxX, FMaxY, FMinX, FMinY = pickLocation()
		}

		//rotate the positions
		for i := 0; i < snakeEnd-1; i++ {
			player.Snake[i] = player.Snake[i+1]
		}

		for i, imd := range player.ImdSnake {
			imdFood.Push(pixel.V(FMaxX, FMaxY))
			imdFood.Push(pixel.V(FMinX, FMinY))
			imdFood.Rectangle(0)
			imdFood.Draw(win)
			imd.Draw(win)
			win.Update()
			imd.Clear()
			imdFood.Clear()

			// only turn the snake head
			if i+1 == snakeEnd {
				if win.Pressed(pixelgl.KeyLeft) && lastPressed != 'r' {
					lastPressed = 'l'

					player.Snake[i].MaxX -= speed
					player.Snake[i].MinX -= speed
				} else if win.Pressed(pixelgl.KeyRight) && lastPressed != 'l' {
					lastPressed = 'r'

					player.Snake[i].MaxX += speed
					player.Snake[i].MinX += speed
				} else if win.Pressed(pixelgl.KeyDown) && lastPressed != 'u' {
					lastPressed = 'd'

					player.Snake[i].MaxY -= speed
					player.Snake[i].MinY -= speed
				} else if win.Pressed(pixelgl.KeyUp) && lastPressed != 'd' {
					lastPressed = 'u'

					player.Snake[i].MaxY += speed
					player.Snake[i].MinY += speed
				} else {
					switch lastPressed {
					case 'l':
						player.Snake[i].MaxX -= speed
						player.Snake[i].MinX -= speed
					case 'r':
						player.Snake[i].MaxX += speed
						player.Snake[i].MinX += speed
					case 'd':
						player.Snake[i].MaxY -= speed
						player.Snake[i].MinY -= speed
					case 'u':
						player.Snake[i].MaxY += speed
						player.Snake[i].MinY += speed
					default:
						player.Snake[i].MaxX += speed
						player.Snake[i].MinX += speed
					}

				}
			}
			// check if the player reached the walls
			// reached the top
			if player.Snake[i].MaxY >= height+20 {
				player.Snake[i].MaxY = scale //20.0
				player.Snake[i].MinY = 0.0
				// reached the right bounder
			} else if player.Snake[i].MinX >= width+20 {
				player.Snake[i].MinX = scale //20.0
				player.Snake[i].MaxX = 0.0
				// reached the left bounder
			} else if player.Snake[i].MaxX <= -20.0 {
				player.Snake[i].MaxX = width - scale //580.0
				player.Snake[i].MinX = width         //600.0
				// reached the buttom
			} else if player.Snake[i].MinY <= -20 {
				player.Snake[i].MinY = height - scale //580.0
				player.Snake[i].MaxY = height         //600.0
			}

			imd.Push(pixel.V(player.Snake[i].MaxX, player.Snake[i].MaxY))
			imd.Push(pixel.V(player.Snake[i].MinX, player.Snake[i].MinY))
			imd.Rectangle(0)
		}
		time.Sleep(75 * time.Millisecond)
	}
	defer func() { fmt.Println(lastPressed) }()
}

func main() {
	pixelgl.Run(run)
}

// run1 creates a rectangel that can move around in the corrrect way
func run1() {
	speed := 20.0

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 600, 600),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	var lastPressed rune

	rectP1 := []float64{0, 600}
	rectP2 := []float64{20, 580}
	imd := imdraw.New(nil)

	imd.Color = colornames.Blue
	imd.EndShape = imdraw.SharpEndShape
	imd.Push(pixel.V(rectP1[0], rectP1[1]))
	imd.Push(pixel.V(rectP2[0], rectP2[1]))
	imd.Rectangle(0)

	win.Clear(colornames.Darkgrey)
	//frameCounter := 0

	for !win.Closed() {
		//win.Clear(colornames.White)
		// frameCounter++
		// if frameCounter == 100 {
		// 	imd.Clear()
		// 	frameCounter = 0
		// 	fmt.Println("cleared!")
		// }
		imd.Draw(win)
		win.Update()

		clearLastDrawing(imd, win)

		// if rectP1[0] || rectP1[1] || rectP1[0] || rectP1[1] >= 600 {

		// }
		// if win.Pressed(pixelgl.KeySpace) {
		// 	imd.Clear()
		// 	win.Clear(colornames.Darkgrey)
		// }

		if win.JustPressed(pixelgl.KeyLeft) {
			lastPressed = 'l'
			rectP1[0] -= speed
			rectP2[0] -= speed
		} else if win.JustPressed(pixelgl.KeyRight) {
			lastPressed = 'r'
			rectP1[0] += speed
			rectP2[0] += speed
		} else if win.JustPressed(pixelgl.KeyDown) {
			lastPressed = 'd'
			rectP1[1] -= speed
			rectP2[1] -= speed
		} else if win.JustPressed(pixelgl.KeyUp) {
			lastPressed = 'u'
			rectP1[1] += speed
			rectP2[1] += speed
		} else {
			switch lastPressed {
			case 'l':
				rectP1[0] -= speed
				rectP2[0] -= speed
			case 'r':
				rectP1[0] += speed
				rectP2[0] += speed
			case 'd':
				rectP1[1] -= speed
				rectP2[1] -= speed
			case 'u':
				rectP1[1] += speed
				rectP2[1] += speed
			default:
				rectP1[0] += speed
				rectP2[0] += speed
			}

		}

		imd.Push(pixel.V(rectP1[0], rectP1[1]))
		imd.Push(pixel.V(rectP2[0], rectP2[1]))
		imd.Rectangle(0)

	}
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
