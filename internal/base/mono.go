package base

func Mono(nums []int) bool {
	increasing := false
	started := false
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		if nums[i] < nums[i-1] {
			if increasing && started {
				return false
			}
			started = true
		}

		if nums[i] > nums[i-1] {
			if !increasing && started {
				return false
			}
			increasing = true
			started = true
		}
	}

	return true
}