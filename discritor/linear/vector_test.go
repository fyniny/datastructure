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

	Convey("test vec linear vec: get", t, func() {
		Convey("insert the first node", func() {
			err := vec.Insert(1, ptr(1))
			So(err, ShouldEqual, nil)
		})

		Convey("test get succeed", func() {
			res := vec.GetElem(1).(*int)
			So(*res, ShouldEqual, 1)
		})

		Convey("test get failed", func() {
			res := vec.GetElem(2)
			So(res, ShouldEqual, nil)
		})

		Convey("test insert at the end", func() {
			So(vec.Insert(vec.Length()+1, ptr(24)), ShouldEqual, nil)
			So(*vec.GetElem(vec.Length()).(*int), ShouldEqual, 24)
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
			err := vec.Insert(vec.Length()+2, ptr(4))
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

	Convey("text expand", t, func() {
		for i := vec.Length()+1; i <= minimum; i++ {
			_ = vec.Insert(i, ptr(i))
		}
		_ = vec.Insert(vec.Length()+1, ptr(99))
		So(*vec.GetElem(vec.Length()).(*int), ShouldEqual, 99)
	})

	t.Log(fmt.Sprintf("%+v", vec))
}
