package services

import (
	"errors"
	"time"

	"github.com/MSSkowron/BookRESTAPI/internal/database"
	"github.com/MSSkowron/BookRESTAPI/internal/dtos"
	"github.com/MSSkowron/BookRESTAPI/internal/models"
)

var (
	// ErrInvalidID is returned when the given id is not a positive integer.
	ErrInvalidID = errors.New("id must be a positive integer")
	// ErrInvalidCreatedByID is returned when the given createdByID is not a positive integer.
	ErrInvalidCreatedByID = errors.New("createdByID must be a positive integer")
	// ErrInvalidAuthor is returned when the given author is empty.
	ErrInvalidAuthor = errors.New("author must not be empty")
	// ErrInvalidTitle is returned when the given title is empty.
	ErrInvalidTitle = errors.New("title must not be empty")
	// ErrInvalidAuthorOrTitle is returned when the given author or title is empty.
	ErrInvalidAuthorOrTitle = errors.New("invalid author or title")
	// ErrBookNotFound is returned when the book with the given id does not exist in the database.
	ErrBookNotFound = errors.New("book not found")
)

// BookService is an interface that defines the methods that the BookService struct must implement.
type BookService interface {
	GetBooks() ([]*dtos.BookDTO, error)
	GetBook(int) (*dtos.BookDTO, error)
	AddBook(int, *dtos.BookCreateDTO) (*dtos.BookDTO, error)
	UpdateBook(int, *dtos.BookDTO) (*dtos.BookDTO, error)
	DeleteBook(int) error
}

// BookServiceImpl is a struct that implements the BookService interface.
type BookServiceImpl struct {
	db database.Database
}

// NewBookService creates a new BookServiceImpl.
func NewBookService(db database.Database) *BookServiceImpl {
	return &BookServiceImpl{db: db}
}

// GetBooks returns all books from the database.
func (bs *BookServiceImpl) GetBooks() ([]*dtos.BookDTO, error) {
	books, err := bs.db.SelectAllBooks()
	if err != nil {
		return nil, err
	}

	booksDTO := []*dtos.BookDTO{}
	for _, book := range books {
		booksDTO = append(booksDTO, &dtos.BookDTO{
			ID:        int64(book.ID),
			Author:    book.Author,
			Title:     book.Title,
			CreatedAt: book.CreatedAt,
		})
	}

	return booksDTO, nil
}

// GetBook returns a book with the given id from the database.
func (bs *BookServiceImpl) GetBook(id int) (*dtos.BookDTO, error) {
	if !bs.validateID(id) {
		return nil, ErrInvalidID
	}

	book, err := bs.db.SelectBookByID(id)
	if err != nil || book == nil {
		return nil, ErrBookNotFound
	}

	return &dtos.BookDTO{
		ID:        int64(book.ID),
		Author:    book.Author,
		Title:     book.Title,
		CreatedAt: book.CreatedAt,
	}, nil
}

// AddBook adds a book.
func (bs *BookServiceImpl) AddBook(createdByID int, dto *dtos.BookCreateDTO) (*dtos.BookDTO, error) {
	if !bs.validateID(createdByID) {
		return nil, ErrInvalidCreatedByID
	}
	if !bs.validateAuthor(dto.Author) {
		return nil, ErrInvalidAuthor
	}
	if !bs.validateTitle(dto.Title) {
		return nil, ErrInvalidTitle
	}

	id, err := bs.db.InsertBook(&models.Book{
		CreatedBy: createdByID,
		CreatedAt: time.Now(),
		Author:    dto.Author,
		Title:     dto.Title,
	})
	if err != nil {
		return nil, err
	}

	book, err := bs.db.SelectBookByID(id)
	if err != nil {
		return nil, err
	}

	return &dtos.BookDTO{
		ID:        int64(book.ID),
		CreatedAt: book.CreatedAt,
		Author:    book.Author,
		Title:     book.Title,
	}, nil
}

// UpdateBook updates a book with the given id.
func (bs *BookServiceImpl) UpdateBook(id int, dto *dtos.BookDTO) (*dtos.BookDTO, error) {
	if !bs.validateID(id) {
		return nil, ErrInvalidID
	}
	if !bs.validateAuthor(dto.Author) {
		return nil, ErrInvalidAuthor
	}
	if !bs.validateTitle(dto.Title) {
		return nil, ErrInvalidTitle
	}

	book, err := bs.db.SelectBookByID(id)
	if err != nil || book == nil {
		return nil, ErrBookNotFound
	}

	book.Author = dto.Author
	book.Title = dto.Title
	if err := bs.db.UpdateBook(id, book); err != nil {
		return nil, err
	}

	return &dtos.BookDTO{
		ID:        int64(book.ID),
		Author:    book.Author,
		Title:     book.Title,
		CreatedAt: book.CreatedAt,
	}, nil
}

// DeleteBook deletes a book with the given id.
func (bs *BookServiceImpl) DeleteBook(id int) error {
	if !bs.validateID(id) {
		return ErrInvalidID
	}

	if book, err := bs.db.SelectBookByID(id); err != nil || book == nil {
		return ErrBookNotFound
	}

	return bs.db.DeleteBook(id)
}

// validateID validates the given id.
func (bs *BookServiceImpl) validateID(id int) bool {
	return id > 0
}

// validateAuthor validates the given author.
func (bs *BookServiceImpl) validateAuthor(author string) bool {
	return author != ""
}

// validateTitle validates the given title.
func (bs *BookServiceImpl) validateTitle(title string) bool {
	return title != ""
}
