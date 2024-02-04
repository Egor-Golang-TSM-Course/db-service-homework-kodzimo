package handlers

import (
	"db-service-homework-kodzimo/internal/storage"
	"db-service-homework-kodzimo/pkg/config"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
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

// UserHandlerStruct хранит пользователей и обеспечивает потокобезопасный доступ
type UserHandlerStruct struct {
	sync.Mutex
	users  map[int]User
	nextID int
}

// swagger:route GET /payload payloads REST_Task
// Get payload with given ID.
// responses:
//
//	200: map[int]User
//	500: ServerError
//	404: ServerError
func (h *UserHandlerStruct) getAllUsers(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()
	cfg := config.ReadEnv("")

	ll := logrus.New()
	st, err := storage.NewStorage(cfg, ll)

	// Здесь должна быть описана логика фунции
	users, err := st.Postgres.GetAllUsers()
	if err != nil {
		defer println("getAllUsers func: ")
		ll.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandlerStruct) registerUser(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandlerStruct) authenticateUser(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandlerStruct) deleteUser(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.users)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
}

func UserHandler() {
	handler := &UserHandlerStruct{
		users:  make(map[int]User),
		nextID: 1,
	}

	http.HandleFunc("/users", handler.getAllUsers)
	//http.HandleFunc("/users/register", handler.registerUser)
	//http.HandleFunc("/users/login", handler.authenticateUser)
	//http.HandleFunc("/users/delete", handler.deleteUser)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
