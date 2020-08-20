package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item specified\n")
		return
	}
	pricestr := req.FormValue("price")
	price, err := strconv.Atoi(pricestr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price %v\n", pricestr)
		return
	}
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}

	db[item] = dollars(price)
	fmt.Fprintf(w, "item created %s: %s\n", item, dollars(price))
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item specified\n")
		return
	}

	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item specified\n")
		return
	}
	pricestr := req.FormValue("price")
	price, err := strconv.Atoi(pricestr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price %v\n", pricestr)
		return
	}

	db[item] = dollars(price)
	fmt.Fprintf(w, "item updated %s: %s\n", item, dollars(price))
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item specified\n")
		return
	}
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item %q\n", item)
		return
	}

	delete(db, item)
}

func main() {
	db := database{"shoes": 50, "socks": 5, "chocolates": 99, "milk": 46}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.price)
	http.HandleFunc("/update", db.price)
	http.HandleFunc("/delete", db.price)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
