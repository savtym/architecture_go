package tests

import (
	"testing"
	"../models/morze"
)

type testDecoderPair struct {
	input  string
	output string
}

var decoderTests = []testDecoderPair{
	{"",""},
	{" ",""},
	{"011","w"},
	{"0112","w"},
	{"0113","w"},
	{"0114","w"},
	{"0115",""},
}

func TestDecoder(t *testing.T) {
	for _, pair := range decoderTests {
		morze.InputFileName = "../cash/input.txt"
		morze.SetUserInput(pair.input)
		v := morze.Interpreter()
		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}

