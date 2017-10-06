package morze

import (
	"io/ioutil"
	"log"
)

const inputFilename = "cash/input.txt"

func input_reader() string {
	rawDataIn, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatal("Cannot load input:", err)
	}
	return string(rawDataIn)
}

func Interpreter() string {
	var sub_str, out_str string
	str := input_reader()
	for i := range str {
		tmp := string(str[i])
		if (tmp == "2") || (tmp == "3") || (tmp == "4") {
			out_str = out_str + Decoder(sub_str)
			if string(str[i]) == "3" {
				out_str += " "
			}
			sub_str = ""
		} else {
			sub_str += string(str[i])
		}
	}
	return out_str
}
