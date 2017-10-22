package morze

import (
	"io/ioutil"
	"log"
	"os"
)

const defaultInputFilename = "cash/input.txt"

var InputFileName = defaultInputFilename

func input_reader() string {
	rawDataIn, err := ioutil.ReadFile(InputFileName)
	if err != nil {
		log.Fatal("Cannot load input:", err)
	}
	InputFileName = defaultInputFilename
	return string(rawDataIn)
}

func SetUserInput(input string) {
	if InputFileName == defaultInputFilename {
		os.Remove(InputFileName)
	}
	f, err := os.OpenFile(InputFileName, os.O_CREATE|os.O_WRONLY, 0644)
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
