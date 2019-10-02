package queue

import (
	"fmt"
	"github.com/fyniny/datastruct/iface"
)

type Queue iface.Queue

type queueLinear struct {
	linear iface.LinearList
	init bool
}

func (ql *queueLinear) Init(linear iface.LinearList) {
	if !ql.init && linear != nil {
		ql.linear = linear
		ql.init = true
	}
}

func (ql *queueLinear) Destroy() {
	if ql.init {
		ql.init = false
		ql.linear = nil
	}
}

func (ql *queueLinear) Clear() {
	if ql.linear != nil {
		ql.linear.Clear()
	}
}

func (ql *queueLinear) Head() interface{} {
	if ql.init {
		return ql.linear.GetElem(1)
	}
	return nil
}

func (ql *queueLinear) Push(elem interface{}) error {
	if ql.init {
		return ql.linear.Insert(ql.linear.Length()+1, elem)
	}
	return fmt.Errorf("push into uninitialized queue")
}

func (ql *queueLinear) Delete() interface{} {
	if ql.init {
		return ql.linear.GetElem(1)
	}
	return nil
}

func (ql *queueLinear) Length() int {
	if ql.init {
		return ql.linear.Length()
	}
	return 0
}