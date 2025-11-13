package handlers

import (
"database/sql"
"encoding/json"
"net/http"
"project-management-service/db"
"project-management-service/models"
"strconv"
"time"

"github.com/go-chi/chi/v5"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
rows, err := db.DB.Query("SELECT id, title, description, priority, status, assignee, project, created_at, completed_at FROM tasks")
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
defer rows.Close()

var tasks []models.Task
for rows.Next() {
var t models.Task
var completedAt *time.Time
if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Priority, &t.Status, &t.AssigneeID, &t.ProjectID, &t.CreatedAt, &completedAt); err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
t.CompletedAt = completedAt
tasks = append(tasks, t)
}
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(tasks)
}
func GetTask(w http.ResponseWriter, r *http.Request) {
id := chi.URLParam(r, "id")
var t models.Task
var completedAt *time.Time
err := db.DB.QueryRow("SELECT id, title, description, priority, status, assignee, project, created_at, completed_at FROM tasks WHERE id=$1", id).
Scan(&t.ID, &t.Title, &t.Description, &t.Priority, &t.Status, &t.AssigneeID, &t.ProjectID, &t.CreatedAt, &completedAt)
if err != nil {
if err == sql.ErrNoRows {
http.Error(w, "Task not found", http.StatusNotFound)
return
}
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
t.CompletedAt = completedAt
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(t)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
var t models.Task
if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
http.Error(w, "Invalid data", http.StatusBadRequest)
return
}
if t.Title == "" || t.Priority == "" || t.Status == "" {
http.Error(w, "Missing required fields", http.StatusBadRequest)
return
}
err := db.DB.QueryRow("INSERT INTO tasks(title, description, priority, status, assignee, project) VALUES($1,$2,$3,$4,$5,$6) RETURNING id, created_at",
t.Title, t.Description, t.Priority, t.Status, t.AssigneeID, t.ProjectID).Scan(&t.ID, &t.CreatedAt)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(t)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var t models.Task
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := `UPDATE tasks SET title = ?, description = ?, status = ?, user_id = ?, project_id = ?, updated_at = NOW() WHERE id = ?`
    _, err := database.DB.Exec(query, t.Title, t.Description, t.Status, t.UserID, t.ProjectID, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    query := `DELETE FROM tasks WHERE id = ?`
    _, err := database.DB.Exec(query, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    query := `SELECT id, title, description, status, user_id, project_id, created_at, updated_at FROM tasks WHERE id = ?`
    row := database.DB.QueryRow(query, id)

    var t models.Task
    if err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID, &t.ProjectID, &t.CreatedAt, &t.UpdatedAt); err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Task not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(t)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
    query := `SELECT id, title, description, status, user_id, project_id, created_at, updated_at FROM tasks ORDER BY created_at DESC`
    rows, err := database.DB.Query(query)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var t models.Task
        if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID, &t.ProjectID, &t.CreatedAt, &t.UpdatedAt); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tasks = append(tasks, t)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}
