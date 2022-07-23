package moudle

// Student 学生类 do for test
type Student struct {
	ID   int
	Name string
}

func (s *Student) GetName() string {
	return s.Name
}

func NewStudent(id int, name string) *Student {
	return &Student{
		ID:   id,
		Name: name,
	}
}
