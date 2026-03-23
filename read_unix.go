//go:build linux || darwin || freebsd

package input

import (
	"errors"
	"os"

	"golang.org/x/term"
)

// LineSep is the separator for windows or unix systems
const LineSep = "\n"

// rawRead reads file with raw mode (without prompting to term).
func (i *UI) rawRead(f *os.File) (string, error) {

	// MakeRaw put the term connected to the given file descriptor
	// into raw mode
	fd := int(f.Fd())
	if !term.IsTerminal(fd) {
		return "", errors.New(T("go-input.read-unix.not-term", fd))
	}

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return "", err
	}
	defer term.Restore(fd, oldState)

	return i.rawReadline(f)
}
