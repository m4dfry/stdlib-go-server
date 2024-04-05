package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ItemHandler struct{}

func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var i Item
	err := decoder.Decode(&i)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Created: %v\n", i)

	_, err = w.Write([]byte("Item:" + strconv.Itoa(i.ID) + " created"))
	if err != nil {
		log.Fatalln(err)
	}
}

func (h *ItemHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Println("Received request for item:" + id)
	_, err := w.Write([]byte("Item:" + id + " requested"))
	if err != nil {
		log.Fatalln(err)
	}
}
