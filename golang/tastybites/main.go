package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

/* You are assigned to prepare a backend system for TastyBites, a restaurant food
ordering platform. The application should allow users to order food online, reserve
tables, and let the admin manage bookings and view billing details.
Your task is to design and implement a system based on the following requirements:
*/

/* 1. Model Design:
a. Prepare a backend data model for a food ordering system.
b. The model should include relevant entities such as User, Table, Order,
and MenuItem.
c. Clearly show relationships between the entities (e.g., one user can place
multiple orders, one table can be booked by one user at a time).  */

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Table struct {
	ID       int  `json:"id"`
	Reserved bool `json:"reserved"`
	UserID   int  `json:"user_id"`
	OrderID  int  `json:"order_id"`
	Status  string  `json:"status"`
}

type MenuItem struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID      int        `json:"id"`
	UserID  int        `json:"user_id"`
	TableID int        `json:"table_id"`
	Items   []MenuItem `json:"items"`
	Total   float64    `json:"total"`
	Status  string     `json:"status"`
}

// =====================
// In-memory storage
// =====================

var users = map[int]User{}
var tables = map[int]*Table{}
var menu = map[int]MenuItem{}
var orders = map[int]*Order{}


var orderCounter = 1

// =====================
// Initialization
// =====================

func init() {
	// Create 20 tables
	for i := 1; i <= 20; i++ {
		tables[i] = &Table{ID: i}
	}

	// Sample users
	users[1] = User{ID: 1, Name: "user1"}
	users[2] = User{ID: 2, Name: "user2"}

	// Sample menu
	menu[1] = MenuItem{ID: 1, Name: "Burger", Price: 120}
	menu[2] = MenuItem{ID: 2, Name: "Pizza", Price: 250}
	menu[3] = MenuItem{ID: 3, Name: "Pasta", Price: 180}
	menu[4] = MenuItem{ID: 4, Name: "Cold Drink", Price: 60}
}

/*
2. API Operation:
a. Implement a POST operation that allows a user to place an order by
selecting items from the menu.
b. The order should be associated with a specific table reservation.

*/
func updateTableStatus() {
	for _, table := range tables {

		if !table.Reserved {
			table.Status = "AVAILABLE"

		} else if table.Reserved && table.OrderID == 0 {
			table.Status = "RESERVED"

		} else if table.Reserved && table.OrderID != 0 {
			table.Status = "ACTIVE"
		}
	}
}


func alltables(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	updateTableStatus()
	
	// order := orders[table.OrderID]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tables)
}


func allusers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// order := orders[table.OrderID]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}



func menus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// order := orders[table.OrderID]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}




// Reserve a table
func reserveTable(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		UserID  int `json:"user_id"`
		TableID int `json:"table_id"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	table, exists := tables[req.TableID]
	if !exists {
		http.Error(w, "Table not found", http.StatusNotFound)
		return
	}

	if table.Reserved {
		http.Error(w, "Table already reserved", http.StatusBadRequest)
		return
	}

	table.Reserved = true
	table.UserID = req.UserID

	w.Write([]byte("Table reserved successfully"))
}

/*
3. Table Management:
a. Consider that the restaurant has 20 tables, each with a unique table ID.
b. Once a user books a table, that specific table should be blocked (marked
as reserved) and unavailable for others until the order is completed or
manually released.
*/
func placeOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		UserID  int   `json:"user_id"`
		TableID int   `json:"table_id"`
		Items   []int `json:"items"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	table, exists := tables[req.TableID]
	if !exists {
		http.Error(w, "Table not found", http.StatusNotFound)
		return
	}

	if !table.Reserved || table.UserID != req.UserID {
		http.Error(w, "Table not reserved by this user", http.StatusBadRequest)
		return
	}

	var orderItems []MenuItem
	var total float64

	for _, itemID := range req.Items {
		item, ok := menu[itemID]
		if !ok {
			http.Error(w, "Invalid menu item", http.StatusBadRequest)
			return	
		}
		orderItems = append(orderItems, item)
		total += item.Price
	}

	order := &Order{
		ID:      orderCounter,
		UserID:  req.UserID,
		TableID: req.TableID,
		Items:   orderItems,
		Total:   total,
		Status:  "ACTIVE",
	}

	orders[orderCounter] = order
	table.OrderID = orderCounter
	orderCounter++

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

/*
4. Admin Access:
a. Implement an admin API to access any table’s booking details.
b. When the admin accesses a table (e.g., table 1), they should receive a
summary of the amount that the user needs to pay for the order
associated with that table.
*/
func adminTableDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/admin/table/")
	tableID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid table ID", http.StatusBadRequest)
		return
	}

	table, exists := tables[tableID]
	if !exists {
		http.Error(w, "Table not found", http.StatusNotFound)
		return
	}

	if table.OrderID == 0 {
		w.Write([]byte("No active order for this table"))
		return
	}

	order := orders[table.OrderID]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}




func Orders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}


// =====================
// Main
// =====================

func main() {
	// API endpoints show all tables
	http.HandleFunc("/all_tables", alltables)
	// API endpoints show all users
	http.HandleFunc("/all_users", allusers)
	// API endpoints show menu
	http.HandleFunc("/menu", menus)
	// API endpoints for reserve
	http.HandleFunc("/reserve", reserveTable)




	http.HandleFunc("/order", placeOrder)

	http.HandleFunc("/all_orders", Orders)





	http.HandleFunc("/admin/table/", adminTableDetails)

	http.ListenAndServe(":8080", nil)
}
