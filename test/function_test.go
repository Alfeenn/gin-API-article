package test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Alfeenn/article/controller"
	"github.com/Alfeenn/article/model"
	"github.com/Alfeenn/article/repository"
	"github.com/Alfeenn/article/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateDataSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := SetupTestDB()
	tx, _ := db.Begin()
	router := SetupRouter(db)
	tx.Commit()

	requestBody := strings.NewReader(`{"name" : "spider man release date","status" : "published","visibility" : "Gadget","details" : "Gadget"}`)
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
	assert.Equal(t, "spider man release date", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "published", responseBody["data"].(map[string]interface{})["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["visibility"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["details"])

}

func TestCreateDataFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := SetupTestDB()
	tx, _ := db.Begin()
	tx.Commit()

	router := SetupRouter(db)
	reqBody := strings.NewReader(`{"name":""}`)

	request := httptest.NewRequest(http.MethodPost, "/api/categories", reqBody)
	request.Header.Add("x-api-key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

}

func TestUpdateSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := SetupTestDB()
	tx, _ := db.Begin()
	router := SetupRouter(db)
	categoryRepository := repository.NewRepository()
	model := categoryRepository.Create(context.Background(), tx, model.Article{
		Name: "yhahahahayyuk",
	})
	tx.Commit()

	reqBody := strings.NewReader(`{"name":"yhahahahayyuk"}`)

	request := httptest.NewRequest(http.MethodPut, "/api/categories/"+model.Id, reqBody)
	request.Header.Add("x-API-key", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "yhahahahayyuk", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := SetupTestDB()
	tx, _ := db.Begin()
	router := SetupRouter(db)
	categoryRepository := repository.NewRepository()
	model := categoryRepository.Create(context.Background(), tx, model.Article{
		Name: "article update",
	})
	tx.Commit()

	reqBody := strings.NewReader(`{"name":""}`)

	request := httptest.NewRequest(http.MethodPut, "/api/categories/"+model.Id, reqBody)
	request.Header.Add("x-API-key", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

}

func TestDeleteSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := SetupTestDB()
	tx, _ := db.Begin()
	router := SetupRouter(db)
	categoryRepository := repository.NewRepository()
	model := categoryRepository.Create(context.Background(), tx, model.Article{
		Name: "article delete",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodPost, "/api/categories/"+model.Id, nil)
	request.Header.Add("x-API-key", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

}

func TestDeleteFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := SetupTestDB()
	tx, _ := db.Begin()
	router := SetupRouter(db)

	tx.Commit()
	request := httptest.NewRequest(http.MethodPost, "/api/categories/404", nil)
	request.Header.Add("x-API-key", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])

}
func TestPing(t *testing.T) {
	db := SetupTestDB()
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
