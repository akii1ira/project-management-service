package handlers

import (
	"encoding/json"
	"net/http"
	"project-management-service/db"
	"project-management-service/models"
	"time"
	"github.com/go-chi/chi/v5"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, title, description, start_date, end_date, manager FROM projects")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		var endDate *time.Time
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.StartDate, &endDate, &p.ManagerID); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		p.EndDate = endDate
		projects = append(projects, p)
	}
	json.NewEncoder(w).Encode(projects)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var p models.Project
	var endDate *time.Time
	err := db.DB.QueryRow("SELECT id, title, description, start_date, end_date, manager FROM projects WHERE id=$1", id).
		Scan(&p.ID, &p.Title, &p.Description, &p.StartDate, &endDate, &p.ManagerID)
	if err != nil {
		http.Error(w, "Project not found", 404)
		return
	}
	p.EndDate = endDate
	json.NewEncoder(w).Encode(p)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid data", 400)
		return
	}
	err := db.DB.QueryRow("INSERT INTO projects(title, description, start_date, manager) VALUES($1,$2,$3,$4) RETURNING id",
		p.Title, p.Description, p.StartDate, p.ManagerID).Scan(&p.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(p)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var p models.Project
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid data", 400)
		return
	}
	_, err := db.DB.Exec("UPDATE projects SET title=$1, description=$2, start_date=$3, end_date=$4, manager=$5 WHERE id=$6",
		p.Title, p.Description, p.StartDate, p.EndDate, p.ManagerID, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(p)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := db.DB.Exec("DELETE FROM projects WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Project deleted successfully"))
}
