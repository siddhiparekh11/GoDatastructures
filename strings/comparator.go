package strings


import "sort"

type Person struct {
	Name string
	Age int
}


func Compare(persons []Person) []Person {

	sort.Slice(persons, func(i,j int) bool {

		return persons[i].Age > persons[j].Age
		})
	return persons
} 