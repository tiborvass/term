// +build windows

package windowsconsole

import (
	"fmt"
	"io"
	"os"

	ansiterm "github.com/Azure/go-ansiterm"
	"github.com/Azure/go-ansiterm/winterm"
	"golang.org/x/sys/windows"
)

// ansiWriter wraps a standard output file (e.g., os.Stdout) providing ANSI sequence translation.
type ansiWriter struct {
	file           *os.File
	fd             uintptr
	infoReset      *winterm.CONSOLE_SCREEN_BUFFER_INFO
	command        []byte
	escapeSequence []byte
	inAnsiSequence bool
	parser         *ansiterm.AnsiParser
}

// NewAnsiWriter returns an io.Writer that provides VT100 terminal emulation on top of a
// Windows console output handle.
func NewAnsiWriter(nFile uint32) io.Writer {
	fd, err := windows.GetStdHandle(nFile)
	file, ok := handleToFile[nFile]
	if err != nil || !ok {
		panic(fmt.Errorf("Invalid standard handle: %v (file=%v)", nFile, file))
	}
	info, err := winterm.GetConsoleScreenBufferInfo(fd)
	if err != nil {
		return nil
	}

	parser := ansiterm.CreateParser("Ground", winterm.CreateWinEventHandler(fd, file))

	return &ansiWriter{
		file:           file,
		fd:             fd,
		infoReset:      info,
		command:        make([]byte, 0, ansiterm.ANSI_MAX_CMD_LENGTH),
		escapeSequence: []byte(ansiterm.KEY_ESC_CSI),
		parser:         parser,
	}
}

func (aw *ansiWriter) Fd() uintptr {
	return aw.fd
}

// Write writes len(p) bytes from p to the underlying data stream.
func (aw *ansiWriter) Write(p []byte) (total int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	return aw.parser.Parse(p)
}
