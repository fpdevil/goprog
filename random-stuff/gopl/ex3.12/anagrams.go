package anagrams

func isAnagram(s1, s2 string) bool {
	s1vals := make(map[rune]int)
	for _, c := range s1 {
		s1vals[c]++
	}

	s2vals := make(map[rune]int)
	for _, c := range s2 {
		s2vals[c]++
	}

	for k, v := range s1vals {
		if s2vals[k] != v {
			return false
		}
	}

	for k, v := range s2vals {
		if s1vals[k] != v {
			return false
		}
	}

	return true
}
