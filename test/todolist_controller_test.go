package test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/Ardnh/go-todolist.git/app"
	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/go-playground/validator/v10"
)

func SetUpTestDB() *sql.DB {
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/go-todolist")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func SetupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	router := app.NewRouter(db, validate)

	return router
}

func TruncateTodolist(db *sql.DB) {
	db.Exec("TRUNCATE todolist")
}

func TestRegisterSuccess(t *testing.T) {
	db := SetUpTestDB()
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{
		"firstname": "Uzumaki",
		"lastname": "Naruto", 
		"username": "naruto",
		"password": "12345"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	// body, _ := io.ReadAll(response.Body)
	// var responseBody web.WebResponse
	// json.Unmarshal(body, &responseBody)

	// assert.Equal(t, 200, responseBody.Code)
	// assert.Equal(t, "OK", responseBody.Status)
}
