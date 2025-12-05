package utf

import "unsafe"

type uft interface {
	~int | ~int8 | ~int16 | ~int32 |
		~uint | ~uint8 | ~uint16 | ~uint32
}

const (
	_UTF8  = 1
	_UTF16 = 2
	_UTF32 = 4
)

func StrToNum[to ~[]t, t uft](in string) to {
	var _t t
	lIn := len(in)
	if lIn == 0 {
		return nil
	}
	SIZE := int(unsafe.Sizeof(_t))

	if SIZE == _UTF8 {
		return unsafe.Slice((*t)(unsafe.Pointer(unsafe.StringData(in))), lIn)
	}

	out := make(to, lIn)
	ptrIn := unsafe.Pointer(unsafe.StringData(in))
	ptrOut := unsafe.Pointer(&out[0])

	var l int
	var shift int
loop:
	for i := 0; l < lIn; l++ {
		if i >= lIn {
			break loop
		}
		ptrIn = unsafe.Add(ptrIn, shift)

		c := *(*byte)(ptrIn)
		switch {
		case c&0x80 == 0:
			*(*t)(ptrOut) = t(c)
			shift = 1
		case (c>>5)^0x06 == 0:
			if i+1 > lIn {
				break loop
			}
			*(*t)(ptrOut) = t((uint16(c&0x1f) << 6) | uint16(*(*byte)(unsafe.Add(ptrIn, 1))&0x3f))
			shift = 2
		case (c>>4)^0x0e == 0:
			if i+2 > lIn {
				break loop
			}
			*(*t)(ptrOut) = t((uint16(c&0x0f) << 12) | (uint16(*(*byte)(unsafe.Add(ptrIn, 1))&0x3f) << 6) | uint16(*(*byte)(unsafe.Add(ptrIn, 2))&0x3f))
			shift = 3
		default:
			if i+3 > lIn {
				break loop
			}
			r := (uint32(c&0x07) << 18) | (uint32(*(*byte)(unsafe.Add(ptrIn, 1))&0x3f) << 12) | (uint32(*(*byte)(unsafe.Add(ptrIn, 2))&0x3f) << 6) | uint32(*(*byte)(unsafe.Add(ptrIn, 3))&0x3f)
			switch SIZE {
			case _UTF16:
				if l+1 > lIn {
					break loop
				}
				r &= 0xfeffff
				*(*t)(ptrOut) = t(uint16(0xd800) | uint16(r>>10))
				ptrOut = unsafe.Add(ptrOut, SIZE)
				*(*t)(ptrOut) = t(uint16(0xdc00) | uint16(r&0x03ff))
				l++
			default:
				*(*t)(ptrOut) = t(r)
			}
			shift = 4
		}
		ptrOut = unsafe.Add(ptrOut, SIZE)
		i += shift
	}
	return out[:l:l]
}

func NumToStr[from ~[]t, t uft](in from) string {
	var _t t
	lIn := len(in)
	if lIn == 0 {
		return ""
	}
	SIZE := int(unsafe.Sizeof(_t))

	if SIZE == 1 {
		return unsafe.String((*byte)(unsafe.Pointer(&in[0])), lIn)
	}

	lOut := lIn * SIZE
	out := make([]byte, lOut)

	ptrIn := unsafe.Pointer(&in[0])
	ptrOut := unsafe.Pointer(&out[0])

	var l int
	var shift int
loop:
	for i := 0; i < lIn; i++ {
		if l >= lOut {
			break loop
		}
		ptrOut = unsafe.Add(ptrOut, shift)

		c := *(*t)(ptrIn)
		switch {
		case (uint16(c)>>8)&0xff == 0:
			*(*byte)(ptrOut) = byte(c) & 0x7f
			shift = 1
		case (uint16(c)>>11)^0x1b == 0:
			if l+3 > lOut || i+1 > lIn {
				break loop
			}
			r := (uint32(uint16(c)&0x07ff)<<10 | uint32(uint16(*(*t)(unsafe.Add(ptrIn, SIZE)))&0x03ff) | 0x10000)
			*(*byte)(ptrOut) = 0xf0 | (byte(r>>18) & 0x07)
			*(*byte)(unsafe.Add(ptrOut, 1)) = 0x80 | (byte(r>>12) & 0x3f)
			*(*byte)(unsafe.Add(ptrOut, 2)) = 0x80 | (byte(r>>6) & 0x3f)
			*(*byte)(unsafe.Add(ptrOut, 3)) = 0x80 | (byte(r) & 0x3f)
			ptrIn = unsafe.Add(ptrIn, SIZE)
			i++
			shift = 4
		case (uint16(c)>>8)&0xf8 == 0:
			if l+1 > lOut {
				break loop
			}
			*(*byte)(ptrOut) = 0xc0 | (byte(uint16(c)>>6) & 0x1f)
			*(*byte)(unsafe.Add(ptrOut, 1)) = 0x80 | (byte(c) & 0x3f)
			shift = 2
		case (uint32(c)>>16)&0xff == 0:
			if l+2 > lOut {
				break loop
			}
			*(*byte)(ptrOut) = 0xe0 | (byte(uint16(c)>>12) & 0x0f)
			*(*byte)(unsafe.Add(ptrOut, 1)) = 0x80 | (byte(uint16(c)>>6) & 0x3f)
			*(*byte)(unsafe.Add(ptrOut, 2)) = 0x80 | (byte(c) & 0x3f)
			shift = 3
		default:
			if SIZE == _UTF32 {
				if l+3 > lOut {
					break loop
				}
				*(*byte)(ptrOut) = 0xf0 | (byte(uint32(c)>>18) & 0x07)
				*(*byte)(unsafe.Add(ptrOut, 1)) = 0x80 | (byte(uint32(c)>>12) & 0x3f)
				*(*byte)(unsafe.Add(ptrOut, 2)) = 0x80 | (byte(uint16(c)>>6) & 0x3f)
				*(*byte)(unsafe.Add(ptrOut, 3)) = 0x80 | (byte(c) & 0x3f)
				shift = 4
			}
		}
		ptrIn = unsafe.Add(ptrIn, SIZE)
		l += shift
	}
	return unsafe.String(&out[0], l)
}

func PtrToStr[t uft](ptrIn unsafe.Pointer) string {
	if ptrIn == nil {
		return ""
	}
	startPtr := ptrIn
	var lOut int
	var c t
	SIZE := int(unsafe.Sizeof(c))

	for *(*t)(ptrIn) != 0 {
		lOut += SIZE
		ptrIn = unsafe.Add(ptrIn, SIZE)
	}

	ptrIn = startPtr
	if SIZE == 1 {
		return unsafe.String((*byte)(ptrIn), lOut)
	}

	out := make([]byte, lOut)
	ptrOut := unsafe.Pointer(&out[0])

	var l int
	var shift int
loop:
	for i := 0; ; i++ {
		if l >= lOut {
			break loop
		}
		ptrOut = unsafe.Add(ptrOut, shift)
		c := *(*t)(ptrIn)
		if c == 0 {
			break loop
		}

		switch {
		case (uint16(c)>>8)&0xff == 0:
			*(*byte)(ptrOut) = byte(c) & 0x7f
			shift = 1
		case (uint16(c)>>11)^0x1b == 0:
			if l+3 > lOut || *(*t)(unsafe.Add(ptrIn, SIZE)) == 0 {
				break loop
			}
			r := (uint32(uint16(c)&0x07ff)<<10 | uint32(uint16(*(*t)(unsafe.Add(ptrIn, SIZE)))&0x03ff) | 0x10000)
			*(*byte)(ptrOut) = 0xf0 | (byte(r>>18) & 0x07)
			*(*byte)(unsafe.Add(ptrOut, 1)) = 0x80 | (byte(r>>12) & 0x3f)
			*(*byte)(unsafe.Add(ptrOut, 2)) = 0x80 | (byte(r>>6) & 0x3f)
			*(*byte)(unsafe.Add(ptrOut, 3)) = 0x80 | (byte(r) & 0x3f)
			ptrIn = unsafe.Add(ptrIn, SIZE)
			i++
			shift = 4
		case (uint16(c)>>8)&0xf8 == 0:
			if l+1 > lOut {
				break loop
			}
			*(*byte)(ptrOut) = 0xc0 | (byte(uint16(c)>>6) & 0x1f)
			*(*byte)(unsafe.Add(ptrOut, 1)) = 0x80 | (byte(c) & 0x3f)
			shift = 2
		case (uint32(c)>>16)&0xff == 0:
			if l+2 > lOut {
				break loop
			}
			*(*byte)(ptrOut) = 0xe0 | (byte(uint16(c)>>12) & 0x0f)
			*(*byte)(unsafe.Add(ptrOut, 1)) = 0x80 | (byte(uint16(c)>>6) & 0x3f)
			*(*byte)(unsafe.Add(ptrOut, 2)) = 0x80 | (byte(c) & 0x3f)
			shift = 3
		default:
			if SIZE == _UTF32 {
				if l+3 > lOut {
					break loop
				}
				*(*byte)(ptrOut) = 0xf0 | (byte(uint32(c)>>18) & 0x07)
				*(*byte)(unsafe.Add(ptrOut, 1)) = 0x80 | (byte(uint32(c)>>12) & 0x3f)
				*(*byte)(unsafe.Add(ptrOut, 2)) = 0x80 | (byte(uint16(c)>>6) & 0x3f)
				*(*byte)(unsafe.Add(ptrOut, 3)) = 0x80 | (byte(c) & 0x3f)
				shift = 4
			}
		}
		ptrIn = unsafe.Add(ptrIn, SIZE)
		l += shift
	}
	return unsafe.String(&out[0], l)
}

//go:nosplit
func StrToPtr[tChar uft](in string) *tChar { return &StrToNum[[]tChar](in)[0] }
