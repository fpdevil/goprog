package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	plength = 16 // password length from user
	t       = time.Now().UTC().UnixNano()
)

// characters for password as raw text
const (
	letters  = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	numbers  = `0123456789`
	splchars = `-=~!@#$%^&*()_+[]\{}|;':",./<>?`
)

func init() {
	flag.IntVar(&plength, "l", plength, "specify password length between 8-128")
	flag.Parse()
}

func main() {
	fmt.Println("/// Password Generator ///")
	fmt.Println()

	if plength < 8 || plength > 128 {
		flag.Usage()
		return
	}

	rand.Seed(t)
	pwd := genRandomPwd(plength)

	for v := range pwd {
		fmt.Printf("%c", v)
	}

	fmt.Println()
}

func genRandomPwd(l int) (out chan rune) {
	var (
		runeLetters = []rune(letters)
		runeNums    = []rune(numbers)
		runeSplcs   = []rune(splchars)
	)

	out = make(chan rune, l)

	// this should be placed after the initialization of the out
	// channel as otherwise, we wiil be greeted with the following
	// error
	// panic: close of nil channel
	//
	defer close(out)

	// using select case to increase probability of one character set over another
	for {
		// generate a random number of length equal to runeLetters
		rl := rand.Intn(len(runeLetters))
		rn := rand.Intn(len(runeNums))
		rs := rand.Intn(len(runeSplcs))

		// dividing the default 16 length allocating more values
		// to letters and remaining minor percentages to the
		// numbers and special characters
		select {
		case out <- runeLetters[rl]:
		case out <- runeLetters[rl]:
		case out <- runeLetters[rl]:
		case out <- runeLetters[rl]:
		case out <- runeLetters[rl]:
		case out <- runeLetters[rl]:
		case out <- runeLetters[rl]:
		case out <- runeLetters[rl]:
		case out <- runeLetters[rl]:
		case out <- runeLetters[rl]:

		case out <- runeNums[rn]:
		case out <- runeNums[rn]:
		case out <- runeNums[rn]:

		case out <- runeSplcs[rs]:

		default:
			// either close here or defer close as earlier
			// close(out)
			return
		}
	}
}
