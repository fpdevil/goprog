package main

import (
	"fmt"
	"strconv"
	"strings"
)

// GetCount is a custom type of []int to read a flag into
type GetCount []int

func (c *GetCount) String() string {
	result := ""
	for _, v := range *c {
		if len(result) > 0 {
			result += " ... "
		}
	result += fmt.Sprintf("%v", v)
	}
	return result
}

// Set function is used by the flag package
func (c *GetCount) Set(value string) error {
	values := strings.Split(value, ",")
	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*c = append(*c, i)
	}
	return nil
}
