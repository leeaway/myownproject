package collections

type KMP struct {
	Pattern string  //模式串
	DPArray [][]int // 状态转移
}

func NewKMP(pattern string) *KMP {
	kmp := &KMP{
		Pattern: pattern,
	}
	kmp.buildTransferDPArray()
	return kmp
}

/*
	使用二位dp数组来描述，会比传统一维的Next数组更加容易理解
	dp[j][c]:当前状态为j，下一个字符为c,对应的下一个状态
	X:影子状态，所谓影子状态，就是和当前状态具有相同的前缀
	不懂请参考：https://github.com/labuladong/fucking-algorithm/blob/master/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E4%B9%8BKMP%E5%AD%97%E7%AC%A6%E5%8C%B9%E9%85%8D%E7%AE%97%E6%B3%95.md
*/
func (p *KMP) buildTransferDPArray() {
	M := len(p.Pattern)
	pat := p.Pattern
	var dp [][]int
	//初始化
	for i := 0; i < M; i++ {
		var cur []int
		for j := 0; j < 256; j++ {
			cur = append(cur, 0)
		}
		dp = append(dp, cur)
	}
	dp[0][pat[0]] = 1
	X := 0
	//构建
	for i := 0; i < M; i++ {
		for c := 0; c < 256; c++ {
			dp[i][c] = dp[X][c]
		}
		//推进到下一状态
		dp[i][pat[i]] = i + 1
		//更新影子状态
		X = dp[X][pat[i]]
	}
	p.DPArray = dp
}

func (p *KMP) Search(txt string) int {
	N := len(txt)
	M := len(p.Pattern)
	curStatus := 0
	for i := 0; i < N; i++ {
		if txt[i] == p.Pattern[curStatus] {
			curStatus++
		} else {
			curStatus = p.DPArray[curStatus][txt[i]]
		}
		if curStatus == M {
			return i - M + 1
		}
	}
	return -1
}
