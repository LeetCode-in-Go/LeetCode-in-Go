package s0202_happy_number

// #Easy #Top_Interview_Questions #Hash_Table #Math #Two_Pointers #Algorithm_II_Day_21_Others
// #Programming_Skills_I_Day_4_Loop #Level_2_Day_1_Implementation/Simulation
// #Top_Interview_150_Hashmap #2025_05_22_Time_0_ms_(100.00%)_Space_3.94_MB_(86.55%)

func isHappy(n int) bool {
	if n == 1 || n == 7 {
		return true
	}
	if n > 1 && n < 10 {
		return false
	}
	sum := 0
	a := n
	for a != 0 {
		rem := a % 10
		sum += rem * rem
		a /= 10
	}
	if sum != 1 {
		return isHappy(sum)
	}
	return true
}
