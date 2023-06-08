package main

import (
	"fmt"
)

func main() {
	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "hello").
		AddChildFluent("li", "world")
	fmt.Println(b.String())

	pb := NewPersonBuilder()
	person := pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(123000).
		Build()
	fmt.Println(*person)

	SendEmail(func(b *EmailBuilder) {
		b.From("from@example.com").
			To("to@example.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}
