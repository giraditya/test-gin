package repository

import (
	"context"

	"github.com/giriaditya/test-gin/presentation/request"
	"github.com/giriaditya/test-gin/presentation/response"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookRepository interface {
	FetchAll(ctx context.Context, db *gorm.DB) ([]response.BookFetchResponse, error)
	Create(ctx context.Context, db *gorm.DB, request request.BookCreateRequest) (response.BookCreateResponse, error)
	Update(ctx context.Context, db *gorm.DB, request request.BookUpdateRequest, id int) (response.BookUpdateResponse, error)
	Delete(ctx context.Context, db *gorm.DB, id int) error
}

type BookRepositoryImplementation struct{}

func NewBookRepository() BookRepository {
	return &BookRepositoryImplementation{}
}

func (repository *BookRepositoryImplementation) FetchAll(ctx context.Context, db *gorm.DB) ([]response.BookFetchResponse, error) {
	var result []response.BookFetchResponse
	books := []Book{}
	db.Find(&books)
	for _, v := range books {
		row := response.BookFetchResponse{
			ID:     v.ID,
			Title:  v.Title,
			Author: v.Author,
		}
		result = append(result, row)
	}
	return result, nil
}

func (repository *BookRepositoryImplementation) Create(ctx context.Context, db *gorm.DB, request request.BookCreateRequest) (response.BookCreateResponse, error) {
	var result response.BookCreateResponse
	book := Book{
		Title:  request.Title,
		Author: request.Author,
	}
	db.Create(&book)
	if err := db.Error; err != nil {
		return result, err
	}

	result.ID = book.ID
	result.Title = book.Title
	result.Author = book.Author
	return result, nil
}

func (repository *BookRepositoryImplementation) Update(ctx context.Context, db *gorm.DB, request request.BookUpdateRequest, id int) (response.BookUpdateResponse, error) {
	var result response.BookUpdateResponse
	book := Book{
		Title:  request.Title,
		Author: request.Author,
	}
	db.Model(&book).Where("id = ?", id).Updates(book)
	if err := db.Error; err != nil {
		return result, err
	}

	result.Title = request.Title
	result.Author = request.Author
	return result, nil
}

func (repository *BookRepositoryImplementation) Delete(ctx context.Context, db *gorm.DB, id int) error {
	book := Book{}
	db.Model(&book).Delete(book, id)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}
