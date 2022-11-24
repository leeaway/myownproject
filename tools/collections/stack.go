package collections

/**
 * @author 2416144794@qq.com
 * @date 2022/11/24 18:48
 */

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

func (s *Stack) IsEmpty() bool {
	return len(s.nums) == 0
}

func (s *Stack) Push(a int) {
	s.nums = append(s.nums, a)
}

func (s *Stack) Poll() int {
	if s.IsEmpty() {
		return -1
	}
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.nums[len(s.nums)-1]
}
