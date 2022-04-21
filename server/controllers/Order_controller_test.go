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

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"src/models"
// 	"src/utils"
// 	"testing"
// )

func Test_OrderPayment(t *testing.T) {
	utils.ConnectTestDatabase()

	samplePerson := models.Cart_items_db{
		UserID:        "f61dd0d0-5ba4-4353-9485-6c449ac19640",
		ProductID:     "a",
		Product_name:  "Pepsi",
		Product_photo: "asda",
		Product_price: "$2.99",
		Description:   "A cool drink",
		StoreID:       "ChIJpZbmeDuj6IgRuYWJ6GnlnWw",
		Quantity:      "1",
		CreatedAt:     "4/20/22",
		ModifiedAt:    "",
	}
	bytePerson, _ := json.Marshal(samplePerson)

	req, err := http.NewRequest("POST", "/order/payment", bytes.NewReader(bytePerson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(OrderPayment)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

}
