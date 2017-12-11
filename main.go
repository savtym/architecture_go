package main

import (
	"./models/morze"
	// "./views/src"
)

func main() {
	morze.DefaultValues()
	morze.Threader()
	// src.Show()
}
