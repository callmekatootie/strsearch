package strsearch

/*
    Public function that takes in the text and the pattern to find in that text
    Returns -1 if pattern can't be found, else returns the location (zero based)
    where pattern can be found
    Uses Boyer Moore Horspool algorithm
*/
func Find(text, pattern string) int {
	m := len(text)
	n := len(pattern)

	if n == 0 {
		return 0
	} else if m == n && text == pattern {
		return 0
	}

	//Prepare the table
	jumps := make(map[rune]int)

	for i, char := range pattern {
		if i < n-1 {
			jumps[char] = n - i - 1
		} else {
            // Mismatch of the last character results in a jump
            // equal to the length of the string
            jumps[char] = n
        }
	}

	j := n - 1

	for j < m {
		k := j
        i := n -1

		for i >= 0 {
			if text[k] != pattern[i] {
				if jump, ok := jumps[rune(text[j])]; ok {
                    // Character exists in the pattern
                    // Jump based on the table created earlier
					j += jump
				} else {
                    // Character does not exist in the pattern.
                    // Jump the length of the pattern.
					j += n
				}

                break
			}

			k--
            i--
		}

		if i < 0 {
            // The pattern string has exhausted - We have a match
			return k + 1
		}
	}

	return -1
}
