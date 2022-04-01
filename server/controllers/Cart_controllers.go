package controllers

import (
	"encoding/json"
	"net/http"
	"src/models"
	"src/utils"
	"strconv"
)

func CartAddition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart models.Cart_items
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Save(&cart).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func CartManipulation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart models.Cart_items
	var quantity int64

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT quantity FROM store_inventories WHERE product_id = ?", cart.ProductID).Scan(&quantity).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	quantityInt, _ := strconv.ParseInt(cart.Quantity, 10, 0)
	if quantity >= quantityInt {
		err = utils.DB.Exec("UPDATE store_inventories SET quantity = ?, ModifiedAt = ? where userID = ? and productID = ?", cart.Quantity, cart.ModifiedAt, cart.UserID, cart.ProductID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func CartDisplay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart models.Cart_items
	var userID models.UserIDtab
	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT * FROM user3 WHERE userID = ?", userID).Scan(&cart).Error
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
