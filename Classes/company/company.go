package mypackage

import "fmt"

type PrintInfo interface {
	getMessage() string
}

func (fe FullTimeEmployee) getMessage() string {
	return "Full time employee"
}

func (pe PartTimeEmployee) getMessage() string {
	return "Part time employee"
}

func GetMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	id int
	Id int
}

func (e *Employee) SetId(id int) bool {
	e.id = id
	return true
}
func (e *Employee) GetId() int {
	return e.id
}

type FullTimeEmployee struct {
	Person
	Employee
	EndDate string
}

type PartTimeEmployee struct {
	Person
	Employee
	TaxRate int
}
