package main

func dedupCount() [][]int {
	nums := []int{0, 1, 1, 1, 1, 2, 2, 2, 4, 5}
	if len(nums) == 0 {
		return [][]int{}
	}

	result := make([][]int, 0)
	start := 0

	for end := 1; end < len(nums); end++ {
		if nums[end] != nums[start] {
			result = append(result, []int{nums[start], end - start})
			start = end
		}
	}

	result = append(result, []int{nums[start], len(nums) - start})

	return result
}
