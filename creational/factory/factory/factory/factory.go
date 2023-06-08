package factory

import (
	"design-patterns-go/creational/factory/factory/human"
	"log"
	"sync"
)

var once sync.Once
var employeeFactoryInstance *EmployeeFactory
var studentFactoryInstance *StudentFactory

type Factory interface {
	Create() human.Human
}

func GetHumanFactory(factory string) Factory {
	switch factory {
	case "employee":
		return getEmployeeFactory()
	case "student":
		return getStudentFactory()
	default:
		log.Fatalf("Wrong factory type [%s]", factory)
		return nil
	}
}

type EmployeeFactory struct{}

func getEmployeeFactory() *EmployeeFactory {
	once.Do(func() {
		if employeeFactoryInstance == nil {
			employeeFactoryInstance = &EmployeeFactory{}
		}
	})
	return employeeFactoryInstance
}

func (f *EmployeeFactory) Create() human.Human {
	return &human.Employee{}
}

type StudentFactory struct{}

func getStudentFactory() *StudentFactory {
	once.Do(func() {
		studentFactoryInstance = &StudentFactory{}
	})
	return studentFactoryInstance
}

func (f *StudentFactory) Create() human.Human {
	return &human.Student{}
}
