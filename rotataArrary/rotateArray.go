package rotataarrary

// reverseArray
func ReverseArray(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}

}

// RotateArrayFromRight
func RotateArrayFromRight(nums []int, k int) []int {
	n := len(nums)
	d := k % n

	ReverseArray(nums, n-d, n-1)
	ReverseArray(nums, 0, n-d-1)
	ReverseArray(nums, 0, n-1)

	return nums
}

// RotateArrayFromLeft
func RotateArrayFromLeft(nums []int, k int) []int {
	n := len(nums)
	d := k % n

	ReverseArray(nums, 0, d-1)
	ReverseArray(nums, d, n-1)
	ReverseArray(nums, 0, n-1)

	return nums
}
