package greetings

import "fmt"

func Hello(name string, lastname string) string {
	var addition string

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	addition = message + fmt.Sprintf("to your first go code %v", lastname)
	return addition
}

//testing
func Age(year int) int {
	age := 2021 - year
	return age
}
