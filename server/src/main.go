package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	a.db.AutoMigrate(&student{})
	a.r.HandleFunc("/address", a.returnLat)
	a.r.HandleFunc("/address/", a.returnNearBy)
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

func (a *App) returnNearBy(w http.ResponseWriter, r *http.Request) {

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
	var f interface{}
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
