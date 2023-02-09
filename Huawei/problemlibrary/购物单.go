package problemlibrary

import "fmt"

/**
 * @author 2416144794@qq.com
 * @date 2023/2/6 11:47
 */

/*
王强决定把年终奖用于购物，他把想买的物品分为两类：主件与附件，附件是从属于某个主件的，下表就是一些主件与附件的例子：
主件	附件
电脑	打印机，扫描仪
书柜	图书
书桌	台灯，文具
工作椅	无
如果要买归类为附件的物品，必须先买该附件所属的主件，且每件物品只能购买一次。
每个主件可以有 0 个、 1 个或 2 个附件。附件不再有从属于自己的附件。
王强查到了每件物品的价格（都是 10 元的整数倍），而他只有 N 元的预算。除此之外，他给每件物品规定了一个重要度，用整数 1 ~ 5 表示。他希望在花费不超过 N 元的前提下，使自己的满意度达到最大。
请你帮助王强计算可获得的最大的满意度。


输入描述：
输入的第 1 行，为两个正整数N，m，用一个空格隔开：

（其中 N （ N<32000 ）表示总钱数， m （m <60 ）为可购买的物品的个数。）

从第 2 行到第 m+1 行，第 j 行给出了编号为 j-1 的物品的基本数据，每行有 3 个非负整数 v p q
（其中 v 表示该物品的价格（ v<10000 ）， p 表示该物品的重要度（ 1 ~ 5 ）， q 表示该物品是主件还是附件。如果 q=0 ，表示该物品为主件，如果 q>0 ，表示该物品为附件， q 是所属主件的编号）




输出描述：
 输出一个正整数，为张强可以获得的最大的满意度。
示例1
输入：
1000 5
800 2 0
400 5 1
300 5 1
400 3 0
500 2 0
复制
输出：
2200
复制
示例2
输入：
50 5
20 3 5
20 3 5
10 3 0
10 2 0
10 1 0
复制
输出：
130
复制
说明：
由第1行可知总钱数N为50以及希望购买的物品个数m为5；
第2和第3行的q为5，说明它们都是编号为5的物品的附件；
第4~6行的q都为0，说明它们都是主件，它们的编号依次为3~5；
所以物品的价格与重要度乘积的总和的最大值为10*1+20*3+20*3=130
*/

/*
本质上还是主件的01背包问题，只不过我们还需要考虑附件
things：所有物品
mainItems: 记录所有主件的索引
mainToAttachMap: 记录主件到附件的索引映射关系

定义：dp[m][n],表示m个主件，钱数为n的情况下，可以拿到的最大满意度
转移：1. j < prices[i] || 不买当前主件：dp[i-1][j]
	 2. j >= prices[i] && 买当前主件：
		2.1 只买主件：dp[i-1][j-prices[i]]+satisfy[i]
		2.2 j>=price[i]+prices[attach1]&&购买附件1：dp[i-1][j-prices[i]-prices[attach1]]+satisfy[i]+satisfy[attach1]
		2.3 j>=price[i]+prices[attach2]&&购买附件2：dp[i-1][j-prices[i]-prices[attach2]]+satisfy[i]+satisfy[attach2]
		2.4 j>=price[i]+prices[attach1]+prices[attach2] &&购买附件1&&购买附件2：
			dp[i-1][j-prices[i]-prices[attach1]-prices[attach2]]+satisfy[i]+satisfy[attach1]+satisfy[attach2]

Base Case: 均初始化为0即可
*/

func shoppingList(things [][]int, mainItems []int, mainToAttachListMap map[int][]int, money int) int {
	l := len(mainItems)
	if l == 0 {
		return 0
	}
	m := len(mainItems)
	dp := make([][]int, m+1)
	for i := 0; i <= len(mainItems); i++ {
		dp[i] = make([]int, money+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= money; j++ {
			mainIndex := mainItems[i-1]
			price := things[mainIndex][0]
			sum := things[mainIndex][0] * things[mainIndex][1]
			if j < price {
				continue
			}
			dp[i][j] = max(dp[i][j], max(dp[i-1][j], dp[i-1][j-price]+sum))
			attachIndexes := mainToAttachListMap[mainIndex]
			if len(attachIndexes) == 0 {
				continue
			}
			attachPrice0 := things[attachIndexes[0]][0]
			attachSum0 := things[attachIndexes[0]][0] * things[attachIndexes[0]][1]
			if j-price-attachPrice0 >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-price-attachPrice0]+attachSum0+sum)
			}
			if len(attachIndexes) < 2 {
				continue
			}
			attachPrice1 := things[attachIndexes[1]][0]
			attachSum1 := things[attachIndexes[1]][0] * things[attachIndexes[1]][1]
			if j-price-attachPrice1 >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-price-attachPrice1]+attachSum1+sum)
			}
			if j-price-attachPrice1-attachPrice0 >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-price-attachPrice1-attachPrice0]+attachSum1+attachSum0+sum)
			}
		}
	}
	return dp[m][money]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func CalShoppingList() int {
	m := 0
	n := 0
	fmt.Scan(&n, &m)
	price, importance, mainIndex := 0, 0, 0
	var things [][]int
	mainToAttachIdxMap := make(map[int][]int)
	var mainItemIdxes []int
	for i := 0; i < m; i++ {
		fmt.Scan(&price, &importance, &mainIndex)
		curThing := []int{price, importance, mainIndex}
		things = append(things, curThing)
		if mainIndex == 0 {
			mainItemIdxes = append(mainItemIdxes, i)
			continue
		}
		mainToAttachIdxMap[mainIndex-1] = append(mainToAttachIdxMap[mainIndex-1], i)
	}

	return shoppingList(things, mainItemIdxes, mainToAttachIdxMap, n)
}
