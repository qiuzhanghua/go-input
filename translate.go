package input

import "fmt"

type Format func(string, ...interface{}) string

var T Format = func(stringId string, a ...interface{}) string {
	if format, ok := messages[stringId]; !ok {
		return fmt.Sprintf(format, a...)
	}
	return ""
}

var (
	messages = map[string]string{
		"go-input.read.must-be-file": "reader must be a file",
		"go-input.ErrEmpty":          "default value is not provided but input is empty",
		"go-input.ErrNotNumber":      "input must be number",
		"go-input.ErrOutOfRange":     "input is out of range",
		"go-input.ErrInterrupted":    "interrupted",
	}
)
