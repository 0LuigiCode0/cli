package conv

import (
	"fmt"
	"testing"
)

func TestConv(t *testing.T) {
	t.Run("arr", func(t *testing.T) {
		fmt.Println(ArrNum[[]byte]([]int32{1, 0}))
	})
}

func BenchmarkConv(b *testing.B) {
	b.Run("arr", func(b *testing.B) {
		for b.Loop() {
			ArrNum[[]uint16]([]int32{1, 0})
		}
	})
}
