package cli

import (
	"context"
	"os"
	"os/signal"

	"github.com/0LuigiCode0/CLI/internal/term"
)

// App главный обьект приложения
type App interface {
	Start()
	Window() IWindow
}

type app struct {
	w IWindow
	e *Event
}

// InitApp инициализаци приложения
func InitApp(layout ILayout) (App, error) {
	a := &app{}
	// if layout == nil {
	// 	return nil, errors.New("Layout is nil")
	// }
	// w := &window{
	// layout: layout,
	// }
	// a.w = w
	// a.e = &Event{}

	// detectTerm()

	// newLine, newColumn := w.size()
	// frame := make([][]string, newLine, newLine)
	// for i := range frame {
	// 	frame[i] = make([]string, newColumn, newColumn)
	// }
	// w.lines = newLine
	// w.column = newColumn

	return a, nil
}

// Start запуск приложения
func (a *app) Start() {
	ctx, _ := signal.NotifyContext(context.Background(),
		os.Kill, os.Interrupt)
	err := term.Begin()
	defer term.End()

	if err != nil {
		panic(err)
	}

	go a.e.listen(ctx, a.w)

	<-ctx.Done()
}

// GetValue получить обьект окна
func (a *app) Window() IWindow {
	return a.w
}

// func game(App *app) {
// 	close := make(chan os.Signal)
// 	signal.Notify(close, os.Interrupt, os.Kill)

// 	for {
// 		select {
// 		case <-close:
// 			return
// 		default:
// 			i := int(rand.Float32()*100) % App.w.lines
// 			j := int(rand.Float32()*1000) % App.w.column
// 			x := byte(int(rand.Float32()*100)%App.w.lines + 50)
// 			App.w.setPX(i, j, fmt.Sprintf("\033[5m\033[48;5;%vm\033[38;5;%vm%v\033[0m", x, x+20, string(x)))

// 			time.Sleep(fct)
// 		}
// 	}
// }
