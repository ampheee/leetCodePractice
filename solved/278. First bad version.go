package solved

func isBadVersion(version int) bool {
	var param int
	if param == 0 {
		return false
	}
	return false
}

func firstBadVersion(n int) int {
	var (
		left  = 0
		right = n
	)
	for left <= right {
		mid := (left + right) / 2
		if isBadVersion(mid) == true && isBadVersion(mid-1) == true {
			right = mid - 1
		} else if isBadVersion(mid) == false {
			left = mid + 1
		} else {
			return mid
		}
	}
	return n
}
