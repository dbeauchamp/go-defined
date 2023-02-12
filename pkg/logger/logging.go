package logger

import "log"

const prefix = "[go-defiend]"

func Error(msg string, err error) {
	log.Printf("%v %v: %v \n", prefix, msg, err)
}

func Fatal(msg string, err error) {
	log.Fatalf("%v %v: %v", prefix, msg, err)
}
