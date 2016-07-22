package control_structures

import "strconv"
func fizzBuzz(i int32) string {
	if i % 3 == 0 && i % 5 == 0{
		return "FizzBuzz"
	}else if i % 3 == 0{
		return "Fizz"
	}else if i % 5 == 0{
		return "Buzz"	
	}
	return ""
}
