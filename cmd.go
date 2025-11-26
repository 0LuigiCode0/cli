package cli

import (
	"os"
	"syscall"
	"unsafe"
	_ "unsafe"
)

type iTerminal interface{}

type (
	pws   struct{}
	cmd   struct{}
	xterm struct{}
)

func detectTerm() iTerminal {
	// fmt.Println(os.Getenv("TERM"), os.Getenv("TERM_PROGRAM"), os.Getenv("COMSPEC"), os.Getenv("SHELL"))
	return pws{}
}

func enableMouseTracking() {
	// Включить расширенный режим мыши (SGR)
	os.Stdin.WriteString("\033[?9h")
	// Включить отслеживание кликов и перетаскивания
	os.Stdin.WriteString("\033[?1049h")
	os.Stdin.WriteString("\033[?25l")
}

func disableMouseTracking() {
	os.Stdin.WriteString("\033[?9h")
	os.Stdin.WriteString("\033[?1049l")
	os.Stdin.WriteString("\033[?25h") // Показать курсор
}

const (
	ENABLE_PROCESSED_INPUT = 0x0001
	ENABLE_LINE_INPUT      = 0x0002
	ENABLE_ECHO_INPUT      = 0x0004
	ENABLE_WINDOW_INPUT    = 0x0008
	ENABLE_MOUSE_INPUT     = 0x0010
)

var oldstd = os.Stdout

func termClear() {
	os.Stdout.Write([]byte("\033[0;0H"))
	os.Stdout.Write([]byte("\033[2J"))
}

func termReset() {
	os.Stdout.Write([]byte("\033[0;0H"))
	os.Stdout.Write([]byte("\033[2J"))
}

var (
	modkernel32        = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode = modkernel32.NewProc("GetConsoleMode")
	procSetConsoleMode = modkernel32.NewProc("SetConsoleMode")
)

func getConsoleMode(fd uintptr) (mode uint32, err error) {
	r, _, errs := syscall.SyscallN(procGetConsoleMode.Addr(), fd, uintptr(unsafe.Pointer(&mode)))
	if r == 0 {
		err = errs
	}
	return
}

func setConsoleMode(fd uintptr, mode uint32) (err error) {
	r, _, errs := syscall.SyscallN(procSetConsoleMode.Addr(), fd, uintptr(mode))
	if r == 0 {
		err = errs
	}
	return
}

type ioStatus struct {
	status uintptr
	info   uint64
}

func fastWrite(fd uintptr, buf []byte) (int, error) {
	var res ioStatus
	write(fd, uintptr(unsafe.Pointer(&res)), uintptr(unsafe.Pointer((*byte)(unsafe.SliceData(buf)))), uint64(len(buf)))
	return int(res.info), nil
}

func write(fd uintptr, ioStatus uintptr, buf uintptr, len uint64)
