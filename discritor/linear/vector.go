package linear

import (
	"fmt"
	"reflect"
)


type vector struct {
	len     int
	elem 	[]interface{}
	init    bool
	compare func(inode1, inode2 interface{}) int
	typ     reflect.Type
}

func (v *vector) Init(comparer func(inode1, inode2 interface{}) int, typ interface{}) {
	if !v.init {
		v.init = true
		v.typ = reflect.TypeOf(typ)
		v.len = 0
		v.compare = comparer
	}
}

func (v *vector) Empty() bool {
	if v.init && v.len == 0 {
		return true
	}
	return false
}

func (v *vector) Clear() {
	v.len = 0
	v.elem = nil
}

func (v *vector) GetElem(i int) interface{} {
	if i <= v.len && i > 0 {
		return v.elem[i-1]
	}
	return nil
}

func (v *vector) LocateElem(elem interface{}) int {
	comparer := v.compare
	elemType := reflect.TypeOf(elem).String()
	if comparer == nil {
		// todo: implement elem comparer
		panic("unimplemented comparer function")
	}
	if elemType != v.typ.String() {
		panic("list unmatched type: " + elemType)
	}

	for i, item := range v.elem {
		if comparer(item, elem) == 0 {
			return i+1
		}
	}
	return 0
}

func (v *vector) Insert(i int, elem interface{}) error {
	if i > v.len || i <= 0 {
		return fmt.Errorf("insert position out of range")
	}

	if v.typ.String() != reflect.TypeOf(elem).String() {
		return fmt.Errorf("list only accept type: %s, receive unmatch type: %s", v.typ.String(), reflect.TypeOf(elem).String())
	}

	v.elem = append(v.elem[0: i], v.elem[i-1:]...)
	v.elem[i-1]=elem
	v.len++
	return nil
}

func (v *vector) Delete(i int) interface{} {
	if v.len <= i || i <= 0 {
		return nil
	}

	ret := v.elem[i-1]
	v.elem = append(v.elem[0: i-1], v.elem[i:]...)
	v.len--
	return ret
}

func (v *vector) Length() int {
	return v.len
}
