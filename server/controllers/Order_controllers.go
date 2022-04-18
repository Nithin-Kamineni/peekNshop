package controllers

import (
	"encoding/json"
	"net/http"
	"src/models"
	"src/utils"
	"time"

	"github.com/google/uuid"
)

func OrderPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cartItems []models.Cart_items_db
	var orderProc models.Orders
	orderID := uuid.New().String()
	err := json.NewDecoder(r.Body).Decode(&cartItems)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	for i := 0; i < len(cartItems); i++ {

		currentTime := time.Now()
		orderProc.OrderID = orderID
		orderProc.ProductID = cartItems[i].ProductID
		orderProc.UserID = cartItems[i].UserID
		orderProc.Quantity = cartItems[i].Quantity

		orderProc.Product_name = cartItems[i].Product_name
		orderProc.Product_photo = cartItems[i].Product_photo

		orderProc.Product_name = cartItems[i].Description

		orderProc.StoreID = cartItems[i].StoreID //

		orderProc.OrderedOn = currentTime.String()
		orderProc.DeliveredOn = "" //
		orderProc.PickedUpOn = ""  //

		orderProc.Review = ""
		orderProc.Rating = 0.0
	}
}

func DisplayOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []models.Orders
	var user models.User3

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT orders FROM favorate_stores_objs WHERE user_id = ?", user.ID).Scan(&orders).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func OrderReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart models.Cart_items_db

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Exec("DELETE cart_items_dbs where userID = ?", cart.UserID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func OrderReviewEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart models.Cart_items_db

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Exec("DELETE cart_items_dbs where userID = ?", cart.UserID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
}
