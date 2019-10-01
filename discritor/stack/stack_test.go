package stack

import (
	"github.com/fyniny/datastruct/discritor/linear"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStack(t *testing.T) {
	Convey("test stack", t, func() {
		comparer := func(inode1, inode2 interface{}) int {
			if *inode1.(*int) == *inode2.(*int) {
				return 0
			}
			return 1
		}
		ptr := func(i int) *int {return &i}

		sta := New(linear.Vector(comparer, new(int)))

		Convey("push stack", func() {
			sta.Push(ptr(1))
			sta.Push(ptr(2))
			So(*sta.Top().(*int), ShouldEqual, 2)
			So(*sta.Pop().(*int), ShouldEqual, 2)
			So(*sta.Top().(*int), ShouldEqual, 1)
			So(sta.Length(), ShouldEqual, 1)
		})
	})
}