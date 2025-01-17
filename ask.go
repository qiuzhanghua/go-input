package input

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// Ask asks the user for input using the given query. The response is
// returned as string. Error is returned based on the given option.
// If Loop is true, it continues to ask until it receives valid input.
//
// If the user sends SIGINT (Ctrl+C) while reading input, it catches
// it and return it as an error.
func (i *UI) Ask(query string, opts *Options) (string, error) {
	i.once.Do(i.setDefault)

	// Display the query to the user.
	fmt.Fprintf(i.Writer, "%s", query)

	// resultStr and resultErr are return val of this function
	var resultStr string
	var resultErr error

	loopCount := 0
	for {
		loopCount++

		// Construct the instruction to user.
		var buf bytes.Buffer
		if !opts.HideOrder || loopCount > 1 {
			buf.WriteString(T("go-input.ask.enter-value"))
		}

		if opts.Default != "" && !opts.HideDefault {
			defaultVal := opts.Default
			if opts.MaskDefault {
				defaultVal = maskString(defaultVal)
			}
			buf.WriteString(T("go-input.ask.default-value", defaultVal))
		}

		// Display the instruction to user and ask to input.
		buf.WriteString(": ")
		fmt.Fprint(i.Writer, buf.String())

		// Read user input from UI.Reader.
		line, err := i.read(opts.readOpts())
		if err != nil {
			resultErr = err
			break
		}

		line = strings.TrimSpace(line)

		// line is empty but default is provided returns it
		if line == "" && opts.Default != "" {
			resultStr = opts.Default
			break
		}

		if line == "" && opts.Required {
			if !opts.Loop {
				resultErr = errors.New(T("go-input.ErrEmpty"))
				break
			}

			fmt.Fprint(i.Writer, T("go-input.ask.ErrInputEmpty"))
			continue
		}

		// validate input by custom function
		validate := opts.validateFunc()
		if err := validate(line); err != nil {
			if !opts.Loop {
				resultErr = err
				break
			}

			fmt.Fprint(i.Writer, T("go-input.ask.ErrInputInvalidate", err))
			continue
		}

		// Reach here means it gets ideal input.
		resultStr = line
		break
	}

	// Insert the new line for next output
	fmt.Fprint(i.Writer, "\n")

	return resultStr, resultErr
}
