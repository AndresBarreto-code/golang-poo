package mypackage

type Employee struct {
	id   int
	name string
	Id   int
	Name string
}

func (e *Employee) SetName(name string) bool {
	e.name = name
	return true
}
func (e *Employee) SetId(id int) bool {
	e.id = id
	return true
}
func (e *Employee) GetName() string {
	return e.name
}
func (e *Employee) GetId() int {
	return e.id
}
