package algor_lib

func Fib1(n uint) uint {
	if n < 2 {
		return n
	}

	var a, b uint
	a = 0
	b = 1
	for i := uint(2); i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}

func Fib2(n uint) uint {
	tmpMap := make(map[uint]uint)
	return fib2(n, tmpMap)
}

func fib2(n uint, tmpMap map[uint]uint) uint {
	if n < 2 {
		return n
	}

	tmpMap[n] = fib2(n-2, tmpMap) + fib2(n-1, tmpMap)
	return tmpMap[n]
}
