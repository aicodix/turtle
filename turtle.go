/*
A little fun with Turtle graphics!

Copyright 2017 Ahmet Inan <inan@aicodix.de>
*/

package main

import (
	"os"
	"fmt"
	"image"
	"image/png"
	"image/color"
	"strings"
)

func die(err interface{}) {
	fmt.Println(err)
	os.Exit(1)
}

func rot90(p image.Point) image.Point {
	return image.Point{-p.Y, p.X}
}

func rot270(p image.Point) image.Point {
	return image.Point{p.Y, -p.X}
}

func draw(input string) []image.Point {
	position := image.Point{0, 0}
	delta := image.Point{1, 0}
	var path []image.Point
	path = append(path, position)
	for _, c := range input {
		switch c {
		case 'F': // draw forward
			position = position.Add(delta)
			path = append(path, position)
		case '-': // turn left 90°
			delta = rot90(delta)
		case '+': // turn right 90°
			delta = rot270(delta)
		}
	}
	return path
}

func bounds(path []image.Point) image.Rectangle {
	var rect image.Rectangle
	for _, pos := range path {
		if pos.X < rect.Min.X { rect.Min.X = pos.X }
		if pos.Y < rect.Min.Y { rect.Min.Y = pos.Y }
		if pos.X > rect.Max.X { rect.Max.X = pos.X }
		if pos.Y > rect.Max.Y { rect.Max.Y = pos.Y }
	}
	return rect
}

func line(img *image.NRGBA, a, b image.Point) {
	if a.X == b.X {
		if a.Y > b.Y { a, b = b, a }
		for y := a.Y; y <= b.Y; y++ {
			img.Set(a.X, y, color.Black)
		}
	} else if a.Y == b.Y {
		if a.X > b.X { a, b = b, a }
		for x := a.X; x <= b.X; x++ {
			img.Set(x, a.Y, color.Black)
		}
	} else {
		die("can only do horizontal and vertical lines")
	}
}

func main() {
	scale := 10
	margin := 10
/*
	// Dragon curve
	rules := strings.NewReplacer(
		"X", "X+YF+",
		"Y", "-FX-Y")
	axiom := "FX"
	level := 12

	// Koch curve
	rules := strings.NewReplacer(
		"F", "F+F-F-F+F")
	axiom := "F"
	level := 4
*/
	// Hilbert curve
	rules := strings.NewReplacer(
		"A", "-BF+AFA+FB-",
		"B", "+AF-BFB-FA+")
	axiom := "A"
	level := 5

	tmp := axiom
	for i := 0; i < level; i++ {
		tmp = rules.Replace(tmp)
	}
	// fmt.Println(tmp)
	path := draw(tmp)
	rect := bounds(path)
	rect.Min = rect.Min.Mul(scale).Sub(image.Point{margin, margin})
	rect.Max = rect.Max.Mul(scale).Add(image.Point{margin, margin})
	img := image.NewNRGBA(rect)
	prev := path[0].Mul(scale)
	for _, pos := range path[1:] {
		scaled := pos.Mul(scale)
		line(img, scaled, prev)
		prev = scaled
	}

	name := "output.png"
	file, err := os.Create(name)
	if err != nil { die(err) }
	if err := png.Encode(file, img); err != nil { die(err) }
}

