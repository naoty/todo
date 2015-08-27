package todoutil

import (
	"strconv"
)

func Atois(words []string) ([]int, error) {
	var nums []int
	for _, word := range words {
		num, err := strconv.Atoi(word)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func ContainsInt(xs []int, n int) bool {
	for _, x := range xs {
		if x == n {
			return true
		}
	}
	return false
}
