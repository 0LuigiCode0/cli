package cli

import (
	"os"
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	b.Run("syscoll", func(b *testing.B) {
		for b.Loop() {
			os.Stdin.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		}
	})
	b.Run("asm", func(b *testing.B) {
		for b.Loop() {
			fastWrite(os.Stdin.Fd(), []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		}
	})
}

func TestFastWrite(t *testing.T) {
	fastWrite(os.Stdin.Fd(), []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
