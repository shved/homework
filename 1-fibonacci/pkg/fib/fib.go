package fib

func Calc(seqStep int) int {
	if seqStep == 1 {
		return 0
	}

	if seqStep == 2 {
		return 1
	}

	cur1, cur2 := 0, 1
	for i := 3; i <= seqStep; i++ {
		cur1, cur2 = cur2, cur1+cur2
	}

	return cur2
}
