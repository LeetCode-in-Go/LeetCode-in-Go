package s0763_partition_labels

// #Medium #Top_100_Liked_Questions #String #Hash_Table #Greedy #Two_Pointers
// #Data_Structure_II_Day_7_String #Big_O_Time_O(n)_Space_O(1)

func dailyTemperatures(temperatures []int) []int {
    sol := make([]int, len(temperatures))
    sol[len(temperatures)-1] = 0

    for i := len(sol) - 2; i >= 0; i-- {
        j := i + 1
        for j < len(sol) {
            if temperatures[i] < temperatures[j] {
                sol[i] = j - i
                break
            } else {
                if sol[j] == 0 {
                    break
                }
                j += sol[j]
            }
        }
    }
    return sol
}
