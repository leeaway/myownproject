package dp

import (
	"example.com/m/v2/tools/mathutil"
	"math"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/11 15:27
 */

/*
	递归做法:参考动规的分析
	时间复杂度：O(K*N*N)
*/
func MinNumberOfThrowEgg(eggNum, floorNum int) int {
	memo := make(map[string]int)
	return dpHelper(eggNum, floorNum, memo)
}

//表示k个鸡蛋，N层楼的最坏情况下的最少扔鸡蛋次数
func dpHelper(k, n int, memo map[string]int) int {
	key := buildKeyStr(k, n)
	if k == 1 {
		memo[key] = n
		return n
	}
	if n == 0 {
		memo[key] = 0
		return 0
	}

	val, ok := memo[key]
	if ok {
		return val
	}

	res := math.MaxInt
	for i := 1; i <= n; i++ {
		tmp := mathutil.MaxInt(dpHelper(k-1, i-1, memo), dpHelper(k, n-i, memo)) + 1
		res = mathutil.MinInt(res, tmp)
	}
	memo[key] = res
	return res
}

/*
	N层楼，K个鸡蛋，最坏情况下需要多少次确定鸡蛋在第几层刚好不碎
	定义：dp[K+1][N+1],dp[i][j]表示i个鸡蛋j层楼最坏情况下确定鸡蛋在第几层刚好不碎的次数
	转移：根据第i个鸡蛋在第t层碎没碎，做状态转移
		1.碎了，则需要用i-1个鸡蛋在t-1层楼找到最少需要的次数，即 dp[i-1][t-1]
		2.没碎，则需要用i-1个鸡蛋，在t+1~j层之间找，即dp[i][j-t]
		所以：dp[i][j] = max(dp[i-1][t-1],dp[i][j-t])+1(在第i楼扔了一次),其中1<=t<=j,1<=j<=N
	Base Case:
		1.dp[0][j] = maxInt,没有鸡蛋，无穷大次数，0<j<=N
		2.dp[i][0] = 0,没有楼层，不需要次数,0<i<=K
		3.dp[0][0] = ?,无法确定，所以我们保证鸡蛋K >= 1
		则有：
		1.dp[1][j] = j,1个鸡蛋，只能从1开始，最坏K次；
		2.dp[i][0] = 0,没有楼层，需要0次
	时间复杂度：O(K*N*N)
*/
func MinNumberOfThrowEgg2(K, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
		for j := 1; j <= N && i == 1; j++ {
			dp[1][j] = j
		}
	}

	for i := 2; i <= K; i++ {
		for j := 1; j <= N; j++ {
			//j层楼最多就j次
			res := j
			for t := 1; t < j; t++ {
				tmp := mathutil.MaxInt(dp[i-1][t-1], dp[i][j-t]) + 1
				res = mathutil.MinInt(res, tmp)
			}
			dp[i][j] = res
		}
	}
	return dp[K][N]
}

/*
	上述方法的复杂度过高，需要优化
	重新定义动态规划数组：
	定义：dp[k][m],表示k个鸡蛋投m次可以确定的最大楼层数,我们知道1个鸡蛋N层楼最少需要N次，而在鸡蛋不为0的时候最少次数不可能大于N，所以1<=m<=N,1<=k<=K
	转移：假设本次投鸡蛋的结果
		1. 碎了，则接下来可以投m-1次，剩k-1个鸡蛋，可以确定的最大值为 dp[k-1][m-1]
		2. 没碎，则接下来可以投m-1次，剩k个鸡蛋，可以确定的最大值为dp[k][m-1]
		总的可以确定的最大楼层dp[k][m] = dp[k-1][m-1]+dp[k][m-1]+1,加一是本次偷鸡蛋的楼也确定了
	Base Case:
		1. dp[i][1] = 1，i>0: 表示只能投一次，鸡蛋数大于0，可以确定的层数为1；
		2. dp[1][j] = j,j>0: 表示可以投j次，j>0，但是鸡蛋只有一个，只能从1层一次次往上投，可以确定的层数为j
时间复杂度：O(K*N)
*/

func MinNumberOfThrowEgg3(K, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, N+1)
	}

	//初始化
	for i := 1; i <= K; i++ {
		for j := 1; j <= N && i == 1; j++ {
			dp[1][j] = j
		}
		dp[i][1] = 1
	}

	//转移
	res := N
	for i := 2; i <= K; i++ {
		for j := 2; j <= N; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
			if dp[i][j] >= N && j < res {
				res = j
			}
		}
	}
	return res
}
