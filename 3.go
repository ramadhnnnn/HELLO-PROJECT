package main

import "fmt"

func main(){
	var gr, gr1, gr2, gr3 float64
	var kg, ons, pon, float64
	var kg1, ons1, pon1, float64
	var kg2, ons2, pon2, float64
	var kg3, ons3, pon3, float64
	fmt.Scanln(&gr)
	fmt.Scanln(&gr1)
	fmt.Scanln(&gr2)
	fmt.Scanln(&gr3)
	kg = gr / 1000
	pon = gr / 453.592
	ons = gr / 28.3495
	kg1= gr1 / 1000
	pon1 = gr1 / 453.592
	ons1 = gr1 / 28.3495
	kg2 = gr2 / 1000
	pon2 = gr2 / 453.592
	ons2 = gr2 / 28.3495
	kg3 = gr3 / 1000
	pon3 = gr3 / 453.592
	ons3 = gr3 / 28.3495
	fmt.Println(kg, pon, ons)
	fmt.Println(kg1, pon1, ons1)
	fmt.Println(kg2, pon2, ons2)
	fmt.Println(kg3, pon3, ons3)
}