package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string, lastname string) (string, error) {
	var addition string

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	addition = message + fmt.Sprintf("to your first go code %v", lastname)
	return addition, nil
}

//testing
func Age(year int) int {
	age := 2021 - year
	return age
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
