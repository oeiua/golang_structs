package main

import (
	"fmt"

	"./zoo"
)

/*
0. Declaration
1. Different types
3. Creating instances, instantiation
4. Zero values
5. Consecutive fields with same type
6. User defined type. Structs are the only way to create concrete user-defined types in Golang.
7. Struct Visibility
8. Fields Visibility
9. Embedding
10. Recursive (almost)
11. Anonymous/unnamed structs
12. Structural comparison
13. methods or pointer and value receivers
14. Tagging
15. Implicit dereferencing
16. Assignment
*/

var predatorHead = zoo.Head{Fangs: true}
var herbivorousHead = zoo.Head{Fangs: false}

var topRightLeg = zoo.Leg{Side: "right", Position: "top"}
var topLeftLeg = zoo.Leg{Side: "left", Position: "top"}
var bottomRightLeg = zoo.Leg{Side: "right", Position: "bottom"}
var bottomLeftLeg = zoo.Leg{Side: "left", Position: "bottom"}

var Wolf = zoo.Animal{
	Name: "Wolfie",
	Type: "wolf",
	Legs: []zoo.Leg{topRightLeg, topLeftLeg, bottomRightLeg, bottomLeftLeg},
	Head: predatorHead,
}

var Sheep = zoo.Animal{
	Name: "Dolly",
	Type: "sheep",
	Legs: []zoo.Leg{topRightLeg, topLeftLeg, bottomRightLeg, bottomLeftLeg},
	Head: herbivorousHead,
}

var predatorCage = zoo.Cage{
	Animals: []zoo.Animal{Wolf},
}

var herbivorousCage = zoo.Cage{
	Animals: []zoo.Animal{Sheep},
}

var MyZoo = zoo.Zoo{
	Cages: []*zoo.Cage{&herbivorousCage, &predatorCage},
}

func main() {
	fmt.Printf("There are %v cages in zoo:\n", len(MyZoo.Cages))
	for count, cage := range MyZoo.Cages {

		fmt.Printf("Here %v animal in cage number %v\n", len(cage.Animals), count+1)

		for _, animal := range cage.Animals {
			fmt.Printf("This is %v named %v\n", animal.Type, animal.Name)
			if animal.Head.Fangs {
				fmt.Printf("%v is a predator. He has fangs\n", animal.Name)
			} else {
				fmt.Printf("%v is a herbivorous.\n", animal.Name)
			}
			fmt.Printf("%v has:\n", animal.Name)
			for _, leg := range animal.Legs {
				fmt.Printf(" - %v %v leg\n", leg.Side, leg.Position)
			}
		}
	}
}
