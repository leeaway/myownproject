package dailyExercise

//var BiggerMap = make(map[uint8][]uint8)

func alienOrder(words []string) string {
	BiggerMap := make(map[uint8][]uint8)
	for i:=0;i<len(words);i++ {
		for j:=i+1;j<len(words);j++ {
			if !compareWord(words[j],words[i],BiggerMap) {
				return ""
			}
		}
	}
	occuredCharMap := make(map[uint8]bool)
	for _,word := range words {
		for _,c := range word {
			occuredCharMap[uint8(c)] = true
		}
	}
	charToindegreeMap := make(map[uint8]int)
	for k,_ := range BiggerMap {
		charToindegreeMap[k] = 0
	}
	exist := make(map[uint8]bool)
	for k,v := range BiggerMap {
		exist[k] = true
		list := distinct(v)
		BiggerMap[k] = list
		for _,char := range list {
			charToindegreeMap[char]++
		}
	}
	var zeroQue []uint8
	for char,degree := range charToindegreeMap {
		if degree == 0 {
			zeroQue = append(zeroQue,char)
			delete(exist,char)
		}
	}
	res := ""
	for len(zeroQue) > 0 {
		curSize := len(zeroQue)
		for i:=0;i<curSize;i++ {
			top := zeroQue[0]
			delete(occuredCharMap,top)
			zeroQue = zeroQue[1:]
			res = string(top)+res
			neighbors := BiggerMap[top]
			for _,neigh := range neighbors {
				charToindegreeMap[neigh]--
				if charToindegreeMap[neigh] == 0 {
					zeroQue = append(zeroQue,neigh)
					delete(exist,neigh)
				}
			}
		}
	}
	if len(exist) > 0 {
		return ""
	}
	for k,_ := range occuredCharMap {
		res += string(k)
	}
	return res
}

func compareWord(word1,word2 string,BiggerMap map[uint8][]uint8) bool {
	l:=0
	for l<min(len(word1),len(word2)) {
		if word1[l] != word2[l] {
			BiggerMap[word1[l]] = append(BiggerMap[word1[l]],word2[l])
			return true
		}
		l++
	}
	if len(word1) == l && len(word2) > l {
		return false
	}
	return true
}

func distinct(chars []uint8) []uint8 {
	resMap := make(map[uint8]bool)
	var res []uint8
	for _,char := range chars {
		if !resMap[char] {
			res = append(res,char)
			resMap[char] = true
		}
	}
	return res
}


//func min(a,b int) int {
//	if a<b {
//		return a
//	}
//	return b
//}