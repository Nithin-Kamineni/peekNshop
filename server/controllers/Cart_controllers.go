package controllers

import (
	"encoding/json"
	"net/http"
	"src/src/Carts"
	"src/utils"
)

func cartAddition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart Carts.Cart_items
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

func cartDisplay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart Carts.Cart_items
	var userID Carts.UserIDtab
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
