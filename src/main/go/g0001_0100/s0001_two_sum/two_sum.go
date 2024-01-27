package g0001_0100.s0001_two_sum

func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, v := range nums {
		j, ok := lookup[v]
		lookup[target-v] = i
		if ok {
			return []int{j, i}
		}
	}
	return nil
}
