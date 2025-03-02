package service

import (
	"context"
	"simple-project-native/cmd/myapp1/model"
	"simple-project-native/cmd/myapp1/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(ctx context.Context, book *model.Book) error {
	return s.repo.Create(ctx, book)
}

func (s *BookService) GetBook(ctx context.Context, id int) (*model.Book, error) {
	return s.repo.GetByID(ctx, id)
}
