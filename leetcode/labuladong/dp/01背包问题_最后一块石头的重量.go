package dp

import "example.com/m/v2/tools/mathutil"

/**
 * @author 2416144794@qq.com
 * @date 2022/12/8 15:25
 */

//有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
//
// 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
//
//
// 如果 x == y，那么两块石头都会被完全粉碎；
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
//
//
// 最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。
//
//
//
// 示例 1：
//
//
//输入：stones = [2,7,4,1,8,1]
//输出：1
//解释：
//组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
//组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
//组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
//组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
//
//
// 示例 2：
//
//
//输入：stones = [31,26,33,21,40]
//输出：5
//
//
//
//
// 提示：
//
//
// 1 <= stones.length <= 30
// 1 <= stones[i] <= 100
//
//
// Related Topics 数组 动态规划 👍 578 👎 0

/*
假设想要得到最优解，我们需要按照如下顺序操作石子：[(sa,sb),(sc,sd),...,(si,sj),(sp,sq)]。

其中 abcdijpq 代表了石子编号，字母顺序不代表编号的大小关系。

如果不考虑「有放回」的操作的话，我们可以划分为两个石子堆（正号堆/负号堆）：

将每次操作中「重量较大」的石子放到「正号堆」，代表在这次操作中该石子重量在「最终运算结果」中应用 + 运算符
将每次操作中「重量较少/相等」的石子放到「负号堆」，代表在这次操作中该石子重量在「最终运算结果」中应用 − 运算符
这意味我们最终得到的结果，可以为原来 stones 数组中的数字添加 +/− 符号，所形成的「计算表达式」所表示。

那有放回的石子重量如何考虑？

其实所谓的「有放回」操作，只是触发调整「某个原有石子」所在「哪个堆」中，并不会真正意义上的产生「新的石子重量」。

什么意思呢？

假设有起始石子 a 和 b，且两者重量关系为 a≥b，那么首先会将 a 放入「正号堆」，将 b 放入「负号堆」。重放回操作可以看作产生一个新的重量为 a−b 的“虚拟石子”，将来这个“虚拟石子”也会参与某次合并操作，也会被添加 +/− 符号：

当对“虚拟石子”添加 + 符号，即可 +(a−b)，展开后为 a−b，即起始石子 a 和 b 所在「石子堆」不变
当对“虚拟石子”添加 − 符号，即可 −(a−b)，展开后为 b−a，即起始石子 a 和 b 所在「石子堆」交换
因此所谓不断「合并」&「重放」，本质只是在构造一个折叠的计算表达式，最终都能展开扁平化为非折叠的计算表达式。

综上，即使是包含「有放回」操作，最终的结果仍然可以使用「为原来 stones 数组中的数字添加 +/− 符号，形成的“计算表达式”」所表示。

设+号部分石子和为x,-部分石子和为y，则求x+y = sum的情况下，|x-y|的最小值，不是一般性，我们可以假设x>=y,因为不满足的时候，对换即可，所以就是求x-y的最小非负值，即sum-2y的最小值，相当于从石子中挑出总和不大于sum/2的前提下的最大和

	定义：dp[n][target+1],dp[i][j]表示stones[:i+1]中价值不超过j的最大价值
	状态转移： dp[i][j] = max(dp[i-1][j],dp[i-1][j-stones[i]]+stones[i])，注意j>=stones[i]
	Base Case:
		1. dp[0][k] = stones[0],其中 k >= stones[0]
*/

//leetcode submit region begin(Prohibit modification and deletion)
func lastStoneWeightII(stones []int) int {
	sum := mathutil.Sum(stones)
	w := lastStoneWeightIIHelper(stones, sum/2)
	return sum - 2*w
}

func lastStoneWeightIIHelper(stones []int, target int) int {
	n := len(stones)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
		for j := stones[0]; j <= target && i == 0; j++ {
			dp[i][j] = stones[0]
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = mathutil.MaxInt(dp[i][j], dp[i-1][j])
			if j >= stones[i] {
				dp[i][j] = mathutil.MaxInt(dp[i][j], dp[i-1][j-stones[i]]+stones[i])
			}
		}
	}
	return dp[n-1][target]
}

//leetcode submit region end(Prohibit modification and deletion)
