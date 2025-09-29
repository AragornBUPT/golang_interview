/* 
4的幂
https://leetcode.cn/problems/power-of-four/solutions/798268/4de-mi-by-leetcode-solution-b3ya/?envType=daily-question&envId=2025-08-15

通过 2 的幂题目，可以将非 2 的幂的数排除。这样剩下数的二进制表示都是高位为 1，低位为 0。
由题目可知，n 为 32 位有符号整数，因此我们可以构造一个 0101...01 循环的二进制数mask，mask 和 n 进行与运算，通过结果是否为 0，判断是奇数位为 1，还是偶数位为 1。
 */


package leetcode

func IsPowerOfFour(n int) bool {
	if n == 0 || n&-n != n {
		return false
	}

	mask := 0b01010101010101010101010101010101
	return n&mask != 0
}
