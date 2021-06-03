package helper

import (
	"log"
	"time"
)

// Trace function provides the time spent the method or function under
// which the function is called with the parameter as function name
// call using defer helper.Trace("name"){}
func Trace(msg string) func() {
	start := time.Now()
	log.Printf("* [enter] %s *", msg)
	return func() {
		log.Printf("* [exit] %s took %s *", msg, time.Since(start).String())
	}
}
