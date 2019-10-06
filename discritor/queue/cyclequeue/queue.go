package cyclequeue

import (
	"fmt"

	"github.com/fyniny/datastruct/iface"
)

type state int

const (
	FULL = 1
	EMPTY = -1
	READWRITE = 0
)

// cycle implement a cycle queue
type cycle struct {
	size int
	head int
	tail int
	status int
	init bool
	len  int
	linear iface.LinearList
}

func (c *cycle) Init(list iface.LinearList) {
	if !c.init && list != nil {
		c.init = true
		c.head = 1
		c.tail = 1
		c.len = 0
		c.status = EMPTY
		c.linear = list
	}
}

func (c *cycle) Destroy() {
	if c.init {
		c.init = false
		c.head = 1
		c.tail = 1
		c.len = 0
		c.status = EMPTY
		c.linear = nil
	}
}

func (c *cycle) Clear() {
	if c.init {
		c.linear.Clear()
		c.head = 1
		c.tail = 1
		c.len = 0
		c.status = EMPTY
	}
}

func (c *cycle) Head() interface{} {
	if c.init && c.status != EMPTY {
		return c.linear.GetElem(c.head)
	}

	return nil
}

func (c *cycle) Push(elem interface{}) error {
	if c.init {
		if c.status == FULL {
			return fmt.Errorf("queue is full and reject any other push operation")
		}
		if err := c.linear.Insert(c.tail, elem); err != nil {
			return err
		}
		c.tail = (c.tail+1) % c.size + 1
		if c.tail == c.head {
			c.status = FULL
		} else if c.status == EMPTY {
			c.status = READWRITE
		}
		c.len++
		return nil
	}
	return fmt.Errorf("push into uninitialized queue")
}

func (c *cycle) Delete() interface{} {
	if c.init && c.status != EMPTY {
		ret := c.linear.GetElem(c.head)
		c.len--
		c.head = (c.head+1) % c.size + 1
		if c.head == c.tail {
			c.status = EMPTY
		} else if c.status == FULL {
			c.status = READWRITE
		}
		return ret
	}
	return nil
}

func (c *cycle) Length() int {
	return c.len
}

func (c *cycle) Resize(size int) error {
	if c.init {
		if size <= 0 {
			return fmt.Errorf("queue size must more than one, but size expected to reset is %d", size)
		}
		if size < c.len {
			return fmt.Errorf("current queue length is large than size, therefore resize failed")
		}
		c.size = size
		if c.size == c.len {
			c.status = FULL
		} else if c.status == FULL {
			c.status = READWRITE
		}
		return nil
	}
	return fmt.Errorf("resize uninitialized queue")
}