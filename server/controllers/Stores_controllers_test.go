package controllers

import (
	"net/http"
	"net/http/httptest"
	"src/utils"
	"testing"
)

func Test_Return_Store(t *testing.T) {
	utils.ConnectTestDatabase()

	req, err := http.NewRequest("GET", "/stores/items?store_id=1&user_id=a", nil)
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
