package main

import (
	"encoding/json"
	"net/http"
)

// SearchHandler принимает поисковой запрос (поисковые запросы должны выполнятся в параллель)
type SearchHandler struct {
	keyWord string
}

// Поиск по постам и их коментам
func (h *SearchHandler) search(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(h.keyWord)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// GET запрос
}

// Запуск хендлера. Выполняется в мейне
func searchHandler() {
	handler := &SearchHandler{
		keyWord: "",
	}

	http.HandleFunc("/posts/search", handler.search)
}
