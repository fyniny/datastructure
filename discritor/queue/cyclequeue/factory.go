package cyclequeue

import (
	"fmt"
	"reflect"

	"github.com/fyniny/datastruct/discritor/queue"
	"github.com/fyniny/datastruct/iface"
)

func New(linear iface.LinearList, size int) queue.Queue {
	if size == 0 {
		panic("queue size can't be zero")
	}

	cq := &cycle{
		size: size,
	}

	cq.Init(linear)
	return cq
}

type  cSS struct {
	ss int
}
func Resize(queue queue.Queue, size int) error {
	resizeFunc := reflect.ValueOf(queue).MethodByName("Resize")
	if !resizeFunc.IsValid() {
		return fmt.Errorf("the queue doesn't have Resize method")
	}
	ret := resizeFunc.Call([]reflect.Value{reflect.ValueOf(size)})
	if ret[0].IsNil() {
		return nil
	}
	return ret[0].Interface().(error)
}
