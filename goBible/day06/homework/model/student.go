package model

import "errors"

var (
	ErrBookNotFount = errors.New("this book not fount...")
)

type Student struct {
	Name  string
	Grade string
	Id    string
	Sex   string
	books []*BorrowItem
}

type BorrowItem struct {
	book *Book
	num  int
}

func NewStudent(name, grade, id, sex string) *Student {
	return &Student{
		Name:  name,
		Grade: grade,
		Id:    id,
		Sex:   sex,
	}
}

func (s *Student) AddBook(b *BorrowItem) {
	s.books = append(s.books, b)
}

func (s *Student) DelBook(b *BorrowItem) error {
	for i := 0; i < len(s.books); i++ {
		if s.books[i].book.Name == b.book.Name {
			s.books[i].num -= b.num
			if s.books[i].num == 0 {
				s.books = append(s.books[0:i], s.books[i+1:]...)
			}
			return nil
		}
	}
	return ErrBookNotFount
}

func (s *Student) BookList() []*BorrowItem {
	return s.books
}
