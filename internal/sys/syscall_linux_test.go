package sys

import (
	"math"
	"os"
	"syscall"
	"testing"
	"unsafe"

	"github.com/0LuigiCode0/CLI/internal/utils/lazydll"
	"github.com/0LuigiCode0/CLI/internal/utils/utf"
	"github.com/stretchr/testify/assert"
)

func TestSys(t *testing.T) {
	libc := lazydll.NewLazyDLL("libc.so.6")
	addrGetpid := libc.NewProc("getpid").Addr()
	addrGetEnvironmentVariable := libc.NewProc("getenv").Addr()
	// addrGetProcessTimes := libc.NewProc("GetProcessTimes").Addr()

	libm := lazydll.NewLazyDLL("libm.so.6")
	addrPow := libm.NewProc("jn").Addr()

	t.Run("sys0", func(t *testing.T) {
		pid := syscall.Getpid()
		assert.Equal(t, pid, int(Call[N0](addrGetpid, IsC)))
	})
	t.Run("sys3", func(t *testing.T) {
		name := "X"

		os.Setenv("X", "hello")
		env1, _ := syscall.Getenv("X")
		char := Call[N1](addrGetEnvironmentVariable, IsC,
			uintptr(unsafe.Pointer(utf.StrToPtr[byte](name))))
		assert.Equal(t, env1, utf.PtrToStr[byte](unsafe.Pointer(char)))
	})
	t.Run("sys3 float", func(t *testing.T) {
		f1, f2 := 2, 2.23
		x1 := call3(addrPow, IsC|F2|FOut, uintptr(f1), uintptr(math.Float64bits(f2)), 0)
		x3 := math.Jn(f1, f2)
		assert.Equal(t, x3, math.Float64frombits(uint64(x1)))
	})

	// t.Run("sys6", func(t *testing.T) {
	// 	handle, _ := syscall.GetCurrentProcess()
	// 	var creationTime1, exitTime1, kernelTime1, userTime1 syscall.Filetime
	// 	var creationTime2, exitTime2, kernelTime2, userTime2 syscall.Filetime
	// 	err := syscall.GetProcessTimes(handle, &creationTime1, &exitTime1, &kernelTime1, &userTime1)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	call6(addrGetProcessTimes, ISC,
	// 		uintptr(handle),
	// 		uintptr(unsafe.Pointer(&creationTime2)),
	// 		uintptr(unsafe.Pointer(&exitTime2)),
	// 		uintptr(unsafe.Pointer(&kernelTime2)),
	// 		uintptr(unsafe.Pointer(&userTime2)), 0)
	// 	assert.Equal(t, creationTime1, creationTime2)
	// 	assert.Equal(t, exitTime1, exitTime2)
	// 	assert.Equal(t, kernelTime1, kernelTime2)
	// 	assert.Equal(t, userTime1, userTime2)
	// })
}

func BenchmarkPow(b *testing.B) {
	f1, f2 := 2.8, 2.23
	math32 := lazydll.NewLazyDLL("libm.so.6")
	addrPow := math32.NewProc("pow").Addr()

	b.Run("native", func(b *testing.B) {
		for b.Loop() {
			_ = math.Pow(f1, f2)
		}
	})
	b.Run("asm", func(b *testing.B) {
		for b.Loop() {
			_ = call3(addrPow, IsC|F1|F2|FOut, uintptr(math.Float64bits(f1)), uintptr(math.Float64bits(f2)), 0)
		}
	})
	b.Run("cgo", func(b *testing.B) {
		for b.Loop() {
			_, _, _ = syscall.RawSyscall(addrPow, uintptr(math.Float64bits(f1)), uintptr(math.Float64bits(f2)), 0)
		}
	})
}
