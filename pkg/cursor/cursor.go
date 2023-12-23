package cursor

type Cursor[T any] struct {
	index int
	items []T
}

func NewCursor[T any](items ...T) *Cursor[T] {
	return &Cursor[T]{index: 0, items: items}
}

func (c *Cursor[T]) Inc() T {
	c.index++
	if c.index == len(c.items) {
		c.index = 0
	}

	return c.items[c.index]
}

func (c *Cursor[T]) Dec() T {
	c.index--
	if c.index < 0 {
		c.index = len(c.items) - 1
	}

	return c.items[c.index]
}

func (c *Cursor[T]) CurrentValue() T {
	return c.items[c.index]
}

func (c *Cursor[T]) AllValues() []T {
	return c.items
}

func (c *Cursor[T]) Index() int {
	return c.index
}
