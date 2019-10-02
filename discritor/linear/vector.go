package linear

import (
	"fmt"
	"reflect"
)

const minimum = 20

type vector struct {
	len     int
	elem 	[]interface{}
	init    bool
	compare func(inode1, inode2 interface{}) int
	typ     reflect.Type
	cap		int
}

func (v *vector) Init(comparer func(inode1, inode2 interface{}) int, typ interface{}) {
	if !v.init {
		v.init = true
		v.typ = reflect.TypeOf(typ)
		v.len = 0
		v.elem=make([]interface{}, minimum)
		v.compare = comparer
		v.cap = minimum
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

	for i := 0 ; i < v.Length(); i++ {
		if comparer(v.elem[i], elem) == 0 {
			return i+1
		}
	}
	return 0
}

func (v *vector) Insert(i int, elem interface{}) error {
	if i-1 > v.len || i <= 0 {
		return fmt.Errorf("insert position out of range")
	}

	if v.typ.String() != reflect.TypeOf(elem).String() {
		return fmt.Errorf("list only accept type: %s, receive unmatch type: %s", v.typ.String(), reflect.TypeOf(elem).String())
	}

	defer func() {v.len++}()
	// expand v.elem
	if v.len >= v.cap {
		mul := 2.0
		if v.cap > 1024 {
			mul = 1.5
		}
		v.cap = int(float64(v.cap) * mul)
		news := make([]interface{}, v.cap)
		index := 1
		for index < i {
			news[index-1] = v.elem[index-1]
			index++
		}

		news[index-1] = elem
		index++

		for i < len(v.elem) {
			news[index-1] = v.elem[i]
			i++
			index++
		}
		v.elem = news
		return nil
	}

	for l := v.len; l >= i; l-- {
		v.elem[l] = v.elem[l-1]
	}

	v.elem[i-1] = elem
	return nil
}

func (v *vector) Delete(i int) interface{} {
	if v.len <= i || i <= 0 {
		return nil
	}

	ret := v.elem[i-1]
	for i < v.len {
		v.elem[i-1] = v.elem[i]
		i++
	}
	v.len--
	return ret
}

func (v *vector) Length() int {
	return v.len
}
