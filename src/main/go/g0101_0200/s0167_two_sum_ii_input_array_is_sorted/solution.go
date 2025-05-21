package s0167_two_sum_ii_input_array_is_sorted

// #Medium #Array #Binary_Search #Two_Pointers #Algorithm_I_Day_3_Two_Pointers
// #Binary_Search_I_Day_7 #Top_Interview_150_Two_Pointers
// #2025_05_21_Time_0_ms_(100.00%)_Space_8.04_MB_(18.55%)

func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	i := 0
	j := len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			res[0] = i + 1
			res[1] = j + 1
			return res
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return res
}
