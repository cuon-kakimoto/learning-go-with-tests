package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Redius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface
// Circle has a method called Area that returns a float64 so it satisfies the Shape interface
// string does not have such a method, so it doesn't satisfy the interface
// Hack: interfaceを定義することで型Shapeであるとする
type Shape interface {
	Area() float64
}

// HACK: width, heightの引数だと四角形の外縁であることが明示できないため、引数として四角形を明示
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// HACK:
// func Area(circle Circle) float64 { ... }
// func Area(rectangle Rectangle) float64 { ... }
// 引数を変えた関数の二重定義はできないため、メソッド定義する

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Redius * c.Redius
}

func (t Triangle) Area() float64 {
	// GOOD
	return (t.Base * t.Height) * 0.5
}
