package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	} //如果遇见没有给予名字的情况，可以返回错误信息（可自定义）。

	// If a name was received, return a value that embeds the name，如果名字被接受了，会输出一个消息，并且返回一个名字。
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
