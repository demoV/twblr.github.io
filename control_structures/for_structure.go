package control_structures

func collatzChainLength(n int) int {
	cons := 0
	for (n != 1) {
		if(n % 2 != 0) {
			n = 3 * n + 1
			cons ++
		}else{
			n = n / 2
			cons ++
		}	
	}
	return cons
	
}
