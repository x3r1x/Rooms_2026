package recovery

import (
	"log"
	"runtime/debug"
)

func Recover() {
	if r := recover(); r != nil {
		log.Printf("PANIC: %v\nStack:\n%s", r, debug.Stack())
	}
}
