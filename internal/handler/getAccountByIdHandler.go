package web

import (
	"encoding/json"
	"kafka-go-microservice/internal/application/usecase"
	"net/http"
)

type GetAccountByIdHandler struct {
	UseCase usecase.GetAccountUseCase
}

func (h *GetAccountByIdHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var dto *usecase.GetAccountDTO

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	account, err := h.UseCase.Execute(*dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
