package cli

// ILayout интерфейс слоя
type ILayout interface {
	SetEvent(key key, f func()) ILayout
	SetOnCreate(f func()) ILayout
	SetOnUpdate(f func()) ILayout
	SetOnDelete(f func()) ILayout

	callEvent(key key)
	callOnCreate()
	callOnUpdate()
	callOnDelete()

	SetStyle(style map[Style]interface{}) ILayout
	SetComponents(comp ...IComponent) ILayout
}

type layout struct {
	event      map[key]func()
	components []IComponent
	style      map[Style]interface{}
	onCreate   func()
	onUpdate   func()
	onDelete   func()
}

// Layout создание нового слоя
func Layout() ILayout {
	l := &layout{
		event: map[key]func(){},
	}

	return l
}

func (l *layout) SetEvent(key key, f func()) ILayout { l.event[key] = f; return l }
func (l *layout) SetOnCreate(f func()) ILayout       { l.onCreate = f; return l }
func (l *layout) SetOnUpdate(f func()) ILayout       { l.onUpdate = f; return l }
func (l *layout) SetOnDelete(f func()) ILayout       { l.onDelete = f; return l }

func (l *layout) callEvent(key key) {
	if e, ok := l.event[key]; ok {
		e()
	}
}
func (l *layout) callOnCreate() { l.onCreate() }
func (l *layout) callOnUpdate() { l.onUpdate() }
func (l *layout) callOnDelete() { l.onDelete() }

func (l *layout) SetStyle(style map[Style]interface{}) ILayout {
	l.style = style
	if f := l.onUpdate; f != nil {
		f()
	}
	return l
}

func (l *layout) SetComponents(comp ...IComponent) ILayout {
	l.components = comp
	if f := l.onUpdate; f != nil {
		f()
	}
	return l
}
