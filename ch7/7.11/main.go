package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{sync.Mutex{}, map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.list)     // list
	http.HandleFunc("/price", db.price)   // R of CRUD
	http.HandleFunc("/create", db.create) // C of CRUD
	http.HandleFunc("/update", db.update) // U of CRUD
	http.HandleFunc("/delete", db.delete) // D of CRUD
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

////////////////////////////////
// price unit
type dollars float64

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

///////////////////////////////
// db with sync
type database struct {
	sync.Mutex
	prices map[string]dollars
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.prices {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.prices[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) create(w http.ResponseWriter, r *http.Request) {
	item, err := db.getItem(r, "item")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	strPrice := r.FormValue("price")
	price, err := strconv.ParseFloat(strPrice, 64)
	if err != nil {
		http.Error(w, "price required and must be a valid number", http.StatusBadRequest)
		return
	}

	if _, ok := db.prices[item]; ok {
		http.Error(w, fmt.Sprintf("%s already exists", item), http.StatusBadRequest)
		return
	}

	db.Lock()
	db.prices[item] = dollars(price)
	db.Unlock()
}

func (db *database) update(w http.ResponseWriter, r *http.Request) {
	item, err := db.getItem(r, "item")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	strPrice := r.FormValue("price")
	price, err := strconv.ParseFloat(strPrice, 64)
	if err != nil {
		http.Error(w, "price required and must be a valid number", http.StatusBadRequest)
		return
	}

	if _, ok := db.prices[item]; !ok {
		http.Error(w, fmt.Sprintf("%s does not exist", item), http.StatusNotFound)
		return
	}

	db.Lock()
	db.prices[item] = dollars(price)
	db.Unlock()
}

func (db *database) delete(w http.ResponseWriter, r *http.Request) {
	item, err := db.getItem(r, "item")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := db.prices[item]; !ok {
		http.Error(w, fmt.Sprintf("%s does not exist", item), http.StatusNotFound)
		return
	}

	db.Lock()
	delete(db.prices, item)
	db.Unlock()
}

func (db *database) getItem(r *http.Request, name string) (item string, err error) {
	requestItem := r.FormValue(name)
	if requestItem == "" {
		return "", errors.New("item required")
	}
	return requestItem, nil
}
