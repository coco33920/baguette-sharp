package utils

func Fibonacci(n int) int {
	b := [...]int{1,1}
	for i := 2; i < n; i++ {
		b[i%2] = b[0] + b[1]
	}
	return b[1]
}
