package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := &Square{}
	sq.width = size
	sq.height = size
	return sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)

	expectArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected an area of ", expectArea, ", but go ", actualArea, "\n")
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}
