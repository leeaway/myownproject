package leetcode

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLeetcode688(t *testing.T) {
	Convey("TestLeetcode688", t, func() {
		So(KnightProbability2(3, 2, 0, 0), ShouldEqual, 0.0625)
		So(KnightProbability2(1, 0, 0, 0), ShouldEqual, float64(1))
		So(KnightProbability2(3, 2, 1, 1), ShouldEqual, float64(0))
	})
}

func TestLeetcode838(t *testing.T) {
	Convey("TestLeeCode838", t, func() {
		So(PushDominoes("RR.L"), ShouldEqual, "RR.L")
		So(PushDominoes(".L.R...LR..L.."), ShouldEqual, "LL.RR.LLRRLL..")
	})
}

func TestLeetcode131(t *testing.T) {
	Partition("aab")
	fmt.Println(NumberOfGoodSubsets([]int{4, 2, 3, 15}))
}

func TestLeecode917(t *testing.T) {
	Convey("TestLeetcode917", t, func() {
		So(reverseOnlyLetters("Test1ng-Leet=code-Q!"), ShouldEqual, "Qedo1ct-eeLg=ntse-T!")
		So(reverseOnlyLetters("abcde"), ShouldEqual, "edcba")
		So(reverseOnlyLetters(""), ShouldEqual, "")
		So(reverseOnlyLetters("123"), ShouldEqual, "123")
	})
}

func TestLeetcpde1706(t *testing.T) {
	testCase1 := [][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}
	testCase2 := [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}
	Convey("TestLeetcode1706", t, func() {
		So(findBall(testCase1), ShouldResemble, []int{0, 1, 2, 3, 4, -1})
		So(findBall(testCase2), ShouldResemble, []int{1, -1, -1, -1, -1})
	})
}

func TestLeetcode139(t *testing.T) {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func TestLeetcode537(t *testing.T) {
	//Convey("TestLeetcode537",t, func() {
	//	So(complexNumberMultiply("1+1i","1+1i"),ShouldEqual,"0+2i")
	//	So(complexNumberMultiply("1+-1i","1+-1i"),ShouldEqual,"0+-2i")
	//	So(complexNumberMultiply("2+2i","1-2i"),ShouldEqual,"-2+2i")
	//})
	fmt.Println(complexNumberMultiply("78+-76i", "-86+72i"))
}
