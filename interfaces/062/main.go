// Interfaces - Intro

// An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other type must have.
// Interfaces are a concept that at first may need some time to be understood, take all your time.

package main

import (
	"context"
	"fmt"
)

// This first exercise we are going to create an interface, no execution!
// Creating our first interface:
// declare an interface called "Geometry"
type Geometry interface {
	// create a function signature area(), it's return type will be float64
	area(width float64, height float64) float64
}

func (rect Rect) area(width float64, height float64) float64 {
	return width * height
}

// A rectangle struct
type Rect struct {
	width, height float64
}

func doSomething(ctx context.Context) {
	var rect Rect = Rect{width: 100, height: 200}
	fmt.Println(rect.area(200, 300))
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}
