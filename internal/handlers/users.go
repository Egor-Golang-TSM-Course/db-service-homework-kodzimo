package handlers

import (
	"db-service-homework-kodzimo/internal/storage"
	"db-service-homework-kodzimo/internal/storage/postgres"
	"db-service-homework-kodzimo/pkg/config"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"sync"
	"time"
)

type Role struct {
	Read   bool
	Write  bool
	Delete bool
}

// UserHandlerStruct хранит пользователей и обеспечивает потокобезопасный доступ
type UserHandlerStruct struct {
	sync.Mutex
	users  []postgres.User
	nextID int
}

// swagger:route GET /payload payloads REST_Task
// Get payload with given ID.
// responses:
//
//	200: []User
//	500: ServerError
//	404: ServerError
func (h *UserHandlerStruct) getAllUsers(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()
	cfg := config.ReadEnv("")

	ll := logrus.New()
	st, err := storage.NewStorage(cfg, ll)

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
	cfg := config.ReadEnv("")

	ll := logrus.New()
	st, _ := storage.NewStorage(cfg, ll)

	switch r.Method {
	case "POST":
		var newUser postgres.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = fmt.Fprintf(w, "User: %+v", newUser.Name)
		if err != nil {
			return
		}
		err = st.Postgres.RegisterUser(newUser)
		if err != nil {
			return
		}
	case "GET":
		//err := st.Postgres.RegisterUser()
		//if err != nil {
		//	return
		//}
	}
	_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	if err != nil {
		return
	}

	// Здеcь должна быть описана логика фунции
	// Запрос POST
}

func (h *UserHandlerStruct) authenticateUser(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()
	cfg := config.ReadEnv("")

	ll := logrus.New()
	st, _ := storage.NewStorage(cfg, ll)

	// Здеcь должна быть описана логика фунции
	// Запрос POST

	type Credentials struct {
		Username string `json:"user_name"`
		Password string `json:"user_password"`
	}

	var creds Credentials

	// Декодируем учетные данные из тела запроса
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Проверяем учетные данные пользователя
	isValid := st.Postgres.CheckCredentials(creds.Username, creds.Password)

	if !isValid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Если учетные данные верны, создаем сессию
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "username", Value: creds.Username, Expires: expiration}
	http.SetCookie(w, &cookie)
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
		users:  make([]postgres.User, 0),
		nextID: 1,
	}

	http.HandleFunc("/users", handler.getAllUsers)
	http.HandleFunc("/users/register", handler.registerUser)
	http.HandleFunc("/users/login", handler.authenticateUser)
	//http.HandleFunc("/users/delete", handler.deleteUser)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
