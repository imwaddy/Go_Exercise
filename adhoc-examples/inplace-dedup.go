package main

func dedup() []int {
	nums := []int{0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 4, 5, 5, 5}
	i := 1
	j := i + 1
	for j <= len(nums)-1 {
		if nums[j-1] != nums[j] {
			nums[i] = nums[j]
			i++
			j++
		} else {
			j++
		}
	}
	return nums[0:i]
}
