package controllers

import (
	"encoding/json"
	"net/http"
	"src/models"
	"src/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func OrderPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cartItems []models.Cart_items_db
	var orderProc models.Orders
	var productQuantity int64
	orderID := uuid.New().String()
	err := json.NewDecoder(r.Body).Decode(&cartItems)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	for i := 0; i < len(cartItems); i++ {

		CurrentTime := time.Now()
		orderProc.OrderID = orderID
		orderProc.ProductID = cartItems[i].ProductID
		orderProc.UserID = cartItems[i].UserID
		orderProc.Quantity = cartItems[i].Quantity

		orderProc.Product_name = cartItems[i].Product_name
		orderProc.Product_photo = cartItems[i].Product_photo

		orderProc.Product_name = cartItems[i].Description

		orderProc.StoreID = cartItems[i].StoreID //

		orderProc.OrderedOn = CurrentTime.String()
		orderProc.DeliveredOn = "" //
		orderProc.PickedUpOn = ""  //

		orderProc.Review = ""
		orderProc.Rating = 0.0
	}

	err = utils.DB.Raw("SELECT quantity FROM store_inventories WHERE product_id = ?", orderProc.ProductID).Scan(&productQuantity).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	quantityInt, _ := strconv.ParseInt(orderProc.Quantity, 10, 0)
	if productQuantity-quantityInt > 0 {
		err = utils.DB.Exec("Update store_inventories set quantity = ?, modified_at = ? where productID = ?", productQuantity-quantityInt, time.Now().String(), orderProc.ProductID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = utils.DB.Table("user3").Save(&orderProc).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	} else if (productQuantity - quantityInt) == 0 {
		err = utils.DB.Exec("Delete store_inventories where productID = ?", orderProc.ProductID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		err = utils.DB.Table("user3").Save(&orderProc).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	} else {
		reply := models.SignInReply{Msg: "Not enought quantity of products"} //product id, quantity , name
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}

}

func DisplayOrders(w http.ResponseWriter, r *http.Request) { //secure
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

func OrderReview(w http.ResponseWriter, r *http.Request) { //secure
	w.Header().Set("Content-Type", "application/json")
	var OrderReview models.Orders

	err := json.NewDecoder(r.Body).Decode(&OrderReview)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Exec("Update orders set review = ?, rating = ? where orderID = ? and productID = ?", OrderReview.Review, OrderReview.Rating, OrderReview.OrderID, OrderReview.ProductID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func OrderReviewEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var OrderReview models.Orders

	err := json.NewDecoder(r.Body).Decode(&OrderReview)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Exec("Update orders set review = ?, rating = ? where orderID = ? and productID = ?", OrderReview.Review, OrderReview.Rating, OrderReview.OrderID, OrderReview.ProductID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func OrderDelivary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var OrderReview models.Orders

	err := json.NewDecoder(r.Body).Decode(&OrderReview)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Exec("Update orders set deliveredon = ?, pickedOn = ? where orderID = ? and productID = ?", OrderReview.DeliveredOn, OrderReview.PickedUpOn, OrderReview.OrderID, OrderReview.ProductID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
}
