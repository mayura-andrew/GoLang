package data

type Instructor struct {
	Id int
	FirstName string
	LastName string
	Score int
}
func NewInstructor(name string, lastname string, score int) Instructor {
	return Instructor{FirstName: name, LastName: lastname, Score: score}
}