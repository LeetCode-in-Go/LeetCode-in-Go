package s0066_plus_one

// #Easy #Top_Interview_Questions #Array #Math #Programming_Skills_II_Day_3 #Udemy_Arrays
// #Top_Interview_150_Math #2025_05_14_Time_0_ms_(100.00%)_Space_3.99_MB_(76.14%)

func plusOne(digits []int) []int {
	num := 1
	carry := 0
	var sum int
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			sum = digits[i] + carry + num
		} else {
			sum = digits[i] + carry
		}
		carry = sum / 10
		digits[i] = sum % 10
	}
	if carry != 0 {
		ans := make([]int, len(digits)+1)
		ans[0] = carry
		copy(ans[1:], digits)
		return ans
	}
	return digits
}
