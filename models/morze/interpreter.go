package morze

import (
	"io/ioutil"
	"log"
	"os"
)

const defaultInputFilename = "cash/input.txt"
//
//var InputFileName = defaultInputFilename

func input_reader() string {
	rawDataIn, err := ioutil.ReadFile(path.inputFileName)
	if err != nil {
		log.Fatal("Cannot load input:", err)
	}
	path.inputFileName = defaultInputFilename
	return string(rawDataIn)
}

func SetUserInput(input string) {
	if path.inputFileName == defaultInputFilename {
		os.Remove(path.inputFileName)
	}
	f, err := os.OpenFile(path.inputFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(input); err != nil {
		panic(err)
	}
}

func Interpreter() string {
	var sub_str, out_str string
	str := input_reader()
	for i := range str {
		tmp := string(str[i])
		if (tmp == "2") || (tmp == "3") || (tmp == "4") || (tmp == " ") || (tmp == "\n") {
			out_str = out_str + Decoder(sub_str)
			if string(str[i]) == "3" || string(str[i]) == " " {
				out_str += " "
			}
			sub_str = ""
		} else {
			sub_str += string(str[i])
		}
	}
	return out_str
}
