package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/irwanhub2016/go-crud/config"
    "github.com/irwanhub2016/go-crud/seeders"
    "github.com/irwanhub2016/go-crud/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    rows, err := config.DB.Query("SELECT id, name, email, age FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    err := config.DB.QueryRow("INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id",
        user.Name, user.Email, user.Age).Scan(&user.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    _, err := config.DB.Exec("UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4",
        user.Name, user.Email, user.Age, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    _, err := config.DB.Exec("DELETE FROM users WHERE id = $1", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func SeedHandler(w http.ResponseWriter, r *http.Request) {
    // Connect to the database
    config.Connect()
    defer config.DB.Close()

    // Call the seed function to insert random users
    err := seeders.SeedUsers(config.DB)
    if err != nil {
        log.Printf("Error seeding users: %v", err)
        http.Error(w, "Failed to seed users", http.StatusInternalServerError)
        return
    }

    // Respond with a success message
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Seeding complete! 10 users added.\n"))
}