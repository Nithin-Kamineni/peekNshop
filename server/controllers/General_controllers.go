package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/src/Stores"
	"src/src/Users"
	"src/utils"

	"github.com/google/uuid"
)

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store Stores.Store_inventory
	reply := Users.SignInReply{Msg: "sucessfull"}
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	store.StoreID = uuid.New().String()
	err = utils.DB.Save(&store).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func homePageReload(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(statusCode: 200)
	//w.WriteHeader(statusCode: 200)
	w.Header().Set("Content-Type", "application/json")

	//params := mux.Vars(r)
	//username := params["username"]
	//fmt.Println(username)
	var reply Users.HomePageCity
	var cord Users.Coardinates

	err := json.NewDecoder(r.Body).Decode(&cord)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	//credentials := a.db.First(&s, "email = ?", username)
	fmt.Println("Gainsville")
	reply = Users.HomePageCity{City: "Gainsvile"}
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
