package calc_test

import (
	"src/calc"
	"testing"
	"github.com/gorilla/mux"
	"src"
)

type TestCase struct {
	valueA   int
	valueB   int
	expected int
	actual   int
}

func Router(a *App) *mux.Router {
	router := mux.NewRouter()
	router.HandelFunc("/user", a.userLogin).Methods("GET")
	return router
}

func TestRootEndpoint(t *testing.T) {
	request, _ = http.NewRequest(method:"GET", url: "/user", body: nil)
	response := httptest.NewRecord()
	Router().ServeHTTP(responce, request)
	assert.Equal(t, expected: 200, responce.code, msgAndArgs:"OK responce is expected")
}
