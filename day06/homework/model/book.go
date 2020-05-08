package model

import (
	"errors"
	"time"
)

var (
	ErrLessBookTotal = errors.New("this book is less than you borrow number...")
)

type Book struct {
	Name       string
	Total      int
	Author     string
	CreateTime time.Time
}

func NewBook(name string, total int, author string, createTime time.Time) *Book {
	return &Book{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateTime: createTime,
	}
}

func (b *Book) CanBorrow(n int) bool {
	return b.Total >= n
}

func (b *Book) Borrow(n int) error {
	if b.CanBorrow(n) == false {
		return ErrLessBookTotal
	}

	b.Total -= n
	return
}

func (b *Book) Back(n int) error {
	b.Total += n
	return
}
