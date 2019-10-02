package stack

import "github.com/fyniny/datastruct/iface"

type Stack iface.Stack

type stackLinear struct {
	list iface.LinearList
	init bool
}

func (sl *stackLinear) Init(list iface.LinearList) {
	if list != nil {
		sl.list = list
		sl.init = true
	}
}

func (sl *stackLinear) Destroy() {
	if sl.list != nil {
		sl.list = nil
		sl.init = false
	}
}

func (sl *stackLinear) Empty () bool {
	if sl.init {
		return sl.list.Empty()
	}

	return false
}

func (sl *stackLinear) Top() interface{} {
	if sl.Empty() && !sl.init {
		return nil
	}

	return sl.list.GetElem(1)
}

func (sl *stackLinear) Push(elem interface{}) {
	if !sl.init {
		panic("stack uninitialized")
	}

	err := sl.list.Insert(sl.list.Length()+1, elem)
	if err != nil {
		panic(err)
	}
}

func (sl *stackLinear) Pop() interface{} {
	if sl.Empty() && !sl.init {
		panic("stack uninitialized or empty")
	}

	return sl.list.Delete(sl.list.Length())
}

func (sl *stackLinear) Length() int {
	if sl.init {
		return sl.list.Length()
	}
	return 0
}

func (sl *stackLinear) Clear() {
	if sl.init {
		sl.list.Clear()
	}
}