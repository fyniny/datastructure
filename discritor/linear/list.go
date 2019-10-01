package linear

import (
	"fmt"
	"reflect"
)

type node struct {
	elem interface{}
	next *node
}

type list struct {
	len     int
	head    *node
	tail    *node
	init    bool
	compare func(inode1, inode2 interface{}) int
	typ reflect.Type
}

func (l *list) Init(comparer func(inode1, inode2 interface{}) int, typ interface{}) {
	if !l.init {
		l.init = true
		l.head = &node{
			elem: nil,
			next: nil,
		}
		l.tail = l.head
		l.typ = reflect.TypeOf(typ)
		l.len = 0
		l.compare = comparer
	}
}

func (l *list) Empty() bool {
	if l.init && l.len == 0 {
		return true
	}
	return false
}

func (l *list) Clear() {
	l.len = 0
	l.head.next = nil
	l.tail = l.head
}

func (l *list) GetElem(i int) interface{} {
	ret := l.head
	for ret.next != nil && i > 0 {
		i--
		if i == 0 {
			return ret.next.elem
		}
		ret = ret.next
	}
	return nil
}

func (l *list) LocateElem(elem interface{}) int {
	ret := 0
	loop := l.head
	comparer := l.compare
	elemType := reflect.TypeOf(elem).String()
	if comparer == nil {
		// todo: implement elem comparer
		panic("unimplemented comparer function")
	}
	if elemType != l.typ.String() {
		panic("list unmatched type: " + elemType)
	}

	for loop.next != nil {
		ret++
		if comparer(loop.next.elem, elem) == 0 {
			return ret
		}
		loop = loop.next
	}
	return 0
}

func (l *list) Insert(i int, elem interface{}) error {
	if i-1 > l.len || i <= 0 {
		return fmt.Errorf("insert position out of range")
	}

	if l.typ.String() != reflect.TypeOf(elem).String() {
		return fmt.Errorf("list only accept type: %s, receive unmatch type: %s", l.typ.String(), reflect.TypeOf(elem).String())
	}

	loop := l.head
	for loop != nil {
		i--
		if i == 0 {
			tmp := &node{
				elem: elem,
				next: loop.next,
			}
			loop.next = tmp
			break
		}
		loop = loop.next
	}
	l.len++
	return nil
}

func (l *list) Delete(i int) interface{} {
	if l.len <= i || i <= 0 {
		return nil
	}
	j := 0
	loop := l.head
	l.len--
	for loop.next != nil {
		j++
		if i == j {
			tmp := loop.next
			loop.next = loop.next.next
			return tmp.elem
		}
		loop = loop.next
	}
	return nil
}

func (l *list) Length() int {
	return l.len
}
