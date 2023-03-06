package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Alfeenn/article/controller"
	"github.com/Alfeenn/article/helper"
	"github.com/Alfeenn/article/middleware"
	"github.com/Alfeenn/article/repository"
	"github.com/Alfeenn/article/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golangdb")
	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func TestCreateDataSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB()
	repository := repository.NewRepository()
	service := service.NewService(repository, db)
	controller := controller.NewController(service)
	tx, _ := db.Begin()
	router := gin.Default()
	requestBody := strings.NewReader(`{"name" : "Gadget","status" : "Gadget","visibility" : "Gadget","details" : "Gadget"}`)
	tx.Commit()
	router.Use(middleware.NewMiddleware())
	router.POST("/api/categories", controller.Create)

	request := httptest.NewRequest(http.MethodPost, "/api/categories", requestBody)
	request.Header.Add("x-API-key", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "CREATED", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["visibility"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["details"])

}

func TestPing(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewRepository()
	service := service.NewService(repo, db)
	controller := controller.NewController(service)
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	w := httptest.NewRecorder()
	router.GET("/ping", controller.Ping)
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, req)
	//Assertion
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, "pong", w.Body.String())
}
