package main

import (
	"fmt"
	"github.com/balamuru/pkgdemo2/greetings"
)

func main() {
	message := greetings.Hello("VB2")
	fmt.Println(message)
}
