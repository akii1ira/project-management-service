package routes

import (
    "net/http"

    "github.com/gorilla/mux"
    "project-management-service/handlers"
    "project-management-service/middleware"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    // Middleware
    r.Use(middleware.SetJSONHeader)

    // Root
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Project Management Service API"))
    }).Methods("GET")

    // --- Users ---
    r.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

    // --- Projects ---
    r.HandleFunc("/projects", handlers.GetAllProjects).Methods("GET")
    r.HandleFunc("/projects/{id}", handlers.GetProjectByID).Methods("GET")
    r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
    r.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")
    r.HandleFunc("/projects/{id}", handlers.DeleteProject).Methods("DELETE")

    // --- Tasks ---
    r.HandleFunc("/tasks", handlers.GetAllTasks).Methods("GET")
    r.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")
    r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

    // --- Relationships ---
    r.HandleFunc("/users/{id}/tasks", handlers.GetTasksByUserID).Methods("GET")
    r.HandleFunc("/projects/{id}/tasks", handlers.GetTasksByProjectID).Methods("GET")

    // --- Search ---
    r.HandleFunc("/search/users", handlers.SearchUsers).Methods("GET")
    r.HandleFunc("/search/tasks", handlers.SearchTasks).Methods("GET")
    r.HandleFunc("/search/projects", handlers.SearchProjects).Methods("GET")

    return r
}
