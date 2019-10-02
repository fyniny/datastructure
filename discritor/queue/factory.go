package queue

import "github.com/fyniny/datastruct/iface"

// New returns queue, it is suggested to use list as the lower storage for the queue
// for that, the queue doesn't implement cycle queue, while may be waste too many memory resource
func New(linear iface.LinearList) Queue {
	q :=  &queueLinear{}
	q.Init(linear)
	return q
}
