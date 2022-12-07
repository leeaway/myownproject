package collections

/**
 * @author 2416144794@qq.com
 * @date 2022/12/6 14:31
 */

type StringSet struct {
	entryMap map[string]bool
}

func NewStringSet(items ...string) *StringSet {
	res := &StringSet{entryMap: make(map[string]bool)}
	for _, str := range items {
		res.entryMap[str] = true
	}
	return res
}

func (s *StringSet) ToSlice() []string {
	var res []string
	for k, _ := range s.entryMap {
		res = append(res, k)
	}
	return res
}

func (s *StringSet) Size() int {
	return len(s.entryMap)
}

func (s *StringSet) Contains(key string) bool {
	return s.entryMap[key]
}

func (s *StringSet) Add(items ...string) {
	for _, item := range items {
		if !s.Contains(item) {
			s.entryMap[item] = true
		}
	}
}

func (s *StringSet) Remove(items ...string) {
	for _, item := range items {
		if s.Contains(item) {
			s.entryMap[item] = false
		}
	}
}
