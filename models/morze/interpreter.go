package morze

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

const defaultInputFilename = "cash/input.txt"
//
//var InputFileName = defaultInputFilename

func Threader() string {
	var arr []string
	str := input_reader()
	length := int(len(str) / threadsNumber)
	beginPos := 0

	for i := length; i < len(str) - 1; i++ {
		if str[i] == '3' {
			arr = append(arr, str[beginPos:i+1])
			beginPos = i + 1
			i += length
		} else if str[i] == '4' {
			arr = append(arr, str[beginPos:])
		}
	}

	for item := range arr {

	}


	return ""
}

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
		if (tmp == "2") || (tmp == "3") || (tmp == "4") || (tmp == " ") {
			out_str = out_str + Decoder(sub_str)
			if string(str[i]) == "3"  {
				out_str += " "
			}
			sub_str = ""
		} else {
			sub_str += string(str[i])
		}
	}
	return out_str
}

func text(str string) {
	var sub_str, out_str string
	for i := range str {
		tmp := string(str[i])
		if (tmp == "2") || (tmp == "3") || (tmp == "4") || (tmp == " ") {
			out_str = out_str + Decoder(sub_str)
			if string(str[i]) == "3"  {
				out_str += " "
			}
			sub_str = ""
		} else {
			sub_str += string(str[i])
		}
	}
} 
