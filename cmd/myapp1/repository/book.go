package repository

import (
	"context"
	"database/sql"
	"simple-project-native/cmd/myapp1/model"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(ctx context.Context, book *model.Book) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO books (title, author) VALUES ($1, $2)",
		book.Title, book.Author,
	)
	return err
}

func (r *BookRepository) GetByID(ctx context.Context, id int) (*model.Book, error) {
	var book model.Book
	err := r.db.QueryRowContext(ctx,
		"SELECT id, title, author FROM books WHERE id = $1", id,
	).Scan(&book.ID, &book.Title, &book.Author)

	return &book, err
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, author FROM books")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
