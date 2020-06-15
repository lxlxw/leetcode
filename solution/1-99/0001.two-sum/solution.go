package leetcode

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for k, v := range nums {
		if _, ok := numsMap[target-v]; ok {
			return []int{numsMap[target-v], k}
		}
		numsMap[v] = k
	}
	return []int{}
}
