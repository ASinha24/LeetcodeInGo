package intersectionoftwoarray

/*
Given two integer arrays nums1 and nums2, return an array of their
intersection
. Each element in the result must be unique and you may return the result in any order.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
*/
func IntersectionOfArrays(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]struct{})

	for _, n := range nums1 {
		hashMap[n] = struct{}{}
	}

	resMap := make(map[int]struct{})
	for _, n := range nums2 {
		if _, ok := hashMap[n]; ok {
			resMap[n] = struct{}{}
		}
	}

	intersection := []int{}
	for k, _ := range resMap {
		intersection = append(intersection, k)
	}

	return intersection
}
