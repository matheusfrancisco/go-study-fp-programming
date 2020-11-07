package fib

func FibSimple(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return FibSimple(x-2) + FibSimple(x-1)
	}
}

type Memoized func(int) int

var fibMem Memoized

func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

func FibMemoized(n int) int {
	n += 1
	fibMem = Memoize(func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		return fibMem(n-2) + fibMem(n-1)
	})
	return fibMem(n)
}
