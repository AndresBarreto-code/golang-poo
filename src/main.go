package main

import (
	company "CursoGolangPlatzi/ObjectOriented/Classes/company"
	store "CursoGolangPlatzi/ObjectOriented/Classes/store"
	"fmt"
)

func main() {
	e := store.Employee{
		Id:   2,
		Name: "Name0",
	}
	fmt.Printf("%v\n", e)
	e.SetId(1)
	e.SetName("Name1")
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", e.GetId())
	fmt.Printf("%v\n", e.GetName())
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// Company
	fe := company.FullTimeEmployee{}
	fe.Age = 25
	fe.Name = "Name FE"
	fe.Id = 2
	fe.EndDate = "tomorrow"
	fmt.Printf("%v\n", fe)
	pe := company.PartTimeEmployee{}
	fmt.Printf("%v\n", fe)
	company.GetMessage(pe)
	company.GetMessage(fe)
}
