package main

import "fmt"

type Person struct {
	Id     int
	Name   string
	Adress string
}
type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (s1 *MySlice) Add(val int) {
	*s1 = append(*s1, val)
}

func (s1 *MySlice) Count() int {
	return len(*s1)
}

func main() {
	var acc Account = Account{
		Id:   1,
		Name: "sksakura89",
		Person: Person{
			Adress: "Piter",
		},
	}
	acc.Owner = Person{
		Id:     2,
		Name:   "alexandra",
		Adress: "Moscow",
	}
	acc.Person.UpdateName("NewName1")
	//fmt.Println(acc.Person.Name)
	acc.Person.SetName("NewName2")
	//fmt.Println(acc.Person.Name)

	mySlice := MySlice([]int{1, 2, 3})
	mySlice.Add(1)
	mySlice.Add(2)
	mySlice.Add(33)
	fmt.Println(mySlice, mySlice.Count())
}
