package interview

import (
	"math"
	"strconv"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/2/3 16:07
 */

var res int

func getRemainNum(str string, k int) int {
	res = math.MaxInt
	dfsHepler(str, 0, k, "")
	return res
}

/*
index:	当前在第几个元素，可以选择是否要删除当前的元素，当然k需要大于0
cur: 当前决策后剩余的字符串
*/
func dfsHepler(str string, index, k int, cur string) {
	//遍历完了，递归终止
	if index == len(str) {
		tmp := strToInt(cur)
		if tmp < res {
			res = tmp
		}
		return
	}
	//不能再删了,递归终止
	if k == 0 {
		tmp := strToInt(cur + str[index:])
		if tmp < res {
			res = tmp
		}
		return
	}

	//删除当前
	dfsHepler(str, index+1, k-1, cur)
	//不删除当前
	dfsHepler(str, index+1, k, cur+str[index:index+1])
}

func strToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
