package main

import (
	"fmt"
	"reflect"
	"sort"
)

type WorkHistory struct {
	companyName string
	yearOfExit int
}

type WorkHistoryList []WorkHistory

func (wh WorkHistoryList) Len() int { return len(wh) }

func (wh WorkHistoryList) Less(i, j int) bool {
	return wh[i].yearOfExit < wh[j].yearOfExit
}

func (wh WorkHistoryList) Swap(i, j int) { wh[i], wh[j] = wh[j], wh[i]}


type Person struct {
	firstName string
	lastName string
	jobHistory []WorkHistory
}

type ListOfPersons []Person

func (p ListOfPersons) Len() int { return len(p) }

func (p ListOfPersons) Less(i, j int) bool {
	if p[i].firstName == p[j].firstName {
		return p[i].lastName < p[j].lastName
	}

	return p[i].firstName < p[j].firstName
}

func (p ListOfPersons) Swap(i, j int) { p[i], p[j] = p[j], p[i]}



var personList1 = []Person{
	{"John", "Doe", []WorkHistory{
		{"Thoughtworks", 2008},
		{"Google", 1998},
		{"Amazon", 2005},
	}},
	{"John", "Snow", []WorkHistory{
		{"Infosys", 2016},
		{"Cognizant", 1995},
		{"Wipro", 2003},
		{"GS Labs", 2009},
	}},
	{"Jane", "Doe", []WorkHistory{
		{"Apple", 1999},
		{"Microsoft", 1995},
		{"Tesla", 2003},
	}},
}


var personList2 = []Person{
	{"John", "Doe", []WorkHistory{
		{"Thoughtworks", 2008},
		{"Google", 1998},
		{"Amazon", 2005},
	}},
	{"John", "Snow", []WorkHistory{
		{"Infosys", 2016},
		{"Cognizant", 1995},
		{"Wipro", 2003},
		{"GS Labs", 2009},
	}},
	{"Jane", "Doe", []WorkHistory{
		{"Apple", 1999},
		{"Tesla", 2003},
		{"Microsoft", 1995},
	}},
}





func main() {

	fmt.Println("PersonList1 == personList2", reflect.DeepEqual(personList1, personList2))

	sortOne(personList1)
	sortTwo(personList2)

	fmt.Println("PersonList1 == personList2", reflect.DeepEqual(personList1, personList2))

}

func sortOne(persons []Person) {
	fmt.Println(persons)
	sort.Sort(ListOfPersons(persons))
	for _, person := range persons {
		sort.Sort(WorkHistoryList(person.jobHistory))
	}
	fmt.Println(persons)
}


func sortTwo(persons []Person) {
	fmt.Println(persons)

	sort.Slice(persons, func(i, j int) bool {
		if persons[i].firstName == persons[j].firstName {
			return persons[i].lastName < persons[j].lastName
		}

		return persons[i].firstName < persons[j].firstName
	})

	sort.Sort(ListOfPersons(persons))
	for _, person := range persons {
		sort.Slice(person.jobHistory, func(i, j int) bool {
			return person.jobHistory[i].yearOfExit < person.jobHistory[j].yearOfExit
		})
	}
	fmt.Println(persons)
}
