package book

import (
	"context"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"
)

func init() {
	service.RegisterBook(&sBook{})
}

type sBook struct{}

// Add implements service.IBook.
func (s *sBook) Add(ctx context.Context, book do.Book) (err error) {
	panic("unimplemented")
}

// Del implements service.IBook.
func (s *sBook) Del(ctx context.Context) (err error) {
	panic("unimplemented")
}

// Edit implements service.IBook.
func (s *sBook) Edit(ctx context.Context, book do.Book) (err error) {
	panic("unimplemented")
}

// GetList implements service.IBook.
func (s *sBook) GetList(ctx context.Context) (books []entity.Book, err error) {
	// panic("unimplemented")
	err = dao.Book.Ctx(ctx).Scan(&books)
	return
}
