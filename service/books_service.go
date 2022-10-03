package service

import (
	"context"

	"github.com/giriaditya/test-gin/models"
	"github.com/giriaditya/test-gin/presentation/request"
	"github.com/giriaditya/test-gin/presentation/response"
	"gorm.io/gorm"
)

type BookService interface {
	FetchAll(ctx context.Context) ([]response.BookFetchResponse, error)
	Create(ctx context.Context, request request.BookCreateRequest) (response.BookCreateResponse, error)
	Update(ctx context.Context, request request.BookUpdateRequest, id int) (response.BookUpdateResponse, error)
	Delete(ctx context.Context, id int) error
}

type BookServiceImplementation struct {
	DB *gorm.DB
}

func NewBookService(db *gorm.DB) BookService {
	return &BookServiceImplementation{
		DB: db,
	}
}

func (service *BookServiceImplementation) Create(ctx context.Context, request request.BookCreateRequest) (response.BookCreateResponse, error) {
	var result response.BookCreateResponse

	book := models.Book{
		Title:  request.Title,
		Author: request.Author,
	}

	tx := service.DB.Begin()
	tx.Create(&book)

	if tx.Error != nil {
		tx.Rollback()
		return result, tx.Error
	} else {
		tx.Commit()
	}

	result.ID = book.ID
	result.Title = book.Title
	result.Author = book.Author

	return result, nil
}

func (service *BookServiceImplementation) Update(ctx context.Context, request request.BookUpdateRequest, id int) (response.BookUpdateResponse, error) {
	var result response.BookUpdateResponse

	book := models.Book{
		Title:  request.Title,
		Author: request.Author,
	}

	tx := service.DB.Begin()
	tx.Model(&book).Where("id = ?", id).Updates(book)

	if tx.Error != nil {
		tx.Rollback()
		return result, tx.Error
	} else {
		tx.Commit()
	}

	result.Author = request.Author
	result.Title = request.Title

	return result, nil
}

func (service *BookServiceImplementation) Delete(ctx context.Context, id int) error {
	book := models.Book{}

	tx := service.DB.Begin()
	tx.Model(&book).Delete(book, id)

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	} else {
		tx.Commit()
	}

	return nil
}

func (service *BookServiceImplementation) FetchAll(ctx context.Context) ([]response.BookFetchResponse, error) {
	var results []response.BookFetchResponse
	books := []models.Book{}
	service.DB.Find(&books)

	for _, v := range books {
		result := response.BookFetchResponse{
			Title:  v.Title,
			Author: v.Author,
		}
		results = append(results, result)
	}

	return results, nil
}
