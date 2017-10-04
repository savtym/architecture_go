package morze

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Symb_code struct {
	Symbol_code string
	Symbol      string
}

type Dictionary struct {
	Symb_codes []Symb_code
}

const dictFilename = "rus.json"

func json_reader() Dictionary {
	rawDataIn, err := ioutil.ReadFile(dictFilename)
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}

	var dictionary Dictionary
	print(rawDataIn)
	err = json.Unmarshal(rawDataIn, &dictionary)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}
	return dictionary
}

func decoder(buffer string) string {
	dictionary := json_reader()
	return "morze"
}
