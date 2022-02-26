package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Ola!, %v. Have a nice day", name)
	return message
}
