package s0068_text_justification

// #Hard #Array #String #Simulation #Top_Interview_150_Array/String
// #2025_05_14_Time_0_ms_(100.00%)_Space_4.14_MB_(52.84%)

func fullJustify(words []string, maxWidth int) []string {
	output := make([]string, 0, (len(words)+1)/(1+maxWidth/7))
	sb := make([]byte, 0, maxWidth)
	lineTotal := 0
	numWordsOnLine := 0
	startWord := 0

	for i := 0; i < len(words)-1; i++ {
		lineTotal += len(words[i])
		numWordsOnLine++
		if lineTotal+numWordsOnLine+len(words[i+1]) > maxWidth {
			if numWordsOnLine == 1 {
				sb = append(sb, words[i]...)
				for lineTotal < maxWidth {
					sb = append(sb, ' ')
					lineTotal++
				}
			} else {
				extraSp := (maxWidth - lineTotal) % (numWordsOnLine - 1)
				for j := startWord; j < startWord+numWordsOnLine-1; j++ {
					sb = append(sb, words[j]...)
					if extraSp > 0 {
						sb = append(sb, ' ')
						extraSp--
					}
					for k := 0; k < (maxWidth-lineTotal)/(numWordsOnLine-1); k++ {
						sb = append(sb, ' ')
					}
				}
				sb = append(sb, words[startWord+numWordsOnLine-1]...)
			}
			output = append(output, string(sb))
			startWord = i + 1
			numWordsOnLine = 0
			lineTotal = 0
			sb = make([]byte, 0, maxWidth)
		}
	}

	lineTotal = 0
	for i := startWord; i < len(words); i++ {
		lineTotal += len(words[i])
		sb = append(sb, words[i]...)
		if lineTotal < maxWidth {
			sb = append(sb, ' ')
			lineTotal++
		}
	}
	for lineTotal < maxWidth {
		sb = append(sb, ' ')
		lineTotal++
	}
	output = append(output, string(sb))
	return output
}
