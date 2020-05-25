package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type snake struct {
	Snake []position
	Imd   *imdraw.IMDraw
}

type position struct {
	// MinX, MinY, MaxX, MaxY float64
	MaxX, MaxY, MinX, MinY float64
}

func (s *snake) Init(n int) {
	var x float64
	start := []float64{0, 600, 20, 580}
	var positions []position
	imd := imdraw.New(nil)
	imd.Color = colornames.Greenyellow
	imd.EndShape = imdraw.RoundEndShape
	s.Imd = imd
	for i := 0; i < n; i++ {
		for p := 0; p < n; p++ {
			x = float64(p) * scale
			positions = append(positions, position{x + start[0], start[1], x + start[2], start[3]})
		}
		s.Snake = append(s.Snake, positions[i])
	}
}

func (s *snake) Set() {
	for _, t := range s.Snake {
		s.Imd.Push(pixel.V(t.MaxX, t.MaxY))
		s.Imd.Push(pixel.V(t.MinX, t.MinY))
		s.Imd.Rectangle(10.0)
	}
}

func (s *snake) EatFood(x1, y1, x2, y2 float64) bool {

	// if s.Snake[len(s.Snake)-1].MaxX == x1 && s.Snake[len(s.Snake)-1].MaxY == y1 ||
	// 	s.Snake[len(s.Snake)-1].MinX == x2 && s.Snake[len(s.Snake)-1].MinY == y2 {
	if (s.Snake[len(s.Snake)-1].MaxX+5.0 >= x1 && s.Snake[len(s.Snake)-1].MaxX-5.0 <= x1) &&
		(s.Snake[len(s.Snake)-1].MaxY+5.0 >= y1 && s.Snake[len(s.Snake)-1].MaxY-5.0 <= y1) {
		s.Grow()
		return true
	} else if (s.Snake[len(s.Snake)-1].MinX+5.0 >= x2 && s.Snake[len(s.Snake)-1].MinX-5.0 <= x2) &&
		(s.Snake[len(s.Snake)-1].MinY+5.0 >= y2 && s.Snake[len(s.Snake)-1].MinY-5.0 <= y2) {
		s.Grow()
		return true
	}
	return false
}

func (s *snake) Grow() {
	array1 := []position{s.Snake[0]}
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

func (s *snake) Interact(r1, r2 position) bool {
	if (r1.MaxX+5.0 >= r2.MaxX && r1.MaxX-5.0 <= r2.MaxX) &&
		(r1.MaxY+5.0 >= r2.MaxY && r1.MaxY-5.0 <= r2.MaxY) {
		return true
	} else if (r1.MinX+5.0 >= r2.MinX && r1.MinX-5.0 <= r2.MinX) &&
		(r1.MinY+5.0 >= r2.MinY && r1.MinY-5.0 <= r2.MinY) {
		return true
	}
	return false
}
