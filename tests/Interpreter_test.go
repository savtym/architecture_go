package tests

import (
	"testing"
	"../models/morze"
)

type testInterpreterPair struct {
	input  string
	output string
}

var InterpreterTests = []testInterpreterPair{
	{"10002","b"},
	{"011","p"},
	{"0112","wError"},
	{"10102","c"},
}

func TestInterpreter(t *testing.T) {
	for _, pair := range InterpreterTests {
		morze.SetDictFileName("../dictionary/eng.json")
		morze.SetInputFileName("../cash/input.txt")
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


func BenchmarkTest(b *testing.B) {
	morze.SetDictFileName("../dictionary/eng.json")
	for i := 0; i < b.N; i++ {
		morze.SetInputFileName("../cash/hello.txt")
		morze.Interpreter()
    }
}


func BenchmarkTestParallel(b *testing.B) {
	morze.SetDictFileName("../dictionary/eng.json")
	for i := 0; i < b.N; i++ {
		morze.SetInputFileName("../cash/hello.txt")
		morze.Threader()
    }
}