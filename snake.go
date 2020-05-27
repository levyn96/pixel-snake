package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type snake struct {
	Snake  []pixel.Rect //position
	Imd    *imdraw.IMDraw
	Colors []color.RGBA
}

func (s *snake) Init(n int) {
	Colors := []color.RGBA{colornames.Brown, colornames.Blue, colornames.Crimson, colornames.Green}
	s.Colors = Colors
	var x float64
	start := []float64{0, 600, 20, 580}
	imd := imdraw.New(nil)
	imd.Color = colornames.Greenyellow
	imd.EndShape = imdraw.RoundEndShape
	s.Imd = imd
	for i := 0; i < n; i++ {
		x = float64(i) * scale
		s.Snake = append(s.Snake, pixel.R(x+start[2], start[3], x+start[0], start[1]))
	}
}

func (s *snake) Set() {
	for i, r := range s.Snake {
		// pick the color with respect to the index (0-brown,1-blue,2-crimson,3-green, 4 is brown again and so on)
		s.Imd.Color = s.Colors[i-(i/len(s.Colors)*len(s.Colors))]
		s.Imd.Push(r.Min, r.Max)
		s.Imd.Rectangle(10.0)
	}
}

func (s *snake) EatFood(food pixel.Rect) bool {
	diff := s.Snake[len(s.Snake)-1].Center().Sub(food.Center())
	if (diff.X < 5 && diff.X > -5) && (diff.Y < 5 && diff.Y > -5) {
		s.Grow()
		return true
	}
	return false
}

func (s *snake) Grow() {
	array1 := []pixel.Rect{s.Snake[0]}
	for _, p := range s.Snake {
		array1 = append(array1, p)
	}
	s.Snake = array1
}

func (s *snake) Reset() {
	s.Imd = nil
	s.Snake = nil
	s.Init(4)
	s.Set()
}

func (s *snake) Interact(r1, r2 pixel.Rect) bool {
	diff := r1.Center().Sub(r2.Center())
	if (diff.X < 5 && diff.X > -5) && (diff.Y < 5 && diff.Y > -5) {
		return true
	}
	return false
}
