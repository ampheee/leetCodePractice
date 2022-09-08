package main

//func romanToInt(s string) int {
//	var totalInt int
//	for i := 0; i < len(s); i++ {
//		switch s[i] {
//		case 'I':
//			totalInt++
//		case 'V':
//			if i >= 1 {
//				if s[i-1] == 'I' {
//					totalInt += 3
//					break
//				}
//			}
//			totalInt += 5
//		case 'X':
//			if i >= 1 {
//				if s[i-1] == 'I' {
//					totalInt += 8
//					break
//				}
//			}
//			totalInt += 10
//		case 'L':
//			if i >= 1 {
//				if s[i-1] == 'X' {
//					totalInt += 30
//					break
//				}
//			}
//			totalInt += 50
//		case 'C':
//			if i >= 1 {
//				if s[i-1] == 'X' {
//					totalInt += 80
//					break
//				}
//			}
//			totalInt += 100
//		case 'D':
//			if i >= 1 {
//				if s[i-1] == 'C' {
//					totalInt += 300
//					break
//				}
//			}
//			totalInt += 500
//		case 'M':
//			if i >= 1 {
//				if s[i-1] == 'C' {
//					totalInt += 800
//					break
//				}
//			}
//			totalInt += 1000
//		}
//	}
//	return totalInt
//}

func romanToInt(s string) int {
	var (
		total int
		prev  rune
	)
	romanIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for _, val := range s {
		total += romanIntMap[val]
		if romanIntMap[prev] < romanIntMap[val] {
			total -= romanIntMap[prev] * 2
		}
		prev = val
	}
	return total
}
