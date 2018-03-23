package main

import "fmt"

func main() {
	m := make([]string, 3, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Todd]
}

func changeMe(z []string) {
	z[0] = "Todd1212121221212"
	z[1] = "Todd555555"

	fmt.Println(z) // [Todd]
}
