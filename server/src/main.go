package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	//"os/user"
	"src/Carts"
	"src/Offers"
	"src/Stores"
	"src/Users"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//"strconv"
)

type student struct {
	ID       string `gorm:"primary_key" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type App struct {
	db *gorm.DB
	r  *mux.Router
}

type Cart_items_db struct {
	UserID     string `gorm:"-" json:"id"`
	ProductID  string `json:"productID"`
	Quantity   string `json:"quantity"`
	CreatedAt  string `json:"created"`
	ModifiedAt string `json:"modified"`
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func main() {

	db, err := gorm.Open(sqlite.Open("Users.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	app := App{
		db: db,
		r:  mux.NewRouter(),
	}
	app.start()
}

func (a *App) start() {
	a.db.AutoMigrate(&Users.User3{})
	a.db.AutoMigrate(Offers.Offer{})
	a.r.HandleFunc("/address", a.returnLat) //returning lat
	a.r.HandleFunc("/address/", a.returnNearBy)
	a.r.HandleFunc("/offers", a.returnOffers)
	a.r.HandleFunc("/user", a.userLogin).Methods("GET")
	a.r.HandleFunc("/user", a.userSignUp).Methods("POST")
	a.db.AutoMigrate(&Carts.Cart_items{})
	a.db.AutoMigrate(&Stores.Store_inventory{})
	a.r.HandleFunc("/address", a.returnLat) //returning lat
	a.r.HandleFunc("/stores/", a.returnNearBy)
	a.r.HandleFunc("/stores/add/{storeID}", a.addInventory).Methods("POST")
	a.r.HandleFunc("/stores/edit/{storeID}", a.editInventory).Methods("POST")
	a.r.HandleFunc("/stores/delete/", a.deleteInventory).Methods("POST")
	a.r.HandleFunc("/stores/items", a.returnStoreInv).Methods("POST")
	a.r.HandleFunc("/stores/items/{product_id}", a.returnProductPage)
	a.r.HandleFunc("/user", a.userLogin).Methods("POST")
	a.r.HandleFunc("/user", a.userSignUp).Methods("POST")
	a.r.HandleFunc("/user/forgotpassword", a.ForgotUserDetails).Methods("POST")
	a.r.HandleFunc("/userStatus", a.userStatus).Methods("POST")     //this
	a.r.HandleFunc("/userCheck", a.userStatusCheck).Methods("POST") //this
	a.r.HandleFunc("/cart", a.cartDisplay).Methods("POST")          //this
	a.r.HandleFunc("/cart/additem", a.cartAddition).Methods("POST") //this
	a.r.HandleFunc("/contact", a.contact).Methods("POST")           //this
	a.r.HandleFunc("/user", a.changeUserDetails).Methods("PUT")
	a.r.HandleFunc("/user/orders", a.sendUserOrders).Methods("POST")
	a.r.HandleFunc("/students/", a.getAllStudents).Methods("GET")
	a.r.HandleFunc("/students/", a.addStudent).Methods("POST")
	a.r.HandleFunc("/students/{id}", a.updateStudent).Methods("PUT")
	a.r.HandleFunc("/students/{id}", a.deleteStudent).Methods("DELETE")

	http.Handle("/", a.r)

	// Users -> main.go

	// latNlon -> main.go

	// nearBy -> main.go

	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(a.r)))
	// log.Fatal(http.ListenAndServe(":10000", a.r))
}

// func PaymentsAuthorization(w http.ResponseWriter, r *http.Request) {

// 	// Read body
// 	b, err := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	// Unmarshal
// 	var authorizationRequestDto dto.AuthorizationRequestDto
// 	err = json.Unmarshal(b, &authorizationRequestDto)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	// Prepare Payment Response
// 	w.Header().Set("content-type", "application/json")

// 	// Basic Validation - Business Account
// 	var businessAccountId = r.Header.Get("From")
// 	var businessAccount model.Account
// 	if len(businessAccountId) > 0 {
// 		businessAccount, _ = getAccount(db, businessAccountId)
// 	} else {
// 		json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(authorizationRequestDto.OrderId, "3", "Invalid Merchant"))
// 		return
// 	}
// 	if businessAccount.Id != businessAccountId {
// 		json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(authorizationRequestDto.OrderId, "15", "No Such Issuer"))
// 		return
// 	}

// 	// Basic Validation - Personal Account
// 	var personalAccountId = fmt.Sprintf("%v", authorizationRequestDto.CardNumber)
// 	var personalAccount model.Account
// 	if len(personalAccountId) > 0 {
// 		personalAccount, _ = getAccount(db, personalAccountId)
// 	} else {
// 		json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(authorizationRequestDto.OrderId, "12", "Invalid Card Number"))
// 		return
// 	}
// 	if personalAccount.Id != personalAccountId {
// 		json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(authorizationRequestDto.OrderId, "56", "No Card Record"))
// 		return
// 	}

// 	if personalAccount.CardNumber != authorizationRequestDto.CardNumber ||
// 		personalAccount.CardSecurityCode != authorizationRequestDto.CardSecurityCode ||
// 		personalAccount.CardExpiryYear != authorizationRequestDto.CardExpiryYear ||
// 		personalAccount.CardExpiryMonth != authorizationRequestDto.CardExpiryMonth {
// 		var payment = model.CreateAuthorizationPayment(authorizationRequestDto,
// 			personalAccount,
// 			businessAccount,
// 			"5",
// 			"Do Not Honour")
// 		savePayment(db, payment)
// 		businessAccount.Statement = append(businessAccount.Statement, payment.Id)
// 		saveAccount(db, businessAccount)
// 		json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(payment.Id, "5", "Do Not Honour"))
// 		return
// 	}

// 	if personalAccount.Available < authorizationRequestDto.Amount {
// 		var payment = model.CreateAuthorizationPayment(authorizationRequestDto,
// 			personalAccount,
// 			businessAccount,
// 			"51",
// 			"Insufficient Funds")
// 		savePayment(db, payment)
// 		businessAccount.Statement = append(businessAccount.Statement, payment.Id)
// 		saveAccount(db, businessAccount)
// 		json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(payment.Id, "51", "Insufficient Funds"))
// 		return
// 	}

// 	// Successful Payment
// 	personalAccount.Available = personalAccount.Available - authorizationRequestDto.Amount
// 	personalAccount.Blocked = personalAccount.Blocked + authorizationRequestDto.Amount
// 	saveAccount(db, personalAccount)
// 	businessAccount.Blocked = businessAccount.Blocked + authorizationRequestDto.Amount
// 	saveAccount(db, businessAccount)
// 	var payment = model.CreateAuthorizationPayment(authorizationRequestDto,
// 		personalAccount,
// 		businessAccount,
// 		"0",
// 		"Approved")
// 	savePayment(db, payment)
// 	businessAccount.Statement = append(businessAccount.Statement, payment.Id)
// 	saveAccount(db, businessAccount)
// 	personalAccount.Statement = append(personalAccount.Statement, payment.Id)
// 	saveAccount(db, personalAccount)
// 	json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(payment.Id, "0", "Approved"))

// 	return
// }

func (a *App) sendUserOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 Users.User3
	var s2 Users.User3
	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = a.db.Raw("SELECT ID,userOrders FROM user3 WHERE id = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(s1.Acesskey)
	fmt.Println(s2.Acesskey)
	fmt.Println()
	if s2.Acesskey == s1.Acesskey {
		err = a.db.Exec("UPDATE user3 SET firstname = ?, lastname = ?, email = ?, password = ? where ID = ?", s1.Firstname, s1.Lastname, s1.Email, s1.Password, s2.ID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := Users.SignInReply{Msg: "sucessfully changed your details"}
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

func (a *App) addInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store Stores.Store_inventory
	reply := Users.SignInReply{Msg: "sucessfull"}
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	store.StoreID = uuid.New().String()
	err = a.db.Save(&store).Error
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

func (a *App) contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store Stores.Store_inventory
	reply := Users.SignInReply{Msg: "sucessfull"}
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	store.StoreID = uuid.New().String()
	err = a.db.Save(&store).Error
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

func (a *App) editInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var storeUser Stores.Store_inventory
	//var storeDB Stores.Store_inventory
	err := json.NewDecoder(r.Body).Decode(&storeUser)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	// err = a.db.Raw("SELECT ID,acessKey FROM user3 WHERE StoreID = ? and ProductID = ?", s1.ID).Scan(&s2).Error
	// if err != nil {
	// 	sendErr(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// fmt.Println(s1.Acesskey)
	// fmt.Println(s2.Acesskey)
	// fmt.Println()
	//if s2.Acesskey == s1.Acesskey {
	err = a.db.Exec("UPDATE store_inventory SET ProductPrice = ?, ProductName = ?, Quantity = ?, ModifiedAt = ? WHERE StoreID = ? and ProductID = ?", storeUser.ProductPrice, storeUser.ProductName, storeUser.Quantity, storeUser.ModifiedAt, storeUser.StoreID, storeUser.ProductID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	reply := Users.SignInReply{Msg: "sucessfully changed your details"}
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) deleteInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var storeUser Stores.Store_inventory
	//var storeDB Stores.Store_inventory
	err := json.NewDecoder(r.Body).Decode(&storeUser)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	// err = a.db.Raw("SELECT ID,acessKey FROM user3 WHERE StoreID = ? and ProductID = ?", s1.ID).Scan(&s2).Error
	// if err != nil {
	// 	sendErr(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// fmt.Println(s1.Acesskey)
	// fmt.Println(s2.Acesskey)
	// fmt.Println()
	//if s2.Acesskey == s1.Acesskey {
	err = a.db.Exec("DELETE from store_inventory WHERE StoreID = ? and ProductID = ?", storeUser.StoreID, storeUser.ProductID).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	reply := Users.SignInReply{Msg: "sucessfully changed your details"}
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}

	//a.db.Raw("SELECT  FROM user3 WHERE acesskey = ?", username)
	// err = a.db.Save(&s).Error
	// if err != nil {
	// 	sendErr(w, http.StatusInternalServerError, err.Error())
	// }
	//}
}

// func PaymentsCapture(w http.ResponseWriter, r *http.Request) {

// 	// Read Path Parameters
// 	vars := mux.Vars(r)

// 	// Read body
// 	b, err := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	// Unmarshal
// 	var successiveRequestDto dto.SuccessiveRequestDto
// 	err = json.Unmarshal(b, &successiveRequestDto)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	w.Header().Set("content-type", "application/json")

// 	// Basic Validation - Business Account
// 	var businessAccountId = r.Header.Get("From")
// 	var referenceId = vars["authorization_id"]
// 	var businessAccount model.Account
// 	if len(businessAccountId) > 0 {
// 		businessAccount, _ = getAccount(db, businessAccountId)
// 	} else {
// 		json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(referenceId, "3", "Invalid Merchant"))
// 		return
// 	}
// 	if businessAccount.Id != businessAccountId {
// 		json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(referenceId, "15", "No Such Issuer"))
// 		return
// 	}

// 	// Check if previous payment exists
// 	successiveRequestDto.Type = constant.CAPTURE
// 	successiveRequestDto.ReferenceId = referenceId
// 	if len(referenceId) <= 0 {
// 		json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(referenceId, "12", "Invalid Transaction	"))
// 		return
// 	}
// 	var referencedPayment, _ = getPayment(db, referenceId)
// 	if referencedPayment.Id != referenceId {
// 		json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(referenceId, "12", "Invalid Transaction	"))
// 		return
// 	}

// 	// Create successive payment
// 	var successivePayment model.Payment
// 	if referencedPayment.Operation == constant.AUTHORIZATION && referencedPayment.Status == "0" {
// 		if referencedPayment.CurrentAmount < successiveRequestDto.Amount {
// 			successivePayment = model.CreateSuccessivePayment(successiveRequestDto,
// 				referencedPayment,
// 				"13",
// 				"Invalid Amount")
// 			savePayment(db, successivePayment)
// 			businessAccount.Statement = append(businessAccount.Statement, successivePayment.Id)
// 			saveAccount(db, businessAccount)
// 			json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(successivePayment.Id,
// 				"13",
// 				"Invalid Amount"))
// 			return
// 		} else {
// 			referencedPayment.CurrentAmount = referencedPayment.CurrentAmount - successiveRequestDto.Amount
// 			savePayment(db, referencedPayment)
// 			successivePayment = model.CreateSuccessivePayment(successiveRequestDto,
// 				referencedPayment,
// 				"0",
// 				"Approved")
// 			savePayment(db, successivePayment)
// 			var personalAccountId = fmt.Sprintf("%v", referencedPayment.CardNumber)
// 			var personalAccount, _ = getAccount(db, personalAccountId)
// 			personalAccount.Blocked = personalAccount.Blocked - successiveRequestDto.Amount
// 			personalAccount.Statement = append(personalAccount.Statement, successivePayment.Id)
// 			saveAccount(db, personalAccount)
// 			businessAccount.Blocked = businessAccount.Blocked - successiveRequestDto.Amount
// 			businessAccount.Available = businessAccount.Available + successiveRequestDto.Amount
// 			businessAccount.Statement = append(businessAccount.Statement, successivePayment.Id)
// 			saveAccount(db, businessAccount)
// 			json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(successivePayment.Id,
// 				"0",
// 				"Approved"))
// 			return
// 		}

// 	}

// 	json.NewEncoder(w).Encode(dto.CreatePaymentResponseDto(successivePayment.Id, "12", "Invalid Transaction"))
// }

func (a *App) userStatusCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 Users.User3
	var s2 Users.User3
	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = a.db.Raw("SELECT ID,acessKey FROM user3 WHERE id = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(s1.Acesskey)
	fmt.Println(s2.Acesskey)
	fmt.Println()
	if s2.Acesskey == s1.Acesskey {
		err = a.db.Exec("UPDATE user3 SET firstname = ?, lastname = ?, email = ?, password = ? where ID = ?", s1.Firstname, s1.Lastname, s1.Email, s1.Password, s2.ID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := Users.SignInReply{Msg: "sucessfully changed your details"}
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

func (a *App) userStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 Users.User3
	var s2 Users.User3
	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = a.db.Raw("SELECT ID,acessKey FROM user3 WHERE id = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(s1.Acesskey)
	fmt.Println(s2.Acesskey)
	fmt.Println()
	if s2.Acesskey == s1.Acesskey {
		err = a.db.Exec("UPDATE user3 SET firstname = ?, lastname = ?, email = ?, password = ? where ID = ?", s1.Firstname, s1.Lastname, s1.Email, s1.Password, s2.ID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := Users.SignInReply{Msg: "sucessfully changed your details"}
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

func (a *App) changeUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 Users.User3
	var s2 Users.User3
	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = a.db.Raw("SELECT ID,acessKey FROM user3 WHERE id = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(s1.Acesskey)
	fmt.Println(s2.Acesskey)
	fmt.Println()
	if s2.Acesskey == s1.Acesskey {
		err = a.db.Exec("UPDATE user3 SET firstname = ?, lastname = ?, email = ? where ID = ?", s1.Firstname, s1.Lastname, s1.Email, s2.ID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := Users.SignInReply{Msg: "sucessfully changed your details"}
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

func (a *App) changeUserAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s1 Users.User3
	var s2 Users.ChangeUserAddress
	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = a.db.Raw("SELECT ID,acessKey FROM user3 WHERE id = ?", s1.ID).Scan(&s2).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(s1.Acesskey)
	fmt.Println(s2.Acesskey)
	fmt.Println()
	if s2.Acesskey == s1.Acesskey {
		err = a.db.Exec("UPDATE user3 SET Address1= ? where ID = ?", s2.Address, s2.ID).Error //add s1.address insted of s1.firstname
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := Users.SignInReply{Msg: "sucessfully changed your delivary address"}
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

func (a *App) ForgotUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var email Users.RetrevalDetails
	var s Users.User3
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = a.db.Raw("SELECT id,email,acesskey FROM user3 WHERE email = ?", email.Email).Scan(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}

}

func (a *App) userLogin(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(statusCode: 200)
	//w.WriteHeader(statusCode: 200)
	w.Header().Set("Content-Type", "application/json")

	//params := mux.Vars(r)
	//username := params["username"]
	//fmt.Println(username)
	var s Users.User3
	var reply Users.LogInReply
	var cord Users.Coardinates

	username := r.URL.Query().Get("email")
	passkey := r.URL.Query().Get("passkey")
	err := json.NewDecoder(r.Body).Decode(&cord)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	//credentials := a.db.First(&s, "email = ?", username)
	err = a.db.Raw("SELECT id FROM user3 WHERE email = ?", username).Scan(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	// a.db.where("username = ?",username)
	//fmt.Println(&s)
	data, err := json.Marshal(&s)

	if s.ID == "" {
		fmt.Println("User does not exist/registered")
		reply = Users.LogInReply{AccessKey: "", RefreshKey: "", Msg: "User does not exist/registered", UserDetails: s, AllowUsers: false}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}

	} else {
		fmt.Println(s.ID)
		err = a.db.Raw("SELECT * FROM user3 WHERE id = ? AND password = ?", s.ID, passkey).Scan(&s).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		if s.Email == "" {
			fmt.Println("Password is incorrect")
			reply = Users.LogInReply{AccessKey: "", RefreshKey: "", Msg: "Password is incorrect", UserDetails: s, AllowUsers: false}
			err = json.NewEncoder(w).Encode(reply)
			if err != nil {
				sendErr(w, http.StatusInternalServerError, err.Error())
			}
		} else {
			fmt.Println("Login Sucessfull")
			reply = Users.LogInReply{AccessKey: "", RefreshKey: "", Msg: "Login Sucessfull", UserDetails: s, AllowUsers: true, City: "Gainsvile"}
			err = json.NewEncoder(w).Encode(reply)
			if err != nil {
				sendErr(w, http.StatusInternalServerError, err.Error())
			}
		}
		fmt.Println(string(data))
		fmt.Println()
	}
}

// func sendErr(w http.ResponseWriter, i int, s string) {
// 	panic("unimplemented")
// }

func (a *App) userSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s Users.User3
	reply_on_succ := Users.SignInReply{Msg: "Sucessfull"}
	reply_on_fail := Users.SignInReply{Msg: "Email already exists, try using another email."}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		// sendErr(w, http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(reply_on_fail)
		return
	}
	s.ID = uuid.New().String()
	err = a.db.Save(&s).Error
	if err != nil {
		// sendErr(w, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(reply_on_fail)
	} else {
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(reply_on_succ)
	}
}

func (a *App) returnOffers(w http.ResponseWriter, r *http.Request) {
	a.db.Model(&Offers.Offer{}).Create([]map[string]interface{}{
		{"name": "jinzhu_1", "description": "10% off on all items"},
		{"name": "jinzhu_2", "description": "20% off on all items"},
	})
	w.Header().Set("Content-Type", "application/json")
	var all []Offers.Offer
	err := a.db.Find(&all).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(all)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (p Cart_items_db) TableName() string {
	// double check here, make sure the table does exist!!
	if p.UserID != "" {
		return p.UserID
	}
	return "cart_items_db" // default table name
}

func (a *App) cartAddition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart Carts.Cart_items
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = a.db.Save(&cart).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (a *App) cartDisplay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart Carts.Cart_items
	var userID Carts.UserIDtab
	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = a.db.Raw("SELECT * FROM user3 WHERE userID = ?", userID).Scan(&cart).Error
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) returnNearBy(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("long")
	w.Header().Set("Content-Type", "application/json")

	keyword := search
	radius := "1500"
	field := "formatted_address,name,rating,opening_hours,geometry"
	location := lat + "," + long
	// fmt.Println(location)
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "keyword=" + url.QueryEscape(keyword) + "&" +
		"radius=" + url.QueryEscape(radius) + "&" +
		"field=" + url.QueryEscape(field) + "&" +
		"location=" + url.QueryEscape(location) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/place/nearbysearch/json?", params)
	// fmt.Println(path)
	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//data1 := result{}
	var f interface{}
	json.Unmarshal(body, &f)
	fmt.Println(f)

	json.NewEncoder(w).Encode(f)
	defer resp.Body.Close()
}

func (a *App) filterInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inv Carts.Cart_items
	var userID Carts.UserIDtab
	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = a.db.Raw("SELECT * FROM storesInventory WHERE storeID = ?", userID).Scan(&inv).Error
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(inv)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) returnStoreInv(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	keyword := "foods"
	radius := "1500"
	field := "formatted_address,name,rating,opening_hours,geometry"
	location := "29.61872,-82.37299"
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "keyword=" + url.QueryEscape(keyword) + "&" +
		"radius=" + url.QueryEscape(radius) + "&" +
		"field=" + url.QueryEscape(field) + "&" +
		"location=" + url.QueryEscape(location) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/place/nearbysearch/json?", params)
	fmt.Println(path)
	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//data1 := result{}
	var f interface{}
	json.Unmarshal(body, &f)
	fmt.Println(f)

	json.NewEncoder(w).Encode(f)
	defer resp.Body.Close()
}

func (a *App) returnProductPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	keyword := "foods"
	radius := "1500"
	field := "formatted_address,name,rating,opening_hours,geometry"
	location := "29.61872,-82.37299"
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "keyword=" + url.QueryEscape(keyword) + "&" +
		"radius=" + url.QueryEscape(radius) + "&" +
		"field=" + url.QueryEscape(field) + "&" +
		"location=" + url.QueryEscape(location) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/place/nearbysearch/json?", params)
	fmt.Println(path)
	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//data1 := result{}
	var f interface{}
	json.Unmarshal(body, &f)
	fmt.Println(f)

	json.NewEncoder(w).Encode(f)
	defer resp.Body.Close()
}

func (a *App) ConvAddressToCord(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	address := "1600+Amphitheatre+Parkway,+Mountain+View,+CA"
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "address=" + url.QueryEscape(address) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/geocode/json?", params)
	fmt.Println(path)
	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//data1 := result{}
	var f interface{}
	json.Unmarshal(body, &f)
	fmt.Println(f)

	json.NewEncoder(w).Encode(f)
	defer resp.Body.Close()
}

func (a *App) returnLat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	address := "1600+Amphitheatre+Parkway,+Mountain+View,+CA"
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "address=" + url.QueryEscape(address) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/geocode/json?", params)
	fmt.Println(path)
	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//data1 := result{}
	var f interface {
		getHtml() string
	}

	json.Unmarshal(body, &f)
	fmt.Println(f)

	json.NewEncoder(w).Encode(f)
	defer resp.Body.Close()
}

func (a *App) getAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var all []student
	err := a.db.Find(&all).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(all)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) addStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s student
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	s.ID = uuid.New().String()
	err = a.db.Save(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (a *App) updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s student
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	s.ID = mux.Vars(r)["id"]
	err = a.db.Save(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := a.db.Unscoped().Delete(student{ID: mux.Vars(r)["id"]}).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// // … // Existing code from above

// type Book struct {
// 	ID     string  `json:"id"`
// 	Isbn   string  `json:"isbn"`
// 	Title  string  `json:"title"`
// 	Author *Author `json:"author"`
// }

// // Author struct
// type Author struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

// // Init books var as a slice Book struct
// var books []Book

// // Get all books
// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// // Get single book
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) // Gets params
// 	// Loop through books and find one with the id from the params
// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }

// // Add new book
// func createBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)
// }

// // Update book
// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// }

// // Delete book
// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// type Article struct {
// 	Id      string `json:"Id"`
// 	Title   string `json:"Title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// // let's declare a global Articles array
// // that we can then populate in our main function
// // to simulate a database
// var Articles []Article

// func main() {
// 	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
// 	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
// 	handleRequests()
// }

// func handleRequests() {
// 	myRouter := mux.NewRouter().StrictSlash(true)

// 	// Route handles & endpoints
// 	myRouter.HandleFunc("/books", getBooks).Methods("GET")
// 	myRouter.HandleFunc("/books/{id}", getBook).Methods("GET")
// 	myRouter.HandleFunc("/books", createBook).Methods("POST")
// 	myRouter.HandleFunc("/books/{id}", updateBook).Methods("PUT")
// 	myRouter.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
// 	myRouter.HandleFunc("/", homePage)
// 	log.Fatal(http.ListenAndServe(":10000", myRouter))
// }
