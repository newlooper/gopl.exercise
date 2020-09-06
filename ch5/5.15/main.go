package main

import (
	"errors"
)

func main() {

}

func max(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("not enough parameters")
	}
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m, nil
}

func min(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("not enough parameters")
	}
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m, nil
}
