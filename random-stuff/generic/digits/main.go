package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	msg = `
	usage: %s [-b | --bar] <whole-number>
	-b --bar draw an underbar and an overbar
	`
)

// digits represents a 10 X 7 (rows X columns) matrix consisting of string
// values to print as digits
var digits = [][]string{
	{
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  ",
	},
	{
		" 1 ",
		" 1 ",
		" 1 ",
		" 1 ",
		" 1 ",
		" 1 ",
		" 1 ",
	},
	{
		"2222222",
		"      2",
		"      2",
		"2222222",
		"2      ",
		"2      ",
		"2222222",
	},
	{
		"3333333",
		"      3",
		"      3",
		"3333333",
		"      3",
		"      3",
		"3333333",
	},
	{
		"4     4",
		"4     4",
		"4     4",
		"4444444",
		"      4",
		"      4",
		"      4",
	},
	{
		"5555555",
		"5      ",
		"5      ",
		"5555555",
		"      5",
		"      5",
		"5555555",
	},
	{
		"6666666",
		"6      ",
		"6      ",
		"6666666",
		"6     6",
		"6     6",
		"6666666",
	},
	{
		"7777777",
		"      7",
		"     7 ",
		"    7  ",
		"   7   ",
		"  7    ",
		" 7     ",
	},
	{
		"8888888",
		"8     8",
		"8     8",
		"8888888",
		"8     8",
		"8     8",
		"8888888",
	},
	{
		"9999999",
		"9     9",
		"9     9",
		"9999999",
		"      9",
		"      9",
		"      9",
	},
}

func main() {
	args := os.Args
	if len(args) == 1 || args[1] == "-h" || args[1] == "--help" ||
		(len(args) == 2 && (args[1] == "-b" || args[1] == "--bar")) {
		fmt.Printf(msg+"\n", filepath.Base(os.Args[0]))
		return
	}

	var (
		stringOfDigits string
		bar            bool
	)

	if args[1] == "-b" || args[1] == "--bar" {
		bar = true
		stringOfDigits = args[2]
	} else {
		// do not print the bar, but only string integers
		stringOfDigits = args[1]
	}

	// for row ranging from 0 to 10
	for row := range digits[0] {
		line := ""
		// for col ranging from 0 to 7
		for col := range stringOfDigits {
			// fmt.Println(col, " -> ", stringOfDigits[col]-'0')
			// we will  use ascii  difference between  current value  and 0
			// (48) as the ascii difference between the number and 0 is the
			// number itself.
			// 0  :  48
			// 1  :  49
			// 2  :  50
			// 3  :  51
			// 4  :  52
			// 5  :  53
			// 6  :  54
			// 7  :  55
			// 8  :  56
			// 9  :  57
			digit := stringOfDigits[col] - '0'
			if 0 <= digit && digit <= 9 {
				line += digits[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number specified")
			}
		}
		if bar && row == 0 {
			fmt.Println(strings.Repeat("*", len(line)))
		}
		fmt.Println(line)
		if bar && row == len(digits[0])-1 {
			fmt.Println(strings.Repeat("*", len(line)))
		}
	}
}
