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

func Test_HomePageReload(t *testing.T) {
	utils.ConnectTestDatabase()

	samplePerson := models.Coardinates{
		Lat: "29.6213925",
		Lon: "-82.37358560000001",
	}
	bytePerson, _ := json.Marshal(samplePerson)

	req, err := http.NewRequest("POST", "/address/city", bytes.NewReader(bytePerson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomePageReload)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

}

func Test_ConvAddressToCord(t *testing.T) {
	utils.ConnectTestDatabase()

	samplePerson := models.AddressForm{
		Street:  "3800 SW 34th",
		City:    "Gainesville",
		State:   "FL",
		Zipcode: "32608",
	}
	bytePerson, _ := json.Marshal(samplePerson)

	req, err := http.NewRequest("POST", "/address", bytes.NewReader(bytePerson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ConvAddressToCord)
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
