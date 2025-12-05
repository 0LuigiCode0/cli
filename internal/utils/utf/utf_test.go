package utf

import (
	"fmt"
	"testing"
	"unicode/utf16"
	"unsafe"
)

func TestConv(t *testing.T) {
	s := "hello world привет мир 우리 반 친구들은 공부도 열심히 하고  𔓙𔓙𔓙𔓙\000"
	sr := []rune(s)
	t.Run("str num utf16", func(t *testing.T) {
		fmt.Println(StrToNum[[]uint16](s))
		fmt.Println(utf16.Encode([]rune(s)))
	})
	t.Run("str num utf32", func(t *testing.T) {
		fmt.Println(StrToNum[[]uint32](s))
		// fmt.Println(utf32.([]rune(s)))
	})
	t.Run("num str utf16", func(t *testing.T) {
		fmt.Println(NumToStr([]uint16{'g', 'g', 'п', '리', 0xD811, 0xDCD9}))
	})
	t.Run("num str utf32", func(t *testing.T) {
		fmt.Println(NumToStr([]rune(s)))
	})
	t.Run("str ptr", func(t *testing.T) {
		fmt.Println(*StrToPtr[uint32](s))
	})
	t.Run("ptr str", func(t *testing.T) {
		fmt.Println(PtrToStr[rune](unsafe.Pointer(&sr[0])))
	})
}

func BenchmarkConv(b *testing.B) {
	s := "hello world привет мир 우리 반 친구들은 공부도 열심히 하고  𔓙𔓙𔓙𔓙\000"
	// sp := unsafe.Pointer(unsafe.StringData(s))
	sr := []rune(s)
	sl := utf16.Encode(sr)
	// sb := []byte(s)

	b.Run("str unf16", func(b *testing.B) {
		for b.Loop() {
			utf16.Encode([]rune(s))
		}
	})
	b.Run("str num", func(b *testing.B) {
		for b.Loop() {
			StrToNum[[]uint16](s)
		}
	})
	b.Run("unf16 str", func(b *testing.B) {
		for b.Loop() {
			_ = string(utf16.Decode(sl))
		}
	})
	b.Run("num str", func(b *testing.B) {
		for b.Loop() {
			NumToStr(sl)
		}
	})
	b.Run("str ptr", func(b *testing.B) {
		for b.Loop() {
			StrToPtr[uint16](s)
		}
	})
	b.Run("ptr str", func(b *testing.B) {
		for b.Loop() {
			PtrToStr[rune](unsafe.Pointer(&sr[0]))
		}
	})
}
