package dailyExercise

/**
 * @author 2416144794@qq.com
 * @date 2022/11/22 11:42
 */

/*
详解：
数学
提示一 : 从题面分析常见做法，从常见做法复杂度出发考虑其他做法
若不看数据范围，只看题面，容易想到的做法是「多路归并」：起始使用两个指针指向 [a, 2a, 3a, ... ] 和 [b, 2b, 3b, ...] 的开头，不断比较两指针所指向的数值大小，从而决定将谁后移，并不断更新顺位计数。

该做法常见，但其复杂度为 O(n)，对于本题 n=1e9 来说并不可行。

确定线性复杂度的做法不可行后，我们考虑是否存在对数复杂度的做法。

提示二 : 如何考虑常见的对数复杂度做法，如何定义二段性
题目要我们求第 n 个符合要求的数，假设我们想要通过「二分」来找该数值，那么我们需要分析其是否存在「二段性」。

假设在所有「能够被 a 或 b 整除的数」形成的数轴上，我们要找的分割点是 k，我们期望通过「二分」来找到 k 值，那么需要定义某种性质，使得 k 左边的数均满足该性质，k 右边的数均不满足该性质。

不难想到可根据题意来设定该性质：小于 k 的任意数字 x 满足在 [0,x] 范围数的个数不足 k 个，而大于等于 k 的任意数字 x 则不满足该性质。

提示三 : 如何实现高效的 check 函数
当确定使用「二分」来做时，剩下问题转化为：如何快速得知某个 [0,n] 中满足要求的数的个数。

容易联想到「容斥原理」：能被 a 或 b 整除的数的个数 = 能够被 a 整除的数的个数 + 能够被 b 整除的数的个数 - 既能被 a 又能被 b 整除的数的个数。
	Floor(n/a)+Floor(n/b)-Floor(n/c)
其中 c 为 a 和 b 的最小公倍数。

求解最小公倍数 lcm 需要实现最大公约数 gcd，两者模板分别为：

int gcd(int a, int b) {
    return b == 0 ? a : gcd(b, a % b);
}
int lcm(int a, int b) {
    return a * b / gcd(a, b);
}

提示四 : 如何确定值域
一个合格的值域只需要确定答案在值域范围即可，因此我们可以直接定值域大小为 1e20。

或是根据 a 和 b 的取值来大致确定：假设两者中的较大值为 m，此时第 n 个符合要求的数最大不会超过 n×m，因此也可以设定值域大小为 [0,40000n]。
*/
