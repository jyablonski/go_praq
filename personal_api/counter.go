package main

type Counter struct {
	value int
}

// have to creat this via c := &Counter{2}
// This means that Inc() (which is a method with a pointer receiver, i.e., func (c *Counter)
// Inc()) won’t modify the original Counter object because it's working with a copy of the
// struct, not the original struct.
func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
