package s0152_maximum_product_subarray

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Dynamic_Programming
// #Dynamic_Programming_I_Day_6 #Level_2_Day_13_Dynamic_Programming #Udemy_Dynamic_Programming
// #Big_O_Time_O(N)_Space_O(1) #2025_05_08_Time_0_ms_(100.00%)_Space_5.25_MB_(91.63%)

func maxProduct(nums []int) int {
	res := findMax(nums...)
	min, max := 1, 1

	for _, n := range nums {
		if n == 0 {
			min, max = 1, 1
			continue
		}
		buffer := max * n
		max = findMax(n*max, n*min, n)
		min = findMin(buffer, n*min, n)

		res = findMax(res, max)
	}
	return res
}

func findMax(nums ...int) int {
	highest := nums[0]
	for _, n := range nums {
		if n > highest {
			highest = n
		}
	}
	return highest
}

func findMin(nums ...int) int {
	lowest := nums[0]
	for _, n := range nums {
		if n < lowest {
			lowest = n
		}
	}
	return lowest
}
