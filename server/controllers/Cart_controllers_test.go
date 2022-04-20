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

func Test_CartAddition(t *testing.T) {
	utils.ConnectTestDatabase()

	samplePerson := models.Cart_items_db{
		UserID: "f61dd0d0-5ba4-4353-9485-6c449ac19640",
		// SessionID:  "e019ff27-4f2b-43f6-8d8d-ec6e16b7f31c",
		ProductID:  "a",
		Quantity:   "1",
		CreatedAt:  "4/20/22",
		ModifiedAt: "",
	}
	bytePerson, _ := json.Marshal(samplePerson)

	req, err := http.NewRequest("POST", "/cart/additem", bytes.NewReader(bytePerson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CartAddition)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

}

func Test_CartDisplay(t *testing.T) {
	utils.ConnectTestDatabase()

	samplePerson := models.Cart_items_db{
		UserID: "f61dd0d0-5ba4-4353-9485-6c449ac19640",
	}
	bytePerson, _ := json.Marshal(samplePerson)

	req, err := http.NewRequest("POST", "/cart", bytes.NewReader(bytePerson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CartDisplay)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

}

// func Test_Login(t *testing.T) {
// 	utils.ConnectTestDatabase()

// 	req, err := http.NewRequest("GET", "/user?firstname=Sai&lastname=Reddy&email=nitin1@gmail.com&passkey=sai", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(UserLogin)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
// 	}

// }

// func Test_Signup(t *testing.T) {
// 	utils.ConnectTestDatabase()

// 	samplePerson := models.User3{
// 		Firstname: "Sai",
// 		Lastname:  "Reddy",
// 		Email:     "nitin1@gmail.com",
// 		Password:  "sai",
// 	}
// 	bytePerson, _ := json.Marshal(samplePerson)

// 	req, err := http.NewRequest("POST", "/user", bytes.NewReader(bytePerson))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(UserSignUp)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
// 	}

// }
