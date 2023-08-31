package main

import "log"

type all interface{}

func PrintType() {
	var a all = "hello"
	switch a.(type) {
	case int:
		log.Printf("%d is int", a)
	case string:
		log.Printf("%s is string", a)
	}
}
