package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainHandlerOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/cafe?count=4&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.NotEmpty(t, responseRecorder.Body.String())
	require.Equal(t, http.StatusOK, responseRecorder.Code)

}
func TestMainHandlerWhenWrongCity(t *testing.T) {
	wrongCity := "vladikavkaz"
	target := fmt.Sprintf("localhost:8080/cafe?count=4&city=%s", wrongCity)
	req := httptest.NewRequest(http.MethodGet, target, nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.NotEqual(t, "wrong count value", responseRecorder.Body.String())
	require.NotEqual(t, "count missing", responseRecorder.Body.String())
	require.Equal(t, "wrong city value", responseRecorder.Body.String())
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 5
	city := "moscow"
	target := fmt.Sprintf("localhost:8080/cafe?count=%d&city=%s", totalCount, city)
	req := httptest.NewRequest(http.MethodGet, target, nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.NotEqual(t, "wrong city value", responseRecorder.Body.String())
	require.NotEqual(t, "wrong count value", responseRecorder.Body.String())
	require.NotEqual(t, "count missing", responseRecorder.Body.String())
	gotCount := strings.Count(responseRecorder.Body.String(), ",") + 1
	require.Equal(t, len(cafeList[city]), gotCount)
	require.Equal(t, http.StatusOK, responseRecorder.Code)

}
