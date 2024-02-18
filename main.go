package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// Item struct represents an item with ID, Name, and Description
type Item struct {
    ID          string `json:"id,omitempty"`       // ID of the item
    Name        string `json:"name,omitempty"`     // Name of the item
    Description string `json:"description,omitempty"`  // Description of the item
}

var items []Item // Slice to store items

// GetItems handles GET request to retrieve all items
func GetItems(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(items) // Encode items slice to JSON and send as response
}

// GetItem handles GET request to retrieve a specific item by ID
func GetItem(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req) // Get request parameters
    // Loop through items to find matching ID
    for _, item := range items {
        if item.ID == params["id"] { // If ID matches
            json.NewEncoder(w).Encode(item) // Encode matching item to JSON and send as response
            return
        }
    }
    json.NewEncoder(w).Encode(&Item{}) // If no matching ID found, send an empty JSON object
}

// CreateItem handles POST request to create a new item
func CreateItem(w http.ResponseWriter, req *http.Request) {
    var item Item // Create a new Item object
    _ = json.NewDecoder(req.Body).Decode(&item) // Decode JSON request body into the item object
    items = append(items, item) // Append the new item to the items slice
    json.NewEncoder(w).Encode(item) // Encode the new item to JSON and send as response
}

func main() {
    router := mux.NewRouter() // Create a new router using gorilla/mux

    // Initialize some example items
    items = append(items, Item{ID: "1", Name: "Item 1", Description: "Description of Item 1"})
    items = append(items, Item{ID: "2", Name: "Item 2", Description: "Description of Item 2"})

    // Define route handlers for each endpoint
    router.HandleFunc("/items", GetItems).Methods("GET")    // GET /items endpoint
    router.HandleFunc("/items/{id}", GetItem).Methods("GET") // GET /items/{id} endpoint
    router.HandleFunc("/items", CreateItem).Methods("POST")  // POST /items endpoint

    // Start the HTTP server on port 8000 using the router
    log.Fatal(http.ListenAndServe(":8000", router))
}
