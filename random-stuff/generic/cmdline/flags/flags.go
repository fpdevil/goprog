package main

import (
	"flag"
	"fmt"
)

// Config is a struct to hold command line flags
type Config struct {
	subject  string
	isGood   bool
	howGood  int
	getCount GetCount
}

// Setup initializes a config from flags that are passed in
func (c *Config) Setup() {
	flag.StringVar(&c.subject, "subject", "", "subject is a string, with default = empty")
	flag.StringVar(&c.subject, "s", "", "subject is a string, with default = empty (shorthand)")
	flag.BoolVar(&c.isGood, "isgood", false, "is it any good?")
	flag.IntVar(&c.howGood, "howgood", 10, "how good is it out of 10?")
	flag.Var(&c.getCount, "c", "comma delimited list of integers")
}

// GetMessage function uses all of the variables from Config and
// returns a sentence or a statement
func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isGood {
		msg += " is excellent"
	} else {
		msg += " is Average!"
	}

	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Lets count the ways %v", msg, c.howGood, c.getCount)
	return msg
}
