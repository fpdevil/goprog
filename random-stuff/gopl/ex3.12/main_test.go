package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s1, s2 string
		needed bool
	}{
		{"reaps", "pears", true},
		{"abcabc", "cacaca", false},
		{"oranges", "potatoes", true},
	}

	for _, test := range tests {
		received := isAnagram(test.s1, test.s2)
		if received != test.needed {
			t.Errorf("isAnagram(%q, %q), received %v and needed %v", test.s1, test.s2, received, test.needed)
		}
	}
}
