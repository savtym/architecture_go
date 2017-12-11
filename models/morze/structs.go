package morze

var path = PathCode{}

const threadsNumber = 10


type commonMemory struct {
	maps map[int]string
}

type Symb_code struct {
	Symbol_code string
	Symbol      string
}

type PathCode struct {
	dictFilename string
	inputFileName string
}

func SetDictFileName(dictFilename string) {
	path.dictFilename = dictFilename
}

func SetInputFileName(inputFileName string) {
	path.inputFileName = inputFileName
}


func DefaultValues() {
	path.dictFilename = "dictionary/eng.json"
	path.inputFileName = "cash/hello.txt"
}