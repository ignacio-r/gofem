package data

import "fmt"

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d)",
		i.Lastname, i.Firstname, i.Score)
}
