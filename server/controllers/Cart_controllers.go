package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/models"
	"src/utils"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func CartAddition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart models.Cart_items_db
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	cart.Quantity = "1"
	err = utils.DB.Raw("SELECT product_name, product_photo, description, product_price FROM store_inventories WHERE product_id = ?", cart.ProductID).Scan(&cart).Error
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	// cart.Product_name =
	// cart.Product_photo =
	// cart.Description =
	cart.SessionID = uuid.New().String()

	err = utils.DB.Save(&cart).Error
	if err != nil {
		fmt.Print("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func CartManipulation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart models.Cart_items_db
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
		err = utils.DB.Exec("UPDATE cart_items_dbs SET quantity = ?, Modified_At = ? where user_ID = ? and product_ID = ?", cart.Quantity, cart.ModifiedAt, cart.UserID, cart.ProductID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		reply := models.SignInReply{Msg: "Not enough inv to add items in cart"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func CartDeletion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart models.Cart_items_db

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	err = utils.DB.Exec("DELETE cart_items_dbs where user_ID = ? and product_ID = ?", cart.UserID, cart.ProductID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func ClearCart(w http.ResponseWriter, r *http.Request) {
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

func CartDisplay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart []models.Cart_items_db
	var userID models.UserIDtab
	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT * FROM cart_items_dbs WHERE user_id = ?", userID.UserID).Scan(&cart).Error
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	for i := 0; i < len(cart); i++ {
		cart[0].Product_price = strings.Replace(cart[0].Product_price, "$", "", 1)
	}

	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
