package twosum

/*Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].*/

// TwoSum to get indices of 2 integers whose sum is equal to the target
func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, n := range nums {
		numsTofind := target - n

		if hashIndex, ok := hashMap[numsTofind]; ok {
			return []int{hashIndex, i}
		}

		hashMap[n] = i
	}

	return []int{}
}
