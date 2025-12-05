package cli

import (
	"context"
)

// IWindow интерфейс окна
type IWindow interface {
	reSize(ctx context.Context)
	reView(ctx context.Context)
	getLayout() ILayout
}

type window struct {
	lines  int
	column int
	frame  [][]byte
	layout ILayout
}

func (w *window) reSize(ctx context.Context) {
}

func (w *window) reView(ctx context.Context) {
}

func (w *window) getLayout() ILayout { return w.layout }
