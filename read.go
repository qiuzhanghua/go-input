package input

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
)

// readOptions is option for read func
type readOptions struct {
	// mask hides user input and will be matched by maskVal.
	mask    bool
	maskVal string
}

// read: reads input from UI.Reader
func (i *UI) read(opts *readOptions) (string, error) {
	i.once.Do(i.setDefault)

	// sigCh is channel which is watch Interrupted signal (SIGINT)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	var resultStr string
	var resultErr error
	doneCh := make(chan struct{})

	go func() {
		defer close(doneCh)

		if opts.mask {
			f, ok := i.Reader.(*os.File)
			if !ok {
				resultErr = fmt.Errorf(T("go-input.read.must-be-file"))
				return
			}

			i.mask, i.maskVal = opts.mask, opts.maskVal
			resultStr, resultErr = i.rawRead(f)
		} else {
			line, err := i.bReader.ReadString('\n')
			if err != nil && err != io.EOF {
				resultErr = errors.New(T("go-input.ErrReadInput", err))
			}

			resultStr = strings.TrimSuffix(line, LineSep)
			// brute force for the moment
			resultStr = strings.TrimSuffix(line, "\n")
		}
	}()

	select {
	case <-sigCh:
		return "", errors.New(T("go-input.ErrInterrupted"))
	case <-doneCh:
		return resultStr, resultErr
	}
}

// rawReadline tries to return a single line, not including the end-of-line
// bytes with raw Mode (without prompting nothing). Or if provided show some
// value instead of actual value.
func (i *UI) rawReadline(f *os.File) (string, error) {
	var resultBuf []byte
	for {
		var buf [1]byte
		n, err := f.Read(buf[:])
		if err != nil && err != io.EOF {
			return "", err
		}

		if n == 0 || buf[0] == '\n' || buf[0] == '\r' {
			break
		}

		if buf[0] == 3 {
			return "", errors.New(T("go-input.ErrInterrupted"))
		}

		if i.mask {
			_, _ = fmt.Fprint(i.Writer, i.maskVal)
		}

		resultBuf = append(resultBuf, buf[0])
	}

	_, _ = fmt.Fprint(i.Writer, "\n")
	return string(resultBuf), nil
}
