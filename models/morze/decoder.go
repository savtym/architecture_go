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

const dictFilename = "models/morze/eng.json"

func json_reader() []Symb_code {
	rawDataIn, err := ioutil.ReadFile(dictFilename)
	if err != nil {
		log.Fatal("Cannot load dictionary:", err)
	}
	var pingJSON []Symb_code
	err = json.Unmarshal([]byte(rawDataIn), &pingJSON)
	if err != nil {
		log.Fatal("Invalid dictionary format:", err)
	}
	return pingJSON
}

func Decoder(buffer string) string {
	dictionary := json_reader()
	for i := range dictionary {
		if buffer == string(dictionary[i].Symbol_code) {
			return dictionary[i].Symbol
		}
	}
	return "Error"
}
