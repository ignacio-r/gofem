package main

import (
	"fmt"

	"fem.com/go/server/data"
)

func main() {

	max := data.Instructor{
		Id:       3,
		Lastname: "Firtman",
	}
	max.Firstname = "Maximiliano"

	goCourse := data.Course{
		Id:         2,
		Name:       "Go fundamentals",
		Instructor: max,
	}
	swiftWS := data.NewWorkshop("Swift with iOS", max)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}
}
