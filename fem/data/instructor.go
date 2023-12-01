package data

type Instructor struct {
	Id        int
	Firstname string
	Lastname  string
	Score     int
}

func NewInstructor(name string, lastname string) Instructor {
	return Instructor{
		Firstname: name,
		Lastname:  lastname,
	}
}
