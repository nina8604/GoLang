package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type personByAge []user

func (pa personByAge) Len() int {
	return len(pa)
}
func (pa personByAge) Less(i, j int) bool {
	return pa[i].Age < pa[j].Age
}
func (pa personByAge) Swap(i, j int) {
	pa[i], pa[j] = pa[j], pa[i]
}

type personByLast []user

func (pl personByLast) Len() int {
	return len(pl)
}
func (pl personByLast) Less(i, j int) bool {
	return pl[i].Last < pl[j].Last
}
func (pl personByLast) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		sort.Strings(user.Sayings)
		for _, v := range user.Sayings {
			fmt.Println(v)
		}
	}

	// your code goes here

	sort.Sort(personByAge(users))
	fmt.Println("--------- sort by age ------ ")
	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		for _, v := range user.Sayings {
			fmt.Println("\t", v)
		}
	}

	sort.Sort(personByLast(users))
	fmt.Println("--------- sort by last name ------ ")
	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		for _, v := range user.Sayings {
			fmt.Println("\t", v)
		}
	}
}
