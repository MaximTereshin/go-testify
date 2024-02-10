package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	totalCount int    = 4
	expected   string = "wrong city"
)

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=7&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=@MeRiK@", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, expected, responseRecorder.Body.String())
}
