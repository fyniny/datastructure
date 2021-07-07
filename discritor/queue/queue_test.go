package queue

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/fyniny/datastruct/discritor/linear"
)

func TestQueue(t *testing.T) {
	comparer := func(node1, node2 interface{}) int {
		if *node1.(*int) == *node2.(*int) {
			return 0
		}
		return 1
	}

	ptr := func(i int) *int { return &i }

	Convey("test queue", t, func() {
		q := New(linear.List(comparer, new(int)))

		Convey("insert into queue", func() {
			So(q.Push(ptr(1)), ShouldEqual, nil)
			So(q.Push(ptr(2)), ShouldEqual, nil)
			So(*q.Head().(*int), ShouldEqual, 1)
			So(*q.Delete().(*int), ShouldEqual, 1)
			q.Clear()
			So(q.Length(), ShouldEqual, 0)
			q.Destroy()
			So(q.Push(ptr(2)), ShouldNotBeNil)
		})
	})
}