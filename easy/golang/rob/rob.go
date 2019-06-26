package rob

import "math"

func rob(nums []int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}
	if count == 1 {
		return nums[0]
	}
	
	one, two := nums[0], nums[1]
	for i := 2; i < count; i ++ {
		tmp := one
		if two > one {
			one = two
		}
		
		two = tmp + nums[i]
	}
	
	return int(math.Max(float64(one), float64(two)))
}