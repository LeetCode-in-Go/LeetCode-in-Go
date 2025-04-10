package s0001_two_sum

// #Easy #Top_100_Liked_Questions #Top_Interview_Questions #Array #Hash_Table
// #Data_Structure_I_Day_2_Array #Level_1_Day_13_Hashmap #Udemy_Arrays #Top_Interview_150_Hashmap
// #Big_O_Time_O(n)_Space_O(n) #AI_can_be_used_to_solve_the_task
// #2024_01_28_Time_3_ms_(93.85%)_Space_4.2_MB_(58.64%)

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		requiredNum := target - num
		if j, ok := indexMap[requiredNum]; ok {
			return []int{j, i}
		}
		indexMap[num] = i
	}
	return []int{-1, -1}
}
