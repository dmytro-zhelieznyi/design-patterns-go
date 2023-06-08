package human

import "fmt"

type Student struct {
}

func (e *Student) SayHello() {
	fmt.Printf("Hello, I'm a student.\n")
}
