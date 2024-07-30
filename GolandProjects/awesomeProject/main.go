package main

import "fmt"

func main() {
	fmt.Println(stringChecker("{}{}{esvsjg{}}"))
	r := Rectangle{
		width:  2,
		height: 5,
	}
	fmt.Println(r.width)
	r.setWidth(4)
	fmt.Println(r.width)
}

var i int

func stringChecker(s string) bool {
	matches := 0
	for _, char := range s {
		if char == '{' {
			matches += 1
		} else if char == '}' {
			matches -= 1
		}
	}
	return matches == 0
}

// Methods are functions which act on a certain type. Have to have this type to be able to call this func.

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) setWidth(w int) {
	r.width = w
}

//above acts on the original instantiated object
