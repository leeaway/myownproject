package dailyExercise

/**
 * @author 2416144794@qq.com
 * @date 2022/12/12 10:46
 */

//一个字符串的 美丽值 定义为：出现频率最高字符与出现频率最低字符的出现次数之差。
//
//
// 比方说，"abaacc" 的美丽值为 3 - 1 = 2 。
//
//
// 给你一个字符串 s ，请你返回它所有子字符串的 美丽值 之和。
//
//
//
// 示例 1：
//
//
//输入：s = "aabcb"
//输出：5
//解释：美丽值不为零的字符串包括 ["aab","aabc","aabcb","abcb","bcb"] ，每一个字符串的美丽值都为 1 。
//
// 示例 2：
//
//
//输入：s = "aabcbaa"
//输出：17
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 500
// s 只包含小写英文字母。
//
//
// Related Topics 哈希表 字符串 计数 👍 35 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func beautySum(s string) int {
	n := len(s)
	//count[i][j][k],表示s[i:j+1]中字符k的数量
	count := make([][][]int, n)
	preSum := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		if i == n {
			preSum[i] = make([]int, 26)
			continue
		}
		count[i] = make([][]int, n)
		preSum[i] = make([]int, 26)
		for j := 0; j < n; j++ {
			count[i][j] = make([]int, 26)
		}
	}

	for i := 1; i <= len(s); i++ {
		cur := s[i-1] - 'a'
		for j := 0; j < 26; j++ {
			preSum[i][j] = preSum[i-1][j]
			if j == int(cur) {
				preSum[i][j] += 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := 0; k < 26; k++ {
				count[i][j][k] = preSum[j+1][k] - preSum[i][k]
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res += checkAndGetMaxGap(count[i][j])
		}
	}
	return res
}

func checkAndGetMaxGap(nums []int) int {
	count := 0
	minV, maxV := 500, 0
	for _, num := range nums {
		if num == 0 {
			continue
		}
		count++
		if num > maxV {
			maxV = num
		}
		if num < minV {
			minV = num
		}
	}
	if count < 2 {
		return 0
	}
	return maxV - minV
}

//leetcode submit region end(Prohibit modification and deletion)
