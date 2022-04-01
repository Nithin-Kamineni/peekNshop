package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/models"
	"src/utils"

	"github.com/google/uuid"
)

func SendUserOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 models.User3
	var s2 models.User3
	var orders []models.Orders

	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT ID,userOrders FROM user3 WHERE id = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if s2.Acesskey == s1.Acesskey {
		err = utils.DB.Raw("SELECT * FROM orders WHERE userid = ?", s1.ID).Scan(&orders).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = json.NewEncoder(w).Encode(orders)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func AddingFavorateStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 models.FavorateStoresObj
	var s2 models.User3

	err := json.NewDecoder(r.Body).Decode(&s1) //ID, accesskey, storeID
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT * FROM user3 WHERE ID = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if s2.Acesskey == s1.Acesskey {
		err = utils.DB.Exec("UPDATE user3 SET FavorateStores = array_append(FavorateStores, ? ) WHERE where ID = ?", s1.FavorateStore, s1.ID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		reply := models.SignInReply{Msg: "sucessfully added the store to your favorate stores"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func DeleFavorateStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 models.FavorateStoresObj
	var s2 models.User3

	err := json.NewDecoder(r.Body).Decode(&s1) //ID, accesskey, storeID
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT * FROM user3 WHERE ID = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if s2.Acesskey == s1.Acesskey {
		err = utils.DB.Exec("UPDATE user3 SET FavorateStores = array_remove(FavorateStores, 0) WHERE where ID = ?", s1.FavorateStore, s1.ID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		reply := models.SignInReply{Msg: "sucessfully added the store to your favorate stores"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func ShowFavorateStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 models.FavorateStoresObj
	var s2 models.User3
	var FavStores []string
	var storeInf models.Stores_Information
	var storeInfs []models.Stores_Information

	err := json.NewDecoder(r.Body).Decode(&s1) //ID, accesskey, storeID
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT * FROM user3 WHERE ID = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if s2.Acesskey == s1.Acesskey {
		err = utils.DB.Raw("SELECT FavorateStores FROM user3 WHERE ID = ?", s1.ID).Scan(&FavStores).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		//storeID, photo ref, name, address
		for i := 0; i < len(FavStores); i++ {
			err = utils.DB.Raw("SELECT Stores_Information FROM user3 WHERE StoreID = ?", &FavStores[i]).Scan(&storeInf).Error
			storeInfs = append(storeInfs, storeInf)
		}
		err = json.NewEncoder(w).Encode(storeInfs)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func UserStatusCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 models.User3
	var s2 models.User3
	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT ID,acessKey FROM user3 WHERE id = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(s1.Acesskey)
	fmt.Println(s2.Acesskey)
	fmt.Println()
	if s2.Acesskey == s1.Acesskey {
		err = utils.DB.Exec("UPDATE user3 SET firstname = ?, lastname = ?, email = ?, password = ? where ID = ?", s1.Firstname, s1.Lastname, s1.Email, s1.Password, s2.ID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := models.SignInReply{Msg: "sucessfully changed your details"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}

		//a.db.Raw("SELECT  FROM user3 WHERE acesskey = ?", username)
		// err = a.db.Save(&s).Error
		// if err != nil {
		// 	sendErr(w, http.StatusInternalServerError, err.Error())
		// }
	}
}

func UserStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 models.User3
	var s2 models.User3
	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT ID,acessKey FROM user3 WHERE id = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(s1.Acesskey)
	fmt.Println(s2.Acesskey)
	fmt.Println()
	if s2.Acesskey == s1.Acesskey {
		err = utils.DB.Exec("UPDATE user3 SET firstname = ?, lastname = ?, email = ?, password = ? where ID = ?", s1.Firstname, s1.Lastname, s1.Email, s1.Password, s2.ID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := models.SignInReply{Msg: "sucessfully changed your details"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}

		//a.db.Raw("SELECT  FROM user3 WHERE acesskey = ?", username)
		// err = a.db.Save(&s).Error
		// if err != nil {
		// 	sendErr(w, http.StatusInternalServerError, err.Error())
		// }
	}
}

func ChangeUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 models.User3
	var s2 models.User3
	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT ID,acessKey FROM user3 WHERE id = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(s1.Acesskey)
	fmt.Println(s2.Acesskey)
	fmt.Println()
	if s2.Acesskey == s1.Acesskey {
		err = utils.DB.Exec("UPDATE user3 SET firstname = ?, lastname = ?, email = ? where ID = ?", s1.Firstname, s1.Lastname, s1.Email, s2.ID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := models.SignInReply{Msg: "sucessfully changed your details"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}

		//a.db.Raw("SELECT  FROM user3 WHERE acesskey = ?", username)
		// err = a.db.Save(&s).Error
		// if err != nil {
		// 	sendErr(w, http.StatusInternalServerError, err.Error())
		// }
	}
}

func ChangeUserAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 models.User3
	var s2 models.ChangeUserAddress
	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT ID,acessKey FROM user3 WHERE id = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(s1.Acesskey)
	fmt.Println(s2.Acesskey)
	fmt.Println()
	if s2.Acesskey == s1.Acesskey {
		err = utils.DB.Exec("UPDATE user3 SET Address1= ? where ID = ?", s2.Address, s2.ID).Error //add s1.address insted of s1.firstname
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := models.SignInReply{Msg: "sucessfully changed your delivary address"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}

		//a.db.Raw("SELECT  FROM user3 WHERE acesskey = ?", username)
		// err = a.db.Save(&s).Error
		// if err != nil {
		// 	sendErr(w, http.StatusInternalServerError, err.Error())
		// }
	}
}

func ForgotUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var email models.RetrevalDetails
	var s models.User3
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT id,email,acesskey FROM user3 WHERE email = ?", email.Email).Scan(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}

}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(statusCode: 200)
	//w.WriteHeader(statusCode: 200)
	w.Header().Set("Content-Type", "application/json")

	//params := mux.Vars(r)
	//username := params["username"]
	//fmt.Println(username)
	var s models.User3
	var reply models.LogInReply

	username := r.URL.Query().Get("email")
	passkey := r.URL.Query().Get("passkey")
	//credentials := a.db.First(&s, "email = ?", username)
	err := utils.DB.Raw("SELECT id FROM user3 WHERE email = ?", username).Scan(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	// a.db.where("username = ?",username)
	//fmt.Println(&s)
	//data, err := json.Marshal(&s)

	if s.ID == "" {
		fmt.Println("User does not exist/registered")
		reply = models.LogInReply{Msg: "User does not exist/registered", UserDetails: s, AllowUsers: false}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	} else {
		fmt.Println(s.ID)
		err = utils.DB.Raw("SELECT * FROM user3 WHERE id = ? AND password = ?", s.ID, passkey).Scan(&s).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		if s.Email == "" {
			fmt.Println("Password is incorrect")
			reply = models.LogInReply{Msg: "Password is incorrect", UserDetails: s, AllowUsers: false}
			err = json.NewEncoder(w).Encode(reply)
			if err != nil {
				sendErr(w, http.StatusInternalServerError, err.Error())
			}
		} else {
			fmt.Println("Login Sucessfull")
			reply = models.LogInReply{Msg: "Login Sucessfull", UserDetails: s, AllowUsers: true}
			err = json.NewEncoder(w).Encode(reply)
			if err != nil {
				sendErr(w, http.StatusInternalServerError, err.Error())
			}
		}
		//fmt.Println(string(data))
		fmt.Println()
	}
}

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s models.User3
	var s1 models.User3
	//var id string

	reply_on_fail := models.SignInReply{Msg: "Email already exists, try using another email."}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusCreated)
	}
	fmt.Println(s.Email)
	err = utils.DB.Raw("SELECT * FROM user3 WHERE email = ?", s.Email).Scan(&s1).Error
	fmt.Println(s1.ID)
	if s1.ID == "" {
		s.ID = uuid.New().String()
		fmt.Println("if")
		fmt.Println(s.Firstname)
		err = utils.DB.Table("user3").Save(&s).Error
		fmt.Println("if2")
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusCreated)
			reply := models.LogInReply{Msg: "Login and sign-up Sucessfull", UserDetails: s, AllowUsers: true}
			err = json.NewEncoder(w).Encode(reply)
			if err != nil {
				sendErr(w, http.StatusInternalServerError, err.Error())
			}
		}
	} else {
		fmt.Println("else")
		err = json.NewEncoder(w).Encode(reply_on_fail)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}
