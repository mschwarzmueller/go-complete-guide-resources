package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
}

func main() {
	greet("Nice to meet you!")
	greet("How are you?")
	slowGreet("How ... are ... you ...?")
	greet("I hope you're liking the course!")
}
