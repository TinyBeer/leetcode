package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var ans int
	var df [2]int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			df[0] = 0
			df[1] = 0
			continue
		} else if nums[i] > 0 {
			df[0] *= nums[i]
			df[1] *= nums[i]
		} else {
			df[0], df[1] = df[1]*nums[i], df[0]*nums[i]
		}
		if df[0] == 0 && nums[i] > 0 {
			df[0] = nums[i]
		}

		if df[1] == 0 && nums[i] < 0 {
			df[1] = nums[i]
		}

		ans = max(ans, df[0])
	}
	return ans
}
