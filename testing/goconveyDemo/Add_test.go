package goconveyDemo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestSubtract(t *testing.T) {
	Convey("将两数相减", t, func() {
		So(Subtract(1, 2), ShouldEqual, -1)
	})
}

func TestMultiply(t *testing.T) {
	Convey("将两数相乘", t, func() {
		So(Multiply(3, 2), ShouldEqual, 6)
	})
}

func TestDivision(t *testing.T) {
	Convey("将两数相除", t, func() {

		Convey("除以非 0 数", func() {
			num, err := Division(10, 2)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, 5)
		})

		Convey("除以 0", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})
	})
}


/*
F:\airdroid_code\go\src\go_learn_demo\testing\goconveyDemo>go test -v
=== RUN   TestAdd

将两数相加 .


1 total assertion

--- PASS: TestAdd (0.00s)
=== RUN   TestSubtract

将两数相减 .


2 total assertions

--- PASS: TestSubtract (0.00s)
=== RUN   TestMultiply

将两数相乘 .


3 total assertions

--- PASS: TestMultiply (0.00s)
=== RUN   TestDivision

将两数相除
除以非 0 数 ..
除以 0 .


6 total assertions

--- PASS: TestDivision (0.00s)
PASS
ok      go_learn_demo/testing/goconveyDemo      0.447s
*/
