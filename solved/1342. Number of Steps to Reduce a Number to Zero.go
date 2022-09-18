package solved

func numberOfSteps(num int) int {
	var res int
	for num != 0 {
		if num%2 != 0 && num > 1 {
			num -= 1
			res++
		}
		num /= 2
		res++
	}
	return res
}
