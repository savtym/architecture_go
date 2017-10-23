package main

import (
	"./models/morze"
	"./views/src"
)

func main() {
	morze.DefaultValues()
	src.Show()
}
