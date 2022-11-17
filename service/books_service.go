package service

import (
	"context"

	"github.com/giriaditya/test-gin/presentation/request"
	"github.com/giriaditya/test-gin/presentation/response"
	"github.com/giriaditya/test-gin/repository"
	"gorm.io/gorm"
)

type BookService interface {
	FetchAll(ctx context.Context) ([]response.BookFetchResponse, error)
	Create(ctx context.Context, request request.BookCreateRequest) (response.BookCreateResponse, error)
	Update(ctx context.Context, request request.BookUpdateRequest, id int) (response.BookUpdateResponse, error)
	Delete(ctx context.Context, id int) error
}

type BookServiceImplementation struct {
	DB         *gorm.DB
	Repository repository.BookRepository
}

func NewBookService(db *gorm.DB, repository repository.BookRepository) BookService {
	return &BookServiceImplementation{
		DB:         db,
		Repository: repository,
	}
}

func (service *BookServiceImplementation) FetchAll(ctx context.Context) ([]response.BookFetchResponse, error) {
	db := service.DB
	result, err := service.Repository.FetchAll(ctx, db)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *BookServiceImplementation) Create(ctx context.Context, request request.BookCreateRequest) (response.BookCreateResponse, error) {
	db := service.DB.Begin()
	result, err := service.Repository.Create(ctx, db, request)
	if err != nil {
		db.Rollback()
		return result, err
	}
	db.Commit()
	return result, err
}

func (service *BookServiceImplementation) Update(ctx context.Context, request request.BookUpdateRequest, id int) (response.BookUpdateResponse, error) {
	db := service.DB.Begin()
	result, err := service.Repository.Update(ctx, db, request, id)
	if err != nil {
		db.Rollback()
		return result, err
	}
	db.Commit()
	return result, err
}

func (service *BookServiceImplementation) Delete(ctx context.Context, id int) error {
	db := service.DB.Begin()
	err := service.Repository.Delete(ctx, db, id)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return err
}
