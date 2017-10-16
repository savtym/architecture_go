package presenters

import (
//"../models/morze"
)

func OnDecodeClick(s string) string{
	return s
}

func OnChooseFileClick(s string) string{
	println(s)
	return "decoded from file"
}