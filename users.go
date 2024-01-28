package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type User struct {
	ID   int    `json:"user_id"`
	Name string `json:"user_name"`
	Role Role   `json:"user_role"`
}

type Role struct {
	Read   bool
	Write  bool
	Delete bool
}

var Admin Role = Role{Read: true, Write: true, Delete: true}
var Member Role = Role{Read: true, Write: true, Delete: false}
var Guest Role = Role{Read: true, Write: false, Delete: false}

// UserHandler хранит пользователей и обеспечивает потокобезопасный доступ
type UserHandler struct {
	sync.Mutex
	users  map[int]User
	nextID int
}

func (h *UserHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.users)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здесь должна быть описана логика фунции
}

func (h *UserHandler) registerUser(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.users)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// Запрос POST
}

func (h *UserHandler) authenticateUser(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.users)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// Запрос POST
}

func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.users)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
}

func userHandler() {
	handler := &UserHandler{
		users:  make(map[int]User, 0),
		nextID: 1,
	}

	http.HandleFunc("/users", handler.getAllUsers)
	http.HandleFunc("/users/register", handler.registerUser)
	http.HandleFunc("/users/login", handler.authenticateUser)
	http.HandleFunc("/users/delete", handler.deleteUser)

}
