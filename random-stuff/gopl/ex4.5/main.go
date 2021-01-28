package main

import "fmt"

func removedups(s []string) []string {
	i := 0
	for _, j := range s {
		if s[i] == j {
			continue
		}
		i++
		s[i] = j
	}
	return s[:i+1]
}

func main() {
	s := []string{"m", "a", "m", "m", "a", "l"}
	s = removedups(s)
	fmt.Printf("removed duplicates %s\n", s)
}
