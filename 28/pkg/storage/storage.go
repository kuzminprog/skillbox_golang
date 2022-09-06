package storage

import (
	"28/pkg/student"
	"errors"
	"fmt"
)

type Storage map[string]*student.Student

// NewStorage - creates and returns Storage
func NewStorage() Storage {
	return make(Storage)
}

// Put - putting the student in the map with the index name
func (s Storage) Put(value *student.Student) error {
	s[value.GetName()] = value

	if s[value.GetName()] == nil {
		return errors.New("THERE WAS AN ERROR WHEN ADDING DATA")
	} else {
		return nil
	}
}

// Get - returns the student from the map with the index name
func (s Storage) Get(name string) (*student.Student, error) {
	if s[name] == nil {
		return nil, errors.New("THE ELEMENT WITH THIS NAME DOES NOT EXIST")
	} else {
		return s[name], nil
	}
}

// PrintStudents - prints students from the map to the screen
func (s *Storage) PrintStudents() {
	for _, value := range *s {
		fmt.Print(value.GetName(), " ")
		fmt.Print(value.GetAge(), " ")
		fmt.Print(value.GetGrade(), "\n")
	}
}
