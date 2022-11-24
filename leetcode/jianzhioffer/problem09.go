package jianzhioffer

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/10/12 17:44
 */

//功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
//
//
// 示例 1：
//
//
//输入：
//["CQueue","appendTail","deleteHead","deleteHead","deleteHead"]
//[[],[3],[],[],[]]
//输出：[null,null,3,-1,-1]
//
//
// 示例 2：
//
//
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
//
//
// 提示：
//
//
// 1 <= values <= 10000
// 最多会对 appendTail、deleteHead 进行 10000 次调用
//
//
// Related Topics 栈 设计 队列 👍 616 👎 0

/*
方法：
一个入栈和一个出栈，append直接append入栈；
出栈的话，若出栈中还有数据，直接poll，否者将入栈的数据全部offer到出栈中
*/

type CQueue struct {
	inStack, outStack *collections.Stack
}

func Constructor() CQueue {
	return CQueue{
		inStack:  collections.NewStack(),
		outStack: collections.NewStack(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if !this.outStack.IsEmpty() {
		return this.outStack.Poll()
	}
	for !this.inStack.IsEmpty() {
		this.outStack.Push(this.inStack.Poll())
	}
	return this.outStack.Poll()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
