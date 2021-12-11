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
		"go-input.ErrEmpty":               "default value is not provided but input is empty",
		"go-input.ErrNotNumber":           "input must be number",
		"go-input.ErrOutOfRange":          "input is out of range",
		"go-input.ErrInterrupted":         "interrupted",
		"go-input.ErrReadInput":           "failed to read the input: %s",
		"go-input.read.must-be-file":      "reader must be a file",
		"go-input.read-unix.not-terminal": "file descriptor %d is not a terminal",
		"go-input.ask.enter-value":        "\nEnter a value",
		"go-input.ask.default-value":      "(Default is %s)",
		"go-input.ask.ErrInputEmpty":      "Input must not be empty.\n\n",
		"go-input.ask.ErrInputInvalidate": "Failed to validate input string: %s\n\n",
		"go-input.select.enter-number":    "Enter a number",
		"go-input.select.default":         " (Default is %d)",
		"go-input.select.number-empty":    "Input must not be empty. Answer by a number.\n\n",
		"go-input.select.not-number":      "%q is not a valid input. Answer by a number.\n\n",
		"go-input.select.invalid-choice":  "%q is not a valid choice. Choose a number from 1 to %d.\n\n",
		"go-input.select.invalid-string":  "Failed to validate input string: %s\n\n",
	}
)
