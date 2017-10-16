package main

import (
	"fmt"
	"./models/morze"
	"./views/src"
)

func main() {
	fmt.Println(morze.Interpreter())
	src.ShowWindow()


}
