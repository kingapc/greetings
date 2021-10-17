package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string, lastname string) (string, error) {
	var addition string

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	addition = message + fmt.Sprintf("to your first go code %v", lastname)
	return addition, nil
}

//testing
func Age(year int) int {
	age := 2021 - year
	return age
}
