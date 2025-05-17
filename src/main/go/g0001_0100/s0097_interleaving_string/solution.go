package s0097_interleaving_string

// #Medium #String #Dynamic_Programming #Top_Interview_150_Multidimensional_DP
// #2025_05_17_Time_0_ms_(100.00%)_Space_4.85_MB_(28.57%)

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	cache := make([][]*bool, len(s1)+1)
	for i := range cache {
		cache[i] = make([]*bool, len(s2)+1)
	}
	return isInterleaveHelper(s1, s2, s3, 0, 0, 0, cache)
}

func isInterleaveHelper(s1, s2, s3 string, i1, i2, i3 int, cache [][]*bool) bool {
	if cache[i1][i2] != nil {
		return *cache[i1][i2]
	}
	if i1 == len(s1) && i2 == len(s2) && i3 == len(s3) {
		return true
	}
	result := false
	if i1 < len(s1) && s1[i1] == s3[i3] {
		result = isInterleaveHelper(s1, s2, s3, i1+1, i2, i3+1, cache)
	}
	if i2 < len(s2) && s2[i2] == s3[i3] {
		result = result || isInterleaveHelper(s1, s2, s3, i1, i2+1, i3+1, cache)
	}
	cache[i1][i2] = &result
	return result
}
