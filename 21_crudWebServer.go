package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
)

type dollars float32


func (d dollars) String() string {
	return fmt.Sprintf("%.2f", d)
}

type database map[string]dollars

func (d database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "Item -> %s price = %s",item, price )
	}
}

func (d database) add(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	fmt.Println(item,price)

	if _,ok := d[item]; ok {
		msg := fmt.Sprintf("%s item already exists", item)
		http.Error(w, msg,http.StatusBadRequest)
		return 
	}

	p,err := strconv.ParseFloat(price,32)
	if err != nil {
		msg := fmt.Sprintf("Error Parsing price %s", price)
		http.Error(w, msg,http.StatusBadRequest)
		return 
	}

	d[item] = dollars(p)
	fmt.Fprintf(w, "Successfully added item -> %s", item)

}

func main() {

	db := database{
		"shoes":50,
		"socks" : 10,
	}

	// Routes
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/add", db.add)

	log.Fatal(http.ListenAndServe(":8080", nil))



}