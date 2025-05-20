package s0135_candy

// #Hard #Array #Greedy #Top_Interview_150_Array/String
// #2025_05_20_Time_0_ms_(100.00%)_Space_8.28_MB_(51.51%)

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			candies[i+1] = candies[i] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && candies[i-1] < candies[i]+1 {
			candies[i-1] = candies[i] + 1
		}
	}
	total := 0
	for _, candy := range candies {
		total += candy
	}
	return total
}
