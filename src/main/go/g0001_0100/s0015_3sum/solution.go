package s0015_3sum

// #Medium #Top_100_Liked_Questions #Top_Interview_Questions #Array #Sorting #Two_Pointers
// #Data_Structure_II_Day_1_Array #Algorithm_II_Day_3_Two_Pointers #Udemy_Two_Pointers
// #Big_O_Time_O(n*log(n))_Space_O(n^2) #2024_03_07_Time_37_ms_(90.03%)_Space_7.9_MB_(46.68%)

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i, a := range nums {
		if a > 0 {
			break
		}
		if i > 0 && a == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := a + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{a, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	// Return the final result
	return res
}
