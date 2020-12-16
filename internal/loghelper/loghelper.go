package loghelper

import (
	"fmt"
	"log"
)

// Log helps logging with 2 input - one at the time of function call, and
// another at later point.
func Log(before, after string) func() {
	prefix := "======== "

	beforeLog := fmt.Sprintf("%s%s", prefix, before)
	afterLog := fmt.Sprintf("%s%s", prefix, after)

	log.Printf(beforeLog)
	return func() { log.Printf(afterLog) }
}
