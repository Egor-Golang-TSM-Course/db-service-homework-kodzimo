package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Comment struct {
	ID           int       `json:"id"`
	PostID       int       `json:"post_id"`
	Author       User      `json:"author"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"creation_date"`
}

// CommentHandler хранит комментарии и обеспечивает потокобезопасный доступ
type CommentHandler struct {
	sync.Mutex
	comments map[int]Comment
	nextID   int
}

// Добавление комментария к посту
func (h *CommentHandler) createComment(comment Comment, postId int, w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.comments)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// POST запрос
}

// Получение комментариев к посту
func (h *CommentHandler) displayPostComments(postId int, w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.comments)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// GET запрос
}

// Запуск хендлера. Выполняется в мейне
func commentHandler() {
	handler := &CommentHandler{
		comments: make(map[int]Comment, 0),
		nextID:   1,
	}

	http.HandleFunc("/posts/{postId}/comments", handler.createComment(comment, postId))
	http.HandleFunc("/posts/{postId}/comments", handler.displayPostComments(postId))
}
