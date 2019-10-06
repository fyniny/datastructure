package cyclequeue

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/fyniny/datastruct/discritor/linear"
)

func TestCycleQueue(t *testing.T) {
	comparer := func(node1, node2 interface{}) int {
		if *node1.(*int) == *node2.(*int) {
			return 0
		}
		return 1
	}

	ptr := func(i int) *int { return &i }

	Convey("test cycle queue", t, func() {
		cq := New(linear.List(comparer, new(int)), *ptr(1))
		So(cq.Push(ptr(1)), ShouldEqual, nil)
		So(cq.Push(ptr(2)), ShouldNotBeNil)
		So(Resize(cq, 2), ShouldEqual, nil)
		So(cq.Push(ptr(2)), ShouldEqual, nil)
		So(Resize(cq, 1), ShouldNotBeNil)
		So(*cq.Head().(*int), ShouldEqual, 1)
		So(*cq.Delete().(*int), ShouldEqual, 1)
		So(*cq.Delete().(*int), ShouldEqual, 2)
		So(Resize(cq, 1), ShouldEqual, nil)
	})
}
