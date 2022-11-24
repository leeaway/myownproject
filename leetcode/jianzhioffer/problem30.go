package jianzhioffer

/**
 * @author 2416144794@qq.com
 * @date 2022/11/24 17:15
 */

//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
//
//
//
// 示例:
//
// MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.min();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.min();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// 各函数的调用总次数不超过 20000 次
//
//
//
//
// 注意：本题与主站 155 题相同：https://leetcode-cn.com/problems/min-stack/
//
// Related Topics 栈 设计 👍 405 👎 0

/*
	方法：只需要维护一个与原栈相同大小的min栈即可，每次原栈的操作都会实时更新min栈
	如：nums: 3,4,1,2,5
       minV: 3,3,1,1,1
*/

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	nums []int
	minV []int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{nums: []int{}, minV: []int{}}
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)
	if len(this.minV) == 0 {
		this.minV = append(this.minV, x)
		return
	}
	if this.Min() > x {
		this.minV = append(this.minV, x)
		return
	}
	this.minV = append(this.minV, this.Min())
}

func (this *MinStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
	this.minV = this.minV[:len(this.minV)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) Min() int {
	return this.minV[len(this.minV)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
//leetcode submit region end(Prohibit modification and deletion)
