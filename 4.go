package main

import "fmt"

func main() {
	var thn, thn1, thn2, thn3 int
	var kabisat, kabisat1, kabisat2, kabisat3 bool
	fmt.Scanln(&thn)
	fmt.Scanln(&thn1)
	fmt.Scanln(&thn2)
	fmt.Scanln(&thn3)
	kabisat = thn%400 == 0 || thn%4 == 0 && thn%100 != 0
	kabisat1 = thn1%400 == 0 || thn%4 == 0 && thn1%100 != 0
	kabisat2 = thn2%400 == 0 || thn%4 == 0 && thn2%100 != 0
	kabisat3 = thn3%400 == 0 || thn%4 == 0 && thn3%100 != 0
	fmt.Println(kabisat && kabisat1 && kabisat2 && kabisat3)
}
