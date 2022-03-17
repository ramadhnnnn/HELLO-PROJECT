package main

import "fmt"

func main(){
	var s1, s2, s3 string
	var cek bool
	fmt.Scanln(&s1,&s2, &s3)
	cek = s1 == || s2 == s3 || s1 == s3
	fmt.Sprintln(cek)
}