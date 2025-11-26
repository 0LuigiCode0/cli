package cli

import (
	"context"
	"fmt"
	"hash/crc32"
	"os"
	"time"
)

// Event обьект событий
type Event struct{}

// key обьект кнопки
type key string

const (
	KeyEnter     key = "Enter"
	KeySpace     key = "Space"
	KeyBackSpace key = "BackSpace"
	KeyTab       key = "Tab"
	KeyEsc       key = "Esc"
	KeyUp        key = "Up"
	KeyDown      key = "Down"
	KeyLeft      key = "Left"
	KeyRight     key = "Right"
	Key0         key = "0"
	Key1         key = "1"
	Key2         key = "2"
	Key3         key = "3"
	Key4         key = "4"
	Key5         key = "5"
	Key6         key = "6"
	Key7         key = "7"
	Key8         key = "8"
	Key9         key = "9"
	KeyA         key = "a"
	KeyB         key = "b"
	KeyC         key = "c"
	KeyD         key = "d"
	KeyE         key = "e"
	KeyF         key = "f"
	KeyG         key = "g"
	KeyH         key = "h"
	KeyI         key = "i"
	KeyJ         key = "j"
	KeyK         key = "k"
	KeyL         key = "l"
	KeyM         key = "m"
	KeyN         key = "n"
	KeyO         key = "o"
	KeyP         key = "p"
	KeyQ         key = "q"
	KeyR         key = "r"
	KeyS         key = "s"
	KeyT         key = "t"
	KeyU         key = "u"
	KeyV         key = "v"
	KeyW         key = "w"
	KeyX         key = "x"
	KeyY         key = "y"
	KeyZ         key = "z"
	KeyTilda     key = "~"
	KeyPlus      key = "+"
	KeyMinus     key = "-"
	KeyEqual     key = "="
)

var keyList = map[uint32]key{
	hash([3]byte{10, 0, 0}):   KeyEnter,
	hash([3]byte{32, 0, 0}):   KeySpace,
	hash([3]byte{127, 0, 0}):  KeyBackSpace,
	hash([3]byte{9, 0, 0}):    KeyTab,
	hash([3]byte{27, 0, 0}):   KeyEsc,
	hash([3]byte{27, 91, 65}): KeyUp,
	hash([3]byte{27, 91, 66}): KeyDown,
	hash([3]byte{27, 91, 68}): KeyLeft,
	hash([3]byte{27, 91, 67}): KeyRight,
	hash([3]byte{48, 0, 0}):   Key0,
	hash([3]byte{49, 0, 0}):   Key1,
	hash([3]byte{50, 0, 0}):   Key2,
	hash([3]byte{51, 0, 0}):   Key3,
	hash([3]byte{52, 0, 0}):   Key4,
	hash([3]byte{53, 0, 0}):   Key5,
	hash([3]byte{54, 0, 0}):   Key6,
	hash([3]byte{55, 0, 0}):   Key7,
	hash([3]byte{56, 0, 0}):   Key8,
	hash([3]byte{57, 0, 0}):   Key9,
	hash([3]byte{97, 0, 0}):   KeyA,
	hash([3]byte{98, 0, 0}):   KeyB,
	hash([3]byte{99, 0, 0}):   KeyC,
	hash([3]byte{100, 0, 0}):  KeyD,
	hash([3]byte{101, 0, 0}):  KeyE,
	hash([3]byte{102, 0, 0}):  KeyF,
	hash([3]byte{103, 0, 0}):  KeyG,
	hash([3]byte{104, 0, 0}):  KeyH,
	hash([3]byte{105, 0, 0}):  KeyI,
	hash([3]byte{106, 0, 0}):  KeyJ,
	hash([3]byte{107, 0, 0}):  KeyK,
	hash([3]byte{108, 0, 0}):  KeyL,
	hash([3]byte{109, 0, 0}):  KeyM,
	hash([3]byte{110, 0, 0}):  KeyN,
	hash([3]byte{111, 0, 0}):  KeyO,
	hash([3]byte{112, 0, 0}):  KeyP,
	hash([3]byte{113, 0, 0}):  KeyQ,
	hash([3]byte{114, 0, 0}):  KeyR,
	hash([3]byte{115, 0, 0}):  KeyS,
	hash([3]byte{116, 0, 0}):  KeyT,
	hash([3]byte{117, 0, 0}):  KeyU,
	hash([3]byte{118, 0, 0}):  KeyV,
	hash([3]byte{119, 0, 0}):  KeyW,
	hash([3]byte{120, 0, 0}):  KeyX,
	hash([3]byte{121, 0, 0}):  KeyY,
	hash([3]byte{122, 0, 0}):  KeyZ,
	hash([3]byte{96, 0, 0}):   KeyTilda,
	hash([3]byte{43, 0, 0}):   KeyPlus,
	hash([3]byte{45, 0, 0}):   KeyMinus,
	hash([3]byte{61, 0, 0}):   KeyEqual,
}

func (e *Event) listen(ctx context.Context, w IWindow) {
	t := time.NewTicker(fct)
	defer t.Stop()

	mode, err := getConsoleMode(os.Stdin.Fd())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer setConsoleMode(os.Stdin.Fd(), mode)

	setConsoleMode(os.Stdin.Fd(), ENABLE_PROCESSED_INPUT|ENABLE_MOUSE_INPUT)

	enableMouseTracking()
	defer disableMouseTracking()

	buf := make([]byte, 128)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			n, err := os.Stdin.Read(buf)
			if err == nil {
				fmt.Println(buf[:n])
				// if key, ok := keyList[hash(buf)]; ok {
				// 	if l := w.getLayout(); l != nil {
				// 		l.callEvent(key)
				// 	}
				// }
			}
			// os.Stdin.Write([]byte("\033[6n"))
		}
	}
}

func hash(in [3]byte) uint32 {
	h := crc32.NewIEEE()
	h.Write(in[:])
	return h.Sum32()
}
