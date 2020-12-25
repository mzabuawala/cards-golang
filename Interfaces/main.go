package main

import "fmt"

type bot interface {
	getGreetings() string
}

type hindiBot struct{}
type gujaratiBot struct{}
type englishBot struct{}

func main() {
	hb := hindiBot{}
	gb := gujaratiBot{}
	en := englishBot{}

	printGreeting(hb)
	printGreeting(gb)
	printGreeting(en)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}

func (hindiBot) getGreetings() string {
	return "केसे हो?"
}

func (gujaratiBot) getGreetings() string {
	return "કેમ છો?"
}

func (englishBot) getGreetings() string {
	return "How are you?"
}
