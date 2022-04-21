package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"src/models"
	"src/utils"
	"testing"
)

func Test_Return_Store(t *testing.T) {
	utils.ConnectTestDatabase()

	req, err := http.NewRequest("GET", "/stores/items?store_id=ChIJpZbmeDuj6IgRuYWJ6GnlnWw&user_id=f61dd0d0-5ba4-4353-9485-6c449ac19640", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ReturnStoreInv)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

}

func Test_Return_Nearby(t *testing.T) {
	utils.ConnectTestDatabase()

	req, err := http.NewRequest("GET", "/stores/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ReturnNearBy)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

}

func Test_Return_Offers(t *testing.T) {
	utils.ConnectTestDatabase()

	req, err := http.NewRequest("GET", "/offers", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ReturnOffers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

}

func Test_Return_ProdPage(t *testing.T) {
	utils.ConnectTestDatabase()

	req, err := http.NewRequest("GET", "/stores/items/a", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ReturnProductPage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

}

func Test_Return_ProdReview(t *testing.T) {
	utils.ConnectTestDatabase()

	sampleReview := models.Orders{
		Review: "Sai",
	}
	bytePerson, _ := json.Marshal(sampleReview)

	req, err := http.NewRequest("POST", "/stores/items/a", bytes.NewReader(bytePerson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SendProductReview)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

}

func Test_Return_ProdRating(t *testing.T) {
	utils.ConnectTestDatabase()

	sampleRating := models.Orders{
		Rating: 3,
	}
	bytePerson, _ := json.Marshal(sampleRating)

	req, err := http.NewRequest("POST", "/stores/items/a", bytes.NewReader(bytePerson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SendProductRating)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

}
