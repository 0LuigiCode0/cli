package term

import (
	"syscall"
	_ "unsafe"

	"github.com/0LuigiCode0/CLI/internal/c"
)

type iTerminal interface{}

type (
	pws   struct{}
	cmd   struct{}
	xterm struct{}
)

func Detect() iTerminal {
	// fmt.Println(os.Getenv("TERM"), os.Getenv("TERM_PROGRAM"), os.Getenv("COMSPEC"), os.Getenv("SHELL"))
	return pws{}
}

const (
	ENABLE_PROCESSED_INPUT = 0x0001
	ENABLE_WINDOW_INPUT    = 0x0008
	ENABLE_MOUSE_INPUT     = 0x0010
	ENABLE_EXTENDED_FLAGS  = 0x0080
)

var (
	oldConsoleMode uint32
	oldCursor      []byte
)

func Begin() (err error) {
	oldConsoleMode, err = c.GetConsoleMode(syscall.Stdin)
	if err != nil {
		return err
	}
	err = c.SetConsoleMode(syscall.Stdin, ENABLE_MOUSE_INPUT|ENABLE_PROCESSED_INPUT|ENABLE_WINDOW_INPUT|ENABLE_EXTENDED_FLAGS)
	if err != nil {
		return err
	}

	c.WriteAttr(syscall.Stdout, []byte("\033[6n"))
	buf := make([]byte, 16)
	n, err := c.ReadResultAttr(syscall.Stdin, buf)
	if err != nil {
		return err
	}
	oldCursor = buf[:n]
	oldCursor[n-1] = 'H'

	c.WriteAttr(syscall.Stdout, []byte("\033[?1049h"))
	c.WriteAttr(syscall.Stdout, []byte("\033[?25l"))
	c.WriteAttr(syscall.Stdout, []byte("\033[0;0H"))
	// writeAttr(syscall.Stdout, []byte("\033[2J"))
	// writeAttr(syscall.Stdout, []byte("\033[3J"))

	return nil
}

func End() (err error) {
	// writeAttr(syscall.Stdout, []byte("\033[2J"))
	c.WriteAttr(syscall.Stdout, oldCursor)
	c.WriteAttr(syscall.Stdout, []byte("\033[?25h"))
	c.WriteAttr(syscall.Stdout, []byte("\033[?1049l"))

	err = c.SetConsoleMode(syscall.Stdin, oldConsoleMode)
	if err != nil {
		return err
	}
	return
}
