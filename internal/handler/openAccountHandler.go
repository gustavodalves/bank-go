package web

import (
	"encoding/json"
	"kafka-go-microservice/internal/application/usecase"
	"net/http"
)

type OpenAccountHandler struct {
	UseCase *usecase.OpenAccountUseCase
}

func (h *OpenAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OpenAccountDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	err := h.UseCase.Execute(dto)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
