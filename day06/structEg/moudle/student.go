package moudle

type Student struct {
	id      int
	name    string
	age     int
	address string
}

func NewStudent(id int, name string, age int, address string) *Student {
	stu := &Student{
		id:      id,
		name:    name,
		age:     age,
		address: address,
	}
	return stu
}

func (s *Student) GetId() int {
	return s.id
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) GetAge() int {
	return s.age
}

func (s *Student) GetAddress() string {
	return s.address
}
