package human

import "fmt"

type Employee struct {
}

func (e *Employee) SayHello() {
	fmt.Printf("Hello, I'm a employee.\n")
}
