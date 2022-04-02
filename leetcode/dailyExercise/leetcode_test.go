package dailyExercise

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

func TestLeecode1601(t *testing.T) {
	Convey("Test1601", t, func() {
		So(maximumRequests(5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}), ShouldEqual, 5)
		So(maximumRequests(3, [][]int{{0, 0}, {1, 2}, {2, 1}}), ShouldEqual, 3)
		So(maximumRequests(4, [][]int{{0, 3}, {3, 1}, {1, 2}, {2, 0}}), ShouldEqual, 4)
	})
}

func TestLeetcode504(t *testing.T) {
	Convey("Test504", t, func() {
		So(convertToBase7(100), ShouldEqual, "202")
		So(convertToBase7(-7), ShouldEqual, "-10")
	})
}

func TestLeetcode2055(t *testing.T) {
	Convey("test2055", t, func() {
		So(platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}}), ShouldResemble, []int{2, 3})
		So(platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}), ShouldResemble, []int{9, 0, 0, 0, 0})
	})
}

func TestLeecode2049(t *testing.T) {
	Convey("test2049", t, func() {
		So(countHighestScoreNodes([]int{-1, 2, 0, 2, 0}), ShouldEqual, 3)
		So(countHighestScoreNodes([]int{-1, 2, 0}), ShouldEqual, 2)
		So(countHighestScoreNodes([]int{-1, 8, 9, 7, 6, 2, 9, 8, 0, 0}), ShouldEqual, 2)
		So(countHighestScoreNodes([]int{-1, 0, 17, 6, 16, 16, 17, 19, 6, 4, 2, 1, 5, 11, 3, 10, 1, 0, 20, 11, 2}), ShouldEqual, 2)
	})
}

func TestLeetcode393(t *testing.T) {
	Convey("test393", t, func() {
		//So(validUtf8([]int{235,140,4}),ShouldEqual,false)
		//So(validUtf8([]int{197,130,1}),ShouldEqual,true)
		So(validUtf8([]int{255}), ShouldEqual, false)
	})
}

func TestLeetcode2039(t *testing.T) {
	//fmt.Println(networkBecomesIdle([][]int{{0,1},{1,2},{1,3},{2,4}},[]int{0,2,1,1,1}))
	//fmt.Println(networkBecomesIdle([][]int{{0,1},{0,2},{1,2}},[]int{0,10,10}))
	fmt.Println(networkBecomesIdle([][]int{{3, 8}, {4, 13}, {0, 7}, {0, 4}, {1, 8}, {14, 1}, {7, 2}, {13, 10}, {9, 11}, {12, 14}, {0, 6}, {2, 12}, {11, 5}, {6, 9}, {10, 3}}, []int{0, 3, 2, 1, 5, 1, 5, 5, 3, 1, 2, 2, 2, 2, 1}))
}

func TestLeecode2211(t *testing.T) {
	fmt.Println(countCollisions("LLRR"))
}

func TestLeetcode49(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"ddddddddddg", "dgggggggggg"}))
}

func TestLeetcode547(t *testing.T) {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}

func TestLeetcode440(t *testing.T) {
	fmt.Println(getCountWithPrefix(1, 13))
}

func TestLeetCode5253(t *testing.T) {
	Convey("test5253", t, func() {
		So(kthPalindrome([]int{90}, 3), ShouldResemble, []int64{999})
		So(kthPalindrome([]int{98043237}, 15), ShouldResemble, []int64{-1})
		So(kthPalindrome([]int{96}, 5), ShouldResemble, []int64{19591})
		So(kthPalindrome([]int{2}, 2), ShouldResemble, []int64{22})
	})
}

func TestLeetcode2024(t *testing.T) {
	Convey("Test2024", t, func() {
		So(maxConsecutiveAnswers("TTTFFFT", 2), ShouldEqual, 5)
		So(maxConsecutiveAnswers("TTFFFFT", 2), ShouldEqual, 6)
		So(maxConsecutiveAnswers("TTTTTTT", 1), ShouldEqual, 7)
		So(maxConsecutiveAnswers("TFTFTFTTTTFFFF", 1), ShouldEqual, 6)
	})
}

func TestLeetcode728(t *testing.T) {
	Convey("Test728", t, func() {
		So(selfDividingNumbers(47, 85), ShouldResemble, []int{48, 55, 66, 77})
		So(selfDividingNumbers(10, 22), ShouldResemble, []int{11, 12, 15, 22})
	})
}

func TestLeetcode954(t *testing.T) {
	Convey("test954", t, func() {
		So(canReorderDoubled([]int{1, 2, 2, 4}), ShouldEqual, true)
		So(canReorderDoubled([]int{1, 3, 3, 6}), ShouldEqual, false)
		So(canReorderDoubled([]int{4, -2, 2, -4}), ShouldEqual, true)
		So(canReorderDoubled([]int{2, 4, 0, 0, 8, 1}), ShouldEqual, true)
	})
}
