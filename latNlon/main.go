package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type result struct {
	results []struct{ results_info } `json:"results"`
	status  string                   `json:"status"`
}

type results_info struct {
	address_component []address_component_info `json:"address_component"`
	formated_address  string                   `json:"formated_address"`
	geometry          []geometry_info          `json:"geometry"`
	place_id          string                   `json:"place_id"`
	types             []string                 `json:"types"`
}

type address_component_info struct {
	long_name  string   `json:"long_name"`
	short_name string   `json:"short_name"`
	types      []string `json:"types"`
}

type geometry_info struct {
	bounds        []bounds_info `json:"bounds"`
	location      []cord_info   `json:"location"`
	location_type string        `json:"location_type"`
	viewport      []bounds_info `json:"viewport"`
}

type bounds_info struct {
	northeast []cord_info `json:"northeast"`
	southwest []cord_info `json:"southwest"`
}

type cord_info struct {
	lat float64 `json:"lat"`
	lng float64 `json:"lng"`
}

func main() {
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/address", returnLat)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnLat(w http.ResponseWriter, r *http.Request) {

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
