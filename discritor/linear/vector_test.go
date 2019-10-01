package linear

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVector(t *testing.T) {
	ptr := func(i int) *int {
		return &i
	}
	vec := vector{}

	vec.Init(
		func(inode1, inode2 interface{}) int {
			if *(inode1.(*int)) == *(inode2.(*int)) {
				return 0
			}
			return 1
		},
		new(int),
	)

	vec.elem = append(vec.elem, ptr(1))
	vec.len = 1

	Convey("test vec linear vec: get", t, func() {
		Convey("test get succeed", func() {
			res := vec.GetElem(1).(*int)
			So(*res, ShouldEqual, 1)
		})

		Convey("test get failed", func() {
			res := vec.GetElem(2)
			So(res, ShouldEqual, nil)
		})
	})

	Convey("test vec linear vec: insert", t, func() {
		Convey("test insert succeed", func() {
			err := vec.Insert(1, ptr(3))
			So(err, ShouldEqual, nil)
			item := vec.GetElem(1)
			So(*item.(*int), ShouldEqual, 3)
		})

		Convey("test insert failed by out of range", func() {
			err := vec.Insert(4, ptr(4))
			So(err, ShouldNotBeNil)
			t.Log(err)
		})

		Convey("test insert failed by unmatched type", func() {
			err := vec.Insert(1, 3)
			So(err, ShouldNotBeNil)
			t.Log(err)
		})
	})

	Convey("test locateElem", t, func() {
		So(vec.LocateElem(ptr(3)), ShouldEqual,1)
	})

	Convey("test delete", t, func() {
		item := vec.Delete(1)
		So(*item.(*int), ShouldEqual, 3)
		So(vec.LocateElem(item), ShouldEqual, 0)
	})

	t.Log(fmt.Sprintf("%+v", vec))
}
