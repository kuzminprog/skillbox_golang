package student

type Student struct {
	name  string
	age   int
	grade int
}

// NewStudent - creates and returns *Storage
func NewStudent() *Student {
	return new(Student)
}

// GetName - returns the name of the student
func (s *Student) GetName() string {
	return s.name
}

// GetAge - returns the age of the student
func (s *Student) GetAge() int {
	return s.age
}

// GetGrade - returns the grade of the student
func (s *Student) GetGrade() int {
	return s.grade
}

// SetName - places the name in the structure Student
func (s *Student) SetName(name string) {
	s.name = name
}

// SetAge - places the age in the structure Student
func (s *Student) SetAge(age int) {
	s.age = age
}

// SetGrade - places the grade in the structure Student
func (s *Student) SetGrade(grade int) {
	s.grade = grade
}
