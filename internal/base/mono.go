package base

func Mono(nums []int) bool {
	up, down := true, true

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			up = false
		}
		if nums[i] > nums[i-1] {
			down = false
		}
		if !up && !down {
			return false
		}
	}
	return true
}