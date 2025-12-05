package lazydll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyscall(t *testing.T) {
	dll := NewLazyDLL("libc.so.6")
	f := dll.NewProc("clock_gettime").Addr()

	assert.NotEqual(t, 0, f)
}
