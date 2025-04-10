package s0763_partition_labels

// #Medium #String #Hash_Table #Greedy #Two_Pointers #Data_Structure_II_Day_7_String
// #Big_O_Time_O(n)_Space_O(1) #2024_03_10_Time_0_ms_(100.00%)_Space_2.3_MB_(63.50%)

func partitionLabels(s string) []int {
	letters := []rune(s)
	result := make([]int, 0)
	position := make([]int, 26)

	for i, letter := range letters {
		position[letter-'a'] = i
	}

	i := 0
	prev := -1
	max := 0

	for i < len(letters) {
		if position[letters[i]-'a'] > max {
			max = position[letters[i]-'a']
		}
		if i == max {
			result = append(result, i-prev)
			prev = i
		}
		i++
	}

	return result
}
