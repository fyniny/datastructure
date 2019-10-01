// factor export linear list/vector
package linear

import "github.com/fyniny/datastruct/iface"

// List returns linear with list engine
func List (comparer func(inode1, inode2 interface{}) int, typ interface{}) iface.LinearList {
	l := &list{}
	l.Init(comparer, typ)
	return l
}

// Vector returns linear with vector engine
func Vector (comparer func(inode1, inode2 interface{}) int, typ interface{}) iface.LinearList {
	vec := &vector{}
	vec.Init(comparer, typ)
	return vec
}
