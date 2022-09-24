package helpers

import "log"

func FailOnError(err error, msg string, args ...any) {
	if err != nil {
		log.Printf(msg, args)
		log.Fatal(err)
	}
}
