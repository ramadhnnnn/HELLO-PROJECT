package main

import "fmt"

func main() {
	var x, y, x1, y1 int
	var x2, y2, x3, y3 int
	var x4, y4, x5, y5 int

	var gradient, gradient1, gradient2 float64
	fmt.Scanln(&x, &y, &x1, &y1)
	fmt.Scanln(&x2, &y2, &x3, &y3)
	fmt.Scanln(&x4, &y4, &x5, &y5)
	gradient = float64(y-y1) / float64(x-x1)
	gradient1 = float64(y2-y3) / float64(x2-x3)
	gradient2 = float64(y4-y5) / float64(x4-x5)
	fmt.Println(gradient)
	fmt.Println(gradient1)
	fmt.Println(gradient2)
}
