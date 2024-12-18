package main

func QuickSelect(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k-1)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	pivot := left + (right-left)/2
	i := partition(nums, left, right, pivot)
	if i == k {
		return nums[i]
	} else if i < k {
		return quickSelect(nums, i+1, right, k)
	} else {
		return quickSelect(nums, left, i-1, k)
	}
}

func partition(nums []int, left, right, pivot int) int {
	pivotValue := nums[pivot]
	swap(nums, pivot, right)
	pivotIndex := left
	for i := left; i < right; i++ {
		if nums[i] > pivotValue {
			swap(nums, i, pivotIndex)
			pivotIndex++
		}
	}
	swap(nums, pivotIndex, right)
	return pivotIndex
}

func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
