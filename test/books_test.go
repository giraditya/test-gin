package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/giriaditya/test-gin/app"
	"github.com/giriaditya/test-gin/controllers"
	"github.com/giriaditya/test-gin/presentation/request"
	"github.com/giriaditya/test-gin/repository"
	"github.com/giriaditya/test-gin/service"
	"github.com/stretchr/testify/assert"
)

var (
	db              = app.ConnectDatabase()
	booksRepository = repository.NewBookRepository()
	bookService     = service.NewBookService(db, booksRepository)
	bookController  = controllers.NewBookController(bookService)
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestFetch(t *testing.T) {
	r := SetupRouter()
	r.GET("/books", bookController.FetchAll)

	req, _ := http.NewRequest("GET", "/books", bytes.NewBuffer([]byte{}))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w)
}

func TestCreate(t *testing.T) {
	r := SetupRouter()
	r.POST("/books", bookController.Create)

	userRequest := request.BookCreateRequest{
		Title:  "Test Create Book",
		Author: "Self",
	}
	testBody, _ := json.Marshal(userRequest)
	req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(testBody))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdate(t *testing.T) {
	r := SetupRouter()
	r.PUT("/books/:id", bookController.Update)

	id := "5"
	userRequest := request.BookUpdateRequest{
		Title:  "Test Update Book",
		Author: "Self",
	}
	testBody, _ := json.Marshal(userRequest)
	req, _ := http.NewRequest(http.MethodPut, "/books/"+id, bytes.NewBuffer(testBody))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDelete(t *testing.T) {
	r := SetupRouter()
	r.DELETE("/books/:id", bookController.Delete)

	id := "4"
	req, _ := http.NewRequest(http.MethodDelete, "/books/"+id, bytes.NewBuffer([]byte{}))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
