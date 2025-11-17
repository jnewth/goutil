package goutil

import (
	"fmt"
	"log"
)

func Verify(test bool, format string, va ...any) {
	if !test {
		log.Fatalf(format, va...)
	}
}

func Verboseln(report bool, msg string) {
	if report {
		fmt.Println(msg)
	}
}
func Verbosef(report bool, format string, va ...any) {
	if report {
		fmt.Printf(format, va...)
	}
}
