package data

import "fmt"

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v, %v (%d)", i.Id, i.LastName, i.FirstName, i.Score)
}