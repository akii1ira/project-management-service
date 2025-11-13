package handlers

import (
"database/sql"
"encoding/json"
"net/http"
"project-management-service/db"
"project-management-service/models"
"strconv"

"github.com/go-chi/chi/v5"
)
func GetUsers(w http.ResponseWriter, r *http.Request) {
rows, err := db.DB.Query("SELECT id, name, email, role, registered_at FROM users")
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
defer rows.Close()

users := []models.User{}
for rows.Next() {
var u models.User
if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.RegisteredAt); err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
users = append(users, u)
}
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
id := chi.URLParam(r, "id")
row := db.DB.QueryRow("SELECT id, name, email, role, registered_at FROM users WHERE id=$1", id)
var u models.User
err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.RegisteredAt)
if err != nil {
if err == sql.ErrNoRows {
http.Error(w, "User not found", http.StatusNotFound)
return
}
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(u)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
var u models.User
if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
http.Error(w, "Invalid data", http.StatusBadRequest)
return
}
if u.Name == "" || u.Email == "" || u.Role == "" {
http.Error(w, "Missing required fields", http.StatusBadRequest)
return
}
err := db.DB.QueryRow(
"INSERT INTO users(name, email, role) VALUES($1, $2, $3) RETURNING id, registered_at",
u.Name, u.Email, u.Role).Scan(&u.ID, &u.RegisteredAt)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(u)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
id := chi.URLParam(r, "id")
var u models.User
if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
http.Error(w, "Invalid data", http.StatusBadRequest)
return
}
res, err := db.DB.Exec("UPDATE users SET name=$1, email=$2, role=$3 WHERE id=$4",
u.Name, u.Email, u.Role, id)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
rows, _ := res.RowsAffected()
if rows == 0 {
http.Error(w, "User not found", http.StatusNotFound)
return
}
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(u)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
id := chi.URLParam(r, "id")
res, err := db.DB.Exec("DELETE FROM users WHERE id=$1", id)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
rows, _ := res.RowsAffected()
if rows == 0 {
http.Error(w, "User not found", http.StatusNotFound)
return
}
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(map[string]string{"result": "user deleted"})
}

// GetTasksByUser - возвращает задачи, назначенные пользователю (assignee)
func GetTasksByUser(w http.ResponseWriter, r *http.Request) {
id := chi.URLParam(r, "id")
rows, err := db.DB.Query("SELECT id, title, description, priority, status, assignee, project, created_at, completed_at FROM tasks WHERE assignee=$1", id)
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
func SearchUsers(w http.ResponseWriter, r *http.Request) {
qName := r.URL.Query().Get("name")
qEmail := r.URL.Query().Get("email")

if qName == "" && qEmail == "" {
http.Error(w, "No search parameters provided", http.StatusBadRequest)
return
}

var rows *sql.Rows
var err error
if qName != "" && qEmail != "" {
rows, err = db.DB.Query("SELECT id, name, email, role, registered_at FROM users WHERE name ILIKE $1 AND email ILIKE $2", "%"+qName+"%", "%"+qEmail+"%")
} else if qName != "" {
rows, err = db.DB.Query("SELECT id, name, email, role, registered_at FROM users WHERE name ILIKE $1", "%"+qName+"%")
} else {
rows, err = db.DB.Query("SELECT id, name, email, role, registered_at FROM users WHERE email ILIKE $1", "%"+qEmail+"%")
}
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
defer rows.Close()

users := []models.User{}
for rows.Next() {
var u models.User
if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.RegisteredAt); err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
users = append(users, u)
}
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(users)
}