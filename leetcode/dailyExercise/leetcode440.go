package dailyExercise

//获取不大于n的以prefix开头的数字的个数
//字典序中*10等于自然序中+1
func getCountWithPrefix(prefix, n int) int {
	count := 0
	for a, b := prefix, prefix+1; a <= n; a, b = a*10, b*10 {
		//这里因为n为右闭区间，b为右开区间，所以计数其实是min(n,b-1)-a+1，简化一下就是min(n+1,b) - a
		count += min(n+1, b) - a
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findKthNumber(n, k int) int {
	prefix, pos := 1, 1
	for pos < k {
		//当前前缀的数字数量
		curCount := getCountWithPrefix(prefix, n)
		//说明加上当前前缀的所有数都还不够，前缀就要+1，pos += curCount
		if pos+curCount <= k {
			prefix++
			pos += curCount
		} else if pos+curCount > k {
			//说明加上当前前缀的所有数超了，答案必定是这个前缀的数,这是缩小查询范围，prefix *= 10, pos++(因为前缀多了一位)
			prefix *= 10
			pos++
		}
	}
	return prefix
}
