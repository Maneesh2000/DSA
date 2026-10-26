package basicprograms

// Fibonacci returns the nth Fibonacci number, where Fibonacci(0) is 0 and
// Fibonacci(1) is 1. It returns 0 for a negative n.
//
// Time: O(n). Space: O(1).
func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}

	previous, current := 0, 1
	for index := 0; index < n; index++ {
		previous, current = current, previous+current
	}
	return previous
}

// IsPrime reports whether number is a prime number. Values below two are not
// prime. Checking divisors only through the square root is sufficient because
// a larger factor would require a matching smaller factor.
//
// Time: O(sqrt(n)). Space: O(1).
func IsPrime(number int) bool {
	if number < 2 {
		return false
	}
	for divisor := 2; divisor <= number/divisor; divisor++ {
		if number%divisor == 0 {
			return false
		}
	}
	return true
}
