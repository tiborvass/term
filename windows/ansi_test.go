// +build windows

package windowsconsole

import (
	"os"
	"syscall"
	"testing"

	"golang.org/x/sys/windows"
)

func TestAnsiReaderStdinHandle(t *testing.T) {
	rc1 := NewAnsiReader(windows.STD_INPUT_HANDLE)
	ar1 := rc1.(*ansiReader)
	rc2 := NewAnsiReader(syscall.STD_INPUT_HANDLE)
	ar2 := rc2.(*ansiReader)
	if ar1.file != os.Stdin {
		t.Fatalf("ar1.file should be os.Stdin")
	}
	if ar2.file != os.Stdin {
		t.Fatalf("ar2.file should be os.Stdin")
	}
	if ar1.fd != windows.Stdin {
		t.Fatalf("ar1.fd should be windows.Stdin")
	}
	if ar2.fd != windows.Stdin {
		t.Fatalf("ar2.fd should be windows.Stdin")
	}
}

func TestAnsiWriterStdoutHandle(t *testing.T) {
	w1 := NewAnsiWriter(windows.STD_OUTPUT_HANDLE)
	aw1 := w1.(*ansiWriter)
	w2 := NewAnsiWriter(syscall.STD_OUTPUT_HANDLE)
	aw2 := w2.(*ansiWriter)
	if aw1.file != os.Stdout {
		t.Fatalf("aw1.file should be os.Stdout")
	}
	if aw2.file != os.Stdout {
		t.Fatalf("aw2.file should be os.Stdout")
	}
	if aw1.fd != windows.Stdout {
		t.Fatalf("aw1.fd should be windows.Stdout")
	}
	if aw2.fd != windows.Stdout {
		t.Fatalf("aw2.fd should be windows.Stdout")
	}
}
