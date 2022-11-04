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

func TestLeetcode420(t *testing.T) {
	Convey("test420", t, func() {
		So(strongPasswordChecker("1337C0d3"), ShouldEqual, 0)
		So(strongPasswordChecker("1333C0d3"), ShouldEqual, 1)
	})
}

func TestLeetCode746(t *testing.T) {
	Convey("test746", t, func() {
		So(minCostClimbingStairs([]int{10, 15, 20}), ShouldEqual, 15)
	})
}

func TestLeetcode54(t *testing.T) {
	Convey("test54", t, func() {
		So(jump([]int{2, 3, 1, 1, 4}), ShouldEqual, 2)
	})
}

func TestLeetcode310(t *testing.T) {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
}

func TestLeetCode1567(t *testing.T) {
	Convey("test", t, func() {
		So(getMaxLen([]int{0, 1, -2, -3, -4}), ShouldEqual, 3)
	})
}

func TestLeet380(t *testing.T) {
	set := Constructor()
	set.Insert(0)
	set.Insert(1)
	set.Remove(0)
	set.Insert(2)
	set.Remove(1)
	fmt.Println(set.GetRandom())
}

func DeferTest(flag bool) int {
	a := 0
	if flag {
		return a
	}
	defer func() {
		fmt.Println("defer")
	}()
	return a
}

func TestDefer(t *testing.T) {
	Convey("test", t, func() {
		So(DeferTest(true), ShouldEqual, 0)
		So(DeferTest(false), ShouldEqual, 0)
	})
}

func Test905(t *testing.T) {
	Convey("test", t, func() {
		So(sortArrayByParity([]int{3, 1, 2, 4}), ShouldResemble, []int{4, 2, 1, 3})
		So(sortArrayByParity([]int{1, 2, 3, 4}), ShouldResemble, []int{4, 2, 3, 1})
	})
}

func Test2249(t *testing.T) {
	Convey("test", t, func() {
		So(countLatticePoints([][]int{{2, 2, 2}}), ShouldEqual, 25)
	})
}

func Test6043(t *testing.T) {
	fmt.Println(countRectangles([][]int{{1, 1}, {1, 3}, {2, 1}, {2, 3}, {3, 3}}, [][]int{{1, 2}, {1, 1}, {2, 3}}))
}

func Test713(t *testing.T) {
	fmt.Println(numSubarrayProductLessThanK([]int{1, 1, 1}, 1))
}

func Test433(t *testing.T) {
	Convey("test433", t, func() {
		So(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}), ShouldEqual, 2)
		So(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}), ShouldEqual, 3)
	})
}

func Test942(t *testing.T) {
	fmt.Println(diStringMatch("IDID"))
}

func Test675(t *testing.T) {
	Convey("test675", t, func() {
		//So(cutOffTree([][]int{{1,2,1},{0,0,4},{7,6,5}}),ShouldEqual,6)
		//So(cutOffTree([][]int{{1,2,1},{3,0,4},{7,6,5}}),ShouldEqual,10)
		//So(cutOffTree([][]int{{1,2,1},{0,0,0},{7,6,5}}),ShouldEqual,-1)
		So(cutOffTree([][]int{{54581641, 64080174, 24346381, 69107959}, {86374198, 61363882, 68783324, 79706116}, {668150, 92178815, 89819108, 94701471}, {83920491, 22724204, 46281641, 47531096}, {89078499, 18904913, 25462145, 60813308}}), ShouldEqual, 57)
	})
}

func Test467(t *testing.T) {
	fmt.Println(findSubstringInWraproundString("abcdefghijklmnopqrstuvwxyzab"))
}

func Test699(t *testing.T) {
	fmt.Println(fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}}))
	fmt.Println(fallingSquares([][]int{{100, 100}, {200, 100}, {6, 1}}))
	fmt.Println(fallingSquares([][]int{{0, 5}, {2, 3}, {0, 1}}))
}

func Test17_11(t *testing.T) {
	fmt.Println(findClosest([]string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}, "a", "student"))
}

func Test114(t *testing.T) {
	fmt.Println(alienOrder([]string{"z", "x", "z"}))
	fmt.Println(alienOrder([]string{"abc", "ab"}))
}

func Test473(t *testing.T) {
	fmt.Println(makesquare([]int{13, 11, 1, 8, 6, 7, 8, 8, 6, 7, 8, 9, 8}))
	fmt.Println(makesquare([]int{1, 1, 2, 3, 3, 4, 2}))
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
}

func Test1175(t *testing.T) {
	fmt.Println(numPrimeArrangements(100))
}

func Test556(t *testing.T) {
	nextPermutation([]int{2, 3, 1})
}

func Test801(t *testing.T) {
	fmt.Println(minSwap([]int{0, 7, 8, 10, 10, 11, 12, 13, 19, 18}, []int{4, 4, 5, 7, 11, 14, 15, 16, 17, 20}))
}

func Test1790(t *testing.T) {
	Convey("1790", t, func() {
		So(areAlmostEqual("bank", "kanb"), ShouldEqual, true)
		So(areAlmostEqual("bank", "knab"), ShouldEqual, false)
		So(areAlmostEqual("bank", "bank"), ShouldEqual, true)
	})
}

func Test817(t *testing.T) {
	Convey("链表组件", t, func() {
		node3 := &ListNode{
			Val:  3,
			Next: nil,
		}
		node2 := &ListNode{
			Val:  2,
			Next: node3,
		}
		node1 := &ListNode{
			Val:  1,
			Next: node2,
		}
		head := &ListNode{
			Val:  0,
			Next: node1,
		}
		So(numComponents(head, []int{0, 1, 3}), ShouldEqual, 2)
		So(numComponents(head, []int{1, 3, 0}), ShouldEqual, 2)
	})
}

func Test1320(t *testing.T) {
	Convey("test1320", t, func() {
		So(minimumDistance("CAKE"), ShouldEqual, 3)
		So(minimumDistance("HAPPY"), ShouldEqual, 6)
	})
}

func Test940(t *testing.T) {
	Convey("test940", t, func() {
		So(distinctSubseqII("abc"), ShouldEqual, 7)
		So(distinctSubseqII("aaa"), ShouldEqual, 3)
		So(distinctSubseqII("aba"), ShouldEqual, 6)
		So(distinctSubseqII("lee"), ShouldEqual, 5)
	})
}

func Test904(t *testing.T) {
	Convey("水果成篮", t, func() {
		So(totalFruit([]int{1}), ShouldEqual, 1)
		So(totalFruit([]int{1, 2, 2}), ShouldEqual, 3)
		So(totalFruit([]int{0, 1, 2, 2}), ShouldEqual, 3)
		So(totalFruit([]int{1, 2, 3, 2, 2}), ShouldEqual, 4)
		So(totalFruit([]int{1, 0, 1, 4, 1, 4, 1, 2, 3}), ShouldEqual, 5)
		So(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}), ShouldEqual, 5)
	})
}

func Test902(t *testing.T) {
	Convey("最大为N的数字组合", t, func() {
		So(atMostNGivenDigitSet([]string{"1", "7"}, 231), ShouldEqual, 10)
		So(atMostNGivenDigitSet([]string{"1", "2", "3", "4"}, 1), ShouldEqual, 1)
		So(atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100), ShouldEqual, 20)
		So(atMostNGivenDigitSet([]string{"1", "4", "9"}, 1e9), ShouldEqual, 29523)
		So(atMostNGivenDigitSet([]string{"3", "4", "5", "6"}, 64), ShouldEqual, 18)
	})
}

func Test1700(t *testing.T) {
	Convey("无法吃午餐的学生数量", t, func() {
		So(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}), ShouldEqual, 0)
		So(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}), ShouldEqual, 3)
	})
}

func Test915(t *testing.T) {
	Convey("分割数组", t, func() {
		So(partitionDisjoint([]int{1, 1, 1, 1, 1}), ShouldEqual, 1)
		So(partitionDisjoint([]int{5, 0, 3, 8, 6}), ShouldEqual, 3)
		So(partitionDisjoint([]int{1, 2, 3, 4, 5, 6}), ShouldEqual, 1)
		So(partitionDisjoint([]int{1, 3, 2, 4, 5, 7}), ShouldEqual, 1)
		So(partitionDisjoint([]int{1, 1, 1, 0, 6, 12}), ShouldEqual, 4)

		So(partitionDisjoint2([]int{1, 1, 1, 1, 1}), ShouldEqual, 1)
		So(partitionDisjoint2([]int{5, 0, 3, 8, 6}), ShouldEqual, 3)
		So(partitionDisjoint2([]int{1, 2, 3, 4, 5, 6}), ShouldEqual, 1)
		So(partitionDisjoint2([]int{1, 3, 2, 4, 5, 7}), ShouldEqual, 1)
		So(partitionDisjoint2([]int{1, 1, 1, 0, 6, 12}), ShouldEqual, 4)
	})
}

func Test1822(t *testing.T) {
	Convey("数组符号", t, func() {
		So(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}), ShouldEqual, 1)
		So(arraySign([]int{-1, 1, -1, 1, -1}), ShouldEqual, -1)
		So(arraySign([]int{-1, 1, -1, 1, -1, 0, 1}), ShouldEqual, 0)
	})
}

func Test754(t *testing.T) {
	Convey("到达终点数字", t, func() {
		So(reachNumber(2), ShouldEqual, 3)
		So(reachNumber(3), ShouldEqual, 2)
		So(reachNumber(120), ShouldEqual, 15)
		So(reachNumber(1200), ShouldEqual, 51)
	})
}
