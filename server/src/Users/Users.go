package Users

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
	r  *mux.Router
}

type SignInReply struct {
	Msg string
}

type LogInReply struct {
	AccessKey   string
	RefreshKey  string
	Msg         string
	UserDetails User3
}

type User3 struct {
	ID         string `gorm:"primary_key" json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Acesskey   string
	RefreshKey string
}

type RetrevalDetails struct {
	Email string `json:"email"`
}

type Cart_items struct {
	userID     string `gorm:"primary_key" json:"id"`
	productID  string `json:"productID"`
	quantity   string `json:"quantity"`
	createdAt  string `json:"created"`
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

// func (a *App) start() {
// 	a.db.AutoMigrate(&User3{})
// 	a.r.HandleFunc("/user", a.userLogin).Methods("GET")
// 	a.r.HandleFunc("/user/", a.userSignUp).Methods("POST")
// 	//a.r.HandleFunc("/students/{id}", a.updateStudent).Methods("PUT")
// 	//a.r.HandleFunc("/students/{id}", a.deleteStudent).Methods("DELETE")
// 	http.Handle("/", a.r)
// 	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(a.r)))
// 	// log.Fatal(http.ListenAndServe(":10000", a.r))
// }
