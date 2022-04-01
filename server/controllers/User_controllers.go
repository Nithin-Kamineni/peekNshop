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
	var cord models.Coardinates

	username := r.URL.Query().Get("email")
	passkey := r.URL.Query().Get("passkey")
	err := json.NewDecoder(r.Body).Decode(&cord)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	//credentials := a.db.First(&s, "email = ?", username)
	err = utils.DB.Raw("SELECT id FROM user3 WHERE email = ?", username).Scan(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	// a.db.where("username = ?",username)
	//fmt.Println(&s)
	data, err := json.Marshal(&s)

	if s.ID == "" {
		fmt.Println("User does not exist/registered")
		reply = models.LogInReply{AccessKey: "", RefreshKey: "", Msg: "User does not exist/registered", UserDetails: s, AllowUsers: false}
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
			reply = models.LogInReply{AccessKey: "", RefreshKey: "", Msg: "Password is incorrect", UserDetails: s, AllowUsers: false}
			err = json.NewEncoder(w).Encode(reply)
			if err != nil {
				sendErr(w, http.StatusInternalServerError, err.Error())
			}
		} else {
			fmt.Println("Login Sucessfull")
			reply = models.LogInReply{AccessKey: "", RefreshKey: "", Msg: "Login Sucessfull", UserDetails: s, AllowUsers: true, City: "Gainsvile"}
			err = json.NewEncoder(w).Encode(reply)
			if err != nil {
				sendErr(w, http.StatusInternalServerError, err.Error())
			}
		}
		fmt.Println(string(data))
		fmt.Println()
	}
}

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s models.User3
	reply_on_succ := models.SignInReply{Msg: "Sucessfull"}
	reply_on_fail := models.SignInReply{Msg: "Email already exists, try using another email."}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		// sendErr(w, http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(reply_on_fail)
		return
	}
	s.ID = uuid.New().String()
	err = utils.DB.Save(&s).Error
	if err != nil {
		// sendErr(w, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(reply_on_fail)
	} else {
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(reply_on_succ)
	}
}
