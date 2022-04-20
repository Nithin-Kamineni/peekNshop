package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"src/models"
// 	"src/utils"
// 	"testing"
// )

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
