package s0763_partition_labels

// #Medium #Top_100_Liked_Questions #String #Hash_Table #Greedy #Two_Pointers
// #Data_Structure_II_Day_7_String #Big_O_Time_O(n)_Space_O(1)

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
