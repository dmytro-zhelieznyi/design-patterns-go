package main

import "design-patterns-go/creational/factory/factory/factory"

func main() {
	employee := factory.GetHumanFactory("employee").Create()
	student := factory.GetHumanFactory("student").Create()
	employee.SayHello()
	student.SayHello()
}
