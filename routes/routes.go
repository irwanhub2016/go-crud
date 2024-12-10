package routes

import (
    "github.com/gorilla/mux"
    "github.com/irwanhub2016/go-crud/handlers"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
    router.HandleFunc("/seed", handlers.SeedHandler).Methods("POST")

    return router
}
