package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

// Item represents a simple data structure.
type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var items []Item

func main() {
    // Initialize a router using the Gorilla Mux router.
    router := mux.NewRouter()

    // Define API endpoints
    router.HandleFunc("/items", GetItems).Methods("GET")
    router.HandleFunc("/items/{id}", GetItem).Methods("GET")
    router.HandleFunc("/items", CreateItem).Methods("POST")
    router.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
    router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

    // Initialize some sample data
    items = append(items, Item{ID: "1", Name: "Item 1"})
    items = append(items, Item{ID: "2", Name: "Item 2"})

    // Start the server
    fmt.Println("Server is running on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

// GetItems returns a list of all items.
func GetItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// GetItem returns a single item by ID.
func GetItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range items {
        if item.ID == params["id"] {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
}

// CreateItem adds a new item to the list.
func CreateItem(w http.ResponseWriter, r *http.Request) {
    var item Item
    _ = json.NewDecoder(r.Body).Decode(&item)
	fmt.Println("Creating a item")
	fmt.Println(item.ID)
	fmt.Println(item.Name)
	if item.ID == "" {
		fmt.Println("ID not found")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if item.Name == "" {
		fmt.Println("Name not found")
		w.WriteHeader(http.StatusNoContent)
		return
	}
    items = append(items, item)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

// UpdateItem updates an existing item by ID.
func UpdateItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for i, item := range items {
        if item.ID == params["id"] {
            var updatedItem Item
            _ = json.NewDecoder(r.Body).Decode(&updatedItem)
            items[i] = updatedItem
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(updatedItem)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
}

// DeleteItem removes an item by ID.
func DeleteItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for i, item := range items {
        if item.ID == params["id"] {
            items = append(items[:i], items[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
}
