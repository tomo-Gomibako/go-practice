package fibonacci

// Fib fibonacci function
func Fib(n int) int {
	if n < 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

var memo = []int{1, 1}

// Memo fibonacci function with memo
func Memo(n int) int {
	if len(memo) > n {
		return memo[n]
	}

	ret := Memo(n-1) + Memo(n-2)
	memo = append(memo, ret)
	return ret
}
