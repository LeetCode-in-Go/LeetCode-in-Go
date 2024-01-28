package s0001_two_sum

func twoSum(numbers []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range numbers {
		requiredNum := target - num
		if j, ok := indexMap[requiredNum]; ok {
			return []int{j, i}
		}
		indexMap[num] = i
	}
	return []int{-1, -1}
}
