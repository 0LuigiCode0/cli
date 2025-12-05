package c

import (
	"syscall"
	"unsafe"

	"github.com/0LuigiCode0/CLI/internal/sys"
	"github.com/0LuigiCode0/CLI/internal/utils/union"
)

var (
	kernel32              = syscall.NewLazyDLL("kernel32.dll")
	addrGetLastError      = kernel32.NewProc("GetLastError").Addr()
	addrWriteConsoleW     = kernel32.NewProc("WriteConsoleW").Addr()
	addrWriteConsoleA     = kernel32.NewProc("WriteConsoleA").Addr()
	addrReadConsoleA      = kernel32.NewProc("ReadConsoleA").Addr()
	addrReadConsoleInputW = kernel32.NewProc("ReadConsoleInputW").Addr()
	addrGetConsoleMode    = kernel32.NewProc("GetConsoleMode").Addr()
	addrSetConsoleMode    = kernel32.NewProc("SetConsoleMode").Addr()
)

func lastErr() error { return syscall.Errno(sys.Call[sys.N0](addrGetLastError, true)) }

func WriteFrame(stdout syscall.Handle, frame *uint16, l int) (n uint32, err error) {
	r := sys.Call[sys.N5](addrWriteConsoleW, true, uintptr(stdout), uintptr(unsafe.Pointer(frame)), uintptr(l), uintptr(unsafe.Pointer(&n)), 0)
	if r == 0 {
		err = lastErr()
	}
	return
}

func WriteAttr(stdout syscall.Handle, attr []byte) (err error) {
	r := sys.Call[sys.N5](addrWriteConsoleA, true, uintptr(stdout), uintptr(unsafe.Pointer(&attr[0])), uintptr(len(attr)), 0, 0)
	if r == 0 {
		err = lastErr()
	}
	return
}

func ReadResultAttr(handle syscall.Handle, res []byte) (n uint32, err error) {
	r := sys.Call[sys.N5](addrReadConsoleA, true, uintptr(handle), uintptr(unsafe.Pointer(&res[0])), uintptr(len(res)), uintptr(unsafe.Pointer(&n)), 0)
	if r == 0 {
		err = lastErr()
	}
	return
}

type eventType uint16

const (
	KEY_EVENT                eventType = 0x0001
	MOUSE_EVENT              eventType = 0x0002
	WINDOW_BUFFER_SIZE_EVENT eventType = 0x0004
	MENU_EVENT               eventType = 0x0008
	FOCUS_EVENT              eventType = 0x0010
)

type (
	InputRecord struct {
		EventType eventType
		Event     union.U[union.U16]
	}
	KeyEvent struct {
		IsDown    bool
		Repeat    uint16
		KeyCode   uint16
		ScanCode  uint16
		Char      uint16
		CtrlState uint32
	}
	Coord struct {
		X int16
		Y int16
	}
	MouseEvent struct {
		Coord
		ButtonState uint32
		CtrlState   uint32
		Event       uint32
	}
	SizeEvent struct{ Coord }
)

func ReadConsoleInput(handle syscall.Handle, buf *InputRecord, l int) (n uint32, err error) {
	r := sys.Call[sys.N4](addrReadConsoleInputW, true, uintptr(handle), uintptr(unsafe.Pointer(buf)), uintptr(l), uintptr(unsafe.Pointer(&n)))
	if r == 0 {
		err = lastErr()
	}
	return
}

func GetConsoleMode(handle syscall.Handle) (mode uint32, err error) {
	r := sys.Call[sys.N2](addrGetConsoleMode, true, uintptr(handle), uintptr(unsafe.Pointer(&mode)))
	if r == 0 {
		err = lastErr()
	}
	return
}

func SetConsoleMode(handle syscall.Handle, mode uint32) (err error) {
	r := sys.Call[sys.N2](addrSetConsoleMode, true, uintptr(handle), uintptr(mode))
	if r == 0 {
		err = lastErr()
	}
	return
}
