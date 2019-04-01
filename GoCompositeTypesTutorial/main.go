/*
Arrays
Slices
Maps
Structs
Nested Structs
*/

package main

import (
	"fmt"
)

// our Person struct
type Person struct {
	name string
	age int
}

// declaring a new `Person`
var myPerson Person

// our Team struct
type Team struct {
	name string
	players [2]Person
}

func main()  {
	// declaring an empty array of string
	var days01 []string
	// declaring an array with elements 
	days02 := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	// print
	fmt.Println(days01)
	fmt.Println(days02)
	fmt.Println(days02[0])
	fmt.Println(days02[5])
	
	// Slices in Go allow you to access a subset of an underlying array’s elements
	// Slices are made up of three things, a pointer, a length, and a capacity
	weekdays := days02[0:5]
	fmt.Println(weekdays)

	// Maps are Go’s representation of hash tables, a data structure that allows you to map one arbitrary data type to another
	youtubeSubscribers := map[string]int {
		"TutorialEdge": 2240,
		"MKBHD":		6580350,
		"Fun Fun Function": 171220,
	}
	fmt.Println(youtubeSubscribers["MKBHD"])

	// In Go, we have this concept of a struct. These structs allow us to create data types that are aggregates of other data types.
	// declaring a new `elliot`
	elliot := Person{name: "Elliot", age: 24}
	elliot.age = 18
	fmt.Println(elliot)

	// Structs are incredibly extensible due to the fact we can create nested structs within our structs
	var myTeam Team
	fmt.Println(myTeam)
	players := [...]Person{Person{name:"TrungLV01"}, Person{name: "TrungLV02"}}
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)
}