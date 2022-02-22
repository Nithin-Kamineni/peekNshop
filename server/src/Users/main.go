package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("../Users.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	app := App{
		db: db,
		r:  mux.NewRouter(),
	}

	app.start()
}

type App struct {
	db *gorm.DB
	r  *mux.Router
}

type SignInReply struct {
	Msg string
}

type LogInReply struct {
	AccessKey  string
	RefreshKey string
	Msg        string
}

type user3 struct {
	ID        string `gorm:"primary_key" json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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

func (a *App) start() {
	a.db.AutoMigrate(&user3{})
	a.r.HandleFunc("/user", a.userLogin).Methods("GET")
	a.r.HandleFunc("/user/", a.userSignUp).Methods("POST")
	//a.r.HandleFunc("/students/{id}", a.updateStudent).Methods("PUT")
	//a.r.HandleFunc("/students/{id}", a.deleteStudent).Methods("DELETE")
	http.Handle("/", a.r)
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(a.r)))
	// log.Fatal(http.ListenAndServe(":10000", a.r))
}

func (a *App) userLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	//username := params["username"]
	//fmt.Println(username)
	var s user3

	username := r.URL.Query().Get("email")
	passkey := r.URL.Query().Get("passkey")
	//credentials := a.db.First(&s, "email = ?", username)
	a.db.Raw("SELECT * FROM user3 WHERE email = ? AND password = ?", username, passkey).Scan(&s)
	// a.db.where("username = ?",username)
	//fmt.Println(&s)
	data, err := json.Marshal(&s)
	fmt.Printf(string(data))

	var all []user3
	err = a.db.Find(&all).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	//reply := LogInReply(AccessKey:"", RefreshKey:"", Msg:"")
	err = json.NewEncoder(w).Encode(all)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

// func sendErr(w http.ResponseWriter, i int, s string) {
// 	panic("unimplemented")
// }

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

func (a *App) userSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s user3
	reply := SignInReply{Msg: "sucessfull"}
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
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
