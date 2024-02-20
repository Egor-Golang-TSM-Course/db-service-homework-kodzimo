package development

import (
	"db-service-homework-kodzimo/internal/handlers"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Post struct {
	ID           int           `json:"post_id"`
	Author       handlers.User `json:"author"`
	Header       string        `json:"post_header"`
	Body         string        `json:"post_body"`
	CreationDate time.Time     `json:"creation_date"`
}

// PostHandler хранит посты и обеспечивает потокобезопасный доступ
type PostHandler struct {
	sync.Mutex
	posts  map[int]Post
	nextID int
}

func (h *PostHandler) createPost(post Post, w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.posts)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// Запрос POST
}

func (h *PostHandler) displayAllPosts(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.posts)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// Запрос GET
}

// Получение детальной информации о конкретном посте.
func (h *PostHandler) displayPost(id int, w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.posts)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// Запрос GET
}

func (h *PostHandler) updatePost(id int, w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.posts)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// Запрос PUT
}

func (h *PostHandler) deletePost(id int, w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.posts)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// Запрос DELETE
}

// Запуск хендлера. Выполняется в мейне
func postHandler() {
	handler := &PostHandler{
		posts:  make(map[int]Post, 0),
		nextID: 1,
	}

	http.HandleFunc("/posts", handler.createPost(post))
	http.HandleFunc("/posts", handler.displayAllPosts)
	http.HandleFunc("/posts/{id}", handler.displayPost(id))
	http.HandleFunc("/posts/{id}", handler.updatePost(id))
	http.HandleFunc("/posts/{id}", handler.deletePost(id))
}
