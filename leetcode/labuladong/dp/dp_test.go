package dp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/12/7 14:48
 */

func Test_LengthOfLIS(t *testing.T) {
	Convey("LengthOfLIS", t, func() {
		Convey("LengthOfLIS test1", func() {
			So(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}), ShouldEqual, 6)
			So(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}), ShouldEqual, 4)
			So(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}), ShouldEqual, 4)
		})
	})
}

func Test_MaxEnvelopes(t *testing.T) {
	Convey("MaxEnvelopes", t, func() {
		Convey("MaxEnvelopes test1", func() {
			So(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}), ShouldEqual, 3)
			So(maxEnvelopes([][]int{{1, 4}, {2, 4}, {3, 3}, {2, 3}}), ShouldEqual, 1)
		})
	})
}

func Test_MinDistance(t *testing.T) {
	Convey("MinDistance", t, func() {
		Convey("MinDistance test1", func() {
			So(minDistance("abcde", "kde"), ShouldEqual, 3)
			So(minDistance("horse", "ros"), ShouldEqual, 3)
			So(minDistance("intention", "execution"), ShouldEqual, 5)
		})

		//递归做法
		Convey("MinDistance2 test1", func() {
			So(minDistance2("abcde", "kde"), ShouldEqual, 3)
			So(minDistance2("horse", "ros"), ShouldEqual, 3)
			So(minDistance2("intention", "execution"), ShouldEqual, 5)
		})

		//记忆化递归
		Convey("MinDistance3 test1", func() {
			So(minDistance3("abcde", "kde"), ShouldEqual, 3)
			So(minDistance3("horse", "ros"), ShouldEqual, 3)
			So(minDistance3("intention", "execution"), ShouldEqual, 5)
		})
	})
}

func Test_LongestCommonSubString(t *testing.T) {
	Convey("LongestCommonSubString", t, func() {
		Convey("LongestCommonSubString test1", func() {
			So(longestCommonSubsequence("abcde", "ace"), ShouldEqual, 3)
			So(longestCommonSubsequence("abcde", "ce"), ShouldEqual, 2)
			So(longestCommonSubsequence("abcde", "fg"), ShouldEqual, 0)
		})
	})
}
