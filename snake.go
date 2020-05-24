package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type snake struct {
	Snake    []position
	ImdSnake []*imdraw.IMDraw
}

type position struct {
	// MinX, MinY, MaxX, MaxY float64
	MaxX, MaxY, MinX, MinY float64
}

func (s *snake) Init(n int) {
	var x float64
	start := []float64{0, 600, 20, 580}
	var positions []position
	for i := 0; i < n; i++ {
		imd := imdraw.New(nil)
		imd.Color = colornames.Blue
		imd.EndShape = imdraw.SharpEndShape
		tmp := imd
		s.ImdSnake = append(s.ImdSnake, tmp)
		for p := 0; p < n; p++ {
			x = float64(p) * 25
			positions = append(positions, position{x + start[0], start[1], x + start[2], start[3]})
		}
		s.Snake = append(s.Snake, positions[i])
	}
}

func (s *snake) Set() {
	for i, t := range s.ImdSnake {
		t.Push(pixel.V(s.Snake[i].MaxX, s.Snake[i].MaxY))
		t.Push(pixel.V(s.Snake[i].MinX, s.Snake[i].MinY))
		t.Rectangle(0)
	}
}

func (s *snake) EatFood(x1, y1, x2, y2 float64) bool {

	// if s.Snake[len(s.Snake)-1].MaxX == x1 && s.Snake[len(s.Snake)-1].MaxY == y1 ||
	// 	s.Snake[len(s.Snake)-1].MinX == x2 && s.Snake[len(s.Snake)-1].MinY == y2 {
	if (s.Snake[len(s.Snake)-1].MaxX+5.0 >= x1 && s.Snake[len(s.Snake)-1].MaxX-5.0 <= x1) &&
		(s.Snake[len(s.Snake)-1].MaxY+5.0 >= y1 && s.Snake[len(s.Snake)-1].MaxY-5.0 <= y1) {
		return true
	} else if (s.Snake[len(s.Snake)-1].MinX+5.0 >= x2 && s.Snake[len(s.Snake)-1].MinX-5.0 <= x2) &&
		(s.Snake[len(s.Snake)-1].MinY+5.0 >= y2 && s.Snake[len(s.Snake)-1].MinY-5.0 <= y2) {
		return true
	}
	return false
}

func (s *snake) Grow() {
	var array1 []*imdraw.IMDraw
	imd := imdraw.New(nil)
	imd.Color = colornames.Blue
	imd.EndShape = imdraw.SharpEndShape
	tmp := imd
	// add new piece at the begening
	array1 = append(array1, tmp)
	for _, p := range s.ImdSnake {
		// copy the old snake
		array1 = append(array1, p)
	}
	s.ImdSnake = array1
}
