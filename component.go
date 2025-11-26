package cli

// IComponent итерфейс компонента
type IComponent interface {
	SetOnCreate(f func()) IComponent
	SetOnUpdate(f func()) IComponent
	SetOnDelete(f func()) IComponent
	SetOnSelect(f func()) IComponent
	SetOnActive(f func()) IComponent
	SetOnBack(f func()) IComponent
	SetOnUp(f func()) IComponent
	SetOnDown(f func()) IComponent
	SetOnLeft(f func()) IComponent
	SetOnRight(f func()) IComponent

	callOnCreate()
	callOnUpdate()
	callOnDelete()
	callOnTab()
	callOnEnter()
	callOnBackSpace()
	callOnUp()
	callOnDown()
	callOnLeft()
	callOnRight()

	SetComponents(comp ...IComponent) IComponent
	SetStyle(style map[Style]interface{}) IComponent
	SetActive(active bool) IComponent
	Style() map[Style]interface{}
	Active() bool
}

type component struct {
	active      bool
	components  []IComponent
	style       map[Style]interface{}
	onCreate    func()
	onUpdate    func()
	onDelete    func()
	onTab       func()
	onEnter     func()
	onBackSpace func()
	onUp        func()
	onDown      func()
	onLeft      func()
	onRight     func()
}

func (c *component) SetOnCreate(f func()) IComponent { c.onCreate = f; return c }
func (c *component) SetOnUpdate(f func()) IComponent { c.onUpdate = f; return c }
func (c *component) SetOnDelete(f func()) IComponent { c.onDelete = f; return c }
func (c *component) SetOnSelect(f func()) IComponent { c.onTab = f; return c }
func (c *component) SetOnActive(f func()) IComponent { c.onEnter = f; return c }
func (c *component) SetOnBack(f func()) IComponent   { c.onBackSpace = f; return c }
func (c *component) SetOnUp(f func()) IComponent     { c.onUp = f; return c }
func (c *component) SetOnDown(f func()) IComponent   { c.onDown = f; return c }
func (c *component) SetOnLeft(f func()) IComponent   { c.onLeft = f; return c }
func (c *component) SetOnRight(f func()) IComponent  { c.onRight = f; return c }

func (c *component) callOnCreate()    { c.onCreate() }
func (c *component) callOnUpdate()    { c.onUpdate() }
func (c *component) callOnDelete()    { c.onDelete() }
func (c *component) callOnTab()       { c.onTab() }
func (c *component) callOnEnter()     { c.onEnter() }
func (c *component) callOnBackSpace() { c.onBackSpace() }
func (c *component) callOnUp()        { c.onUp() }
func (c *component) callOnDown()      { c.onDown() }
func (c *component) callOnLeft()      { c.onLeft() }
func (c *component) callOnRight()     { c.onRight() }

func (c *component) SetComponents(comp ...IComponent) IComponent {
	c.components = comp
	if f := c.onUpdate; f != nil {
		f()
	}
	return c
}

func (c *component) SetStyle(style map[Style]interface{}) IComponent {
	c.style = style
	if f := c.onUpdate; f != nil {
		f()
	}
	return c
}

func (c *component) SetActive(active bool) IComponent { c.active = active; return c }

func (c *component) Style() map[Style]interface{} { return c.style }
func (c *component) Active() bool                 { return c.active }
