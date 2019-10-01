package linear

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestList(t *testing.T) {
	ptr := func(i int) *int {
		return &i
	}
	list := list{}

	list.Init(
		func(inode1, inode2 interface{}) int {
			if *(inode1.(*int)) == *(inode2.(*int)) {
				return 0
			}
			return 1
		},
		new(int),
	)

	list.head.next =  &node{
		elem: ptr(1),
		next: nil,
	}

	list.len = 1

	Convey("test list linear list: get", t, func() {
		Convey("test get succeed", func() {
			res := list.GetElem(1).(*int)
			So(*res, ShouldEqual, 1)
		})

		Convey("test get failed", func() {
			res := list.GetElem(2)
			So(res, ShouldEqual, nil)
		})
	})

	Convey("test list linear list: insert", t, func() {
		Convey("test insert succeed", func() {
			err := list.Insert(1, ptr(3))
			So(err, ShouldEqual, nil)
			item := list.GetElem(1)
			So(*item.(*int), ShouldEqual, 3)
		})

		Convey("test insert failed by out of range", func() {
			err := list.Insert(4, ptr(4))
			So(err, ShouldNotBeNil)
			t.Log(err)
		})

		Convey("test insert failed by unmatched type", func() {
			err := list.Insert(1, 3)
			So(err, ShouldNotBeNil)
			t.Log(err)
		})
	})

	Convey("test locateElem", t, func() {
		So(list.LocateElem(ptr(3)), ShouldEqual,1)
	})

	Convey("test delete", t, func() {
		item := list.Delete(1)
		So(*item.(*int), ShouldEqual, 3)
		So(list.LocateElem(item), ShouldEqual, 0)
	})
	t.Log(fmt.Sprintf("%+v", list))
}
