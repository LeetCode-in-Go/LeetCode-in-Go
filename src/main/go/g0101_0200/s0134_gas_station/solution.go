package s0134_gas_station

// #Medium #Top_Interview_Questions #Array #Greedy #Top_Interview_150_Array/String
// #2025_05_20_Time_0_ms_(100.00%)_Space_11.11_MB_(34.51%)

func canCompleteCircuit(gas []int, cost []int) int {
	index := 0
	total := 0
	current := 0
	for i := 0; i < len(gas); i++ {
		balance := gas[i] - cost[i]
		total += balance
		current += balance
		if current < 0 {
			index = i + 1
			current = 0
		}
	}
	if total >= 0 {
		return index
	}
	return -1
}
