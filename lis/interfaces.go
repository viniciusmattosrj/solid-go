package main

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}