package lazydll

/*
#include <dlfcn.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"unsafe"

	"github.com/0LuigiCode0/CLI/internal/utils/utf"
)

func NewLazyDLL(lib string) *LazyDll {
	str := utf.StrToNum[[]C.char](lib + "\000")

	h := C.dlopen(&str[0], C.RTLD_LAZY)
	if h == nil {
		panic(lastErr())
	}

	return &LazyDll{h: h}
}

type LazyDll struct {
	h unsafe.Pointer
}

func (dll *LazyDll) NewProc(proc string) *Proc {
	str := utf.StrToNum[[]C.char](proc + "\000")
	return &Proc{dll: dll, name: str}
}

type Proc struct {
	dll  *LazyDll
	name []C.char
}

func (p *Proc) Addr() uintptr {
	h := C.dlsym(p.dll.h, &p.name[0])
	if h == nil {
		panic(lastErr())
	}
	return uintptr(h)
}

func lastErr() error {
	data := C.dlerror()
	if data == nil {
		return nil
	}
	return errors.New(C.GoString(data))
}
