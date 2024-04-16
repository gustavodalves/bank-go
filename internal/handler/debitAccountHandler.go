package web

import (
	"encoding/json"
	"kafka-go-microservice/internal/application/usecase"
	"net/http"
)

type DebitAccountHandler struct {
	UseCase usecase.DebitAccountUseCase
}

func (h *DebitAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var dto usecase.DebitAccountDTO

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
