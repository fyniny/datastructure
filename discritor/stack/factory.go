// factory to export stack
package stack

import "github.com/fyniny/datastruct/iface"

func New(list iface.LinearList) Stack {
	s := &stackLinear{}
	s.Init(list)
	return s
}
