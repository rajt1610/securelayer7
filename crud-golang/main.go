package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "strings"

    _ "modernc.org/sqlite"
)

type Item struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("sqlite", "crud.db")
    if err != nil {
        log.Fatalf("open db: %v", err)
    }
    defer db.Close()

    if err := ensureSchema(); err != nil {
        log.Fatalf("schema: %v", err)
    }

    http.HandleFunc("/items", itemsHandler)
    http.HandleFunc("/items/", itemHandler)

    addr := ":8080"
    log.Printf("Starting server on %s (database: crud.db)", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}

func ensureSchema() error {
    const q = `CREATE TABLE IF NOT EXISTS items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT
    );`
    _, err := db.Exec(q)
    return err
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getItems(w, r)
    case http.MethodPost:
        createItem(w, r)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
    id, err := parseID(r.URL.Path)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        getItem(w, r, id)
    case http.MethodPut:
        updateItem(w, r, id)
    case http.MethodDelete:
        deleteItem(w, r, id)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}

func parseID(path string) (int, error) {
    p := strings.TrimPrefix(path, "/items/")
    p = strings.Trim(p, "/")
    return strconv.Atoi(p)
}

func getItems(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, description FROM items")
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    list := []Item{}
    for rows.Next() {
        var it Item
        if err := rows.Scan(&it.ID, &it.Name, &it.Description); err != nil {
            http.Error(w, "internal error", http.StatusInternalServerError)
            return
        }
        list = append(list, it)
    }
    respondJSON(w, http.StatusOK, list)
}

func getItem(w http.ResponseWriter, r *http.Request, id int) {
    var it Item
    row := db.QueryRow("SELECT id, name, description FROM items WHERE id = ?", id)
    if err := row.Scan(&it.ID, &it.Name, &it.Description); err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "not found", http.StatusNotFound)
            return
        }
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    respondJSON(w, http.StatusOK, it)
}

func createItem(w http.ResponseWriter, r *http.Request) {
    var in Item
    if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    res, err := db.Exec("INSERT INTO items(name, description) VALUES(?, ?)", in.Name, in.Description)
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    id64, err := res.LastInsertId()
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    in.ID = int(id64)
    respondJSON(w, http.StatusCreated, in)
}

func updateItem(w http.ResponseWriter, r *http.Request, id int) {
    var in Item
    if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    res, err := db.Exec("UPDATE items SET name = ?, description = ? WHERE id = ?", in.Name, in.Description, id)
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    if n, _ := res.RowsAffected(); n == 0 {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    in.ID = id
    respondJSON(w, http.StatusOK, in)
}

func deleteItem(w http.ResponseWriter, r *http.Request, id int) {
    res, err := db.Exec("DELETE FROM items WHERE id = ?", id)
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    if n, _ := res.RowsAffected(); n == 0 {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func respondJSON(w http.ResponseWriter, status int, v interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if v == nil {
        return
    }
    _ = json.NewEncoder(w).Encode(v)
}
