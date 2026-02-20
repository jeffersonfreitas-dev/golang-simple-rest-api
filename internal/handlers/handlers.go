package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	usescases "simple-rest-api/internal/usecases"
)

type Handlers struct {
	useCases *usescases.UseCases
}

func New(useCases *usescases.UseCases) *Handlers {
	return &Handlers{
		useCases: useCases,
	}
}

func (h Handlers) Listen(port int) error {
	h.registerUserEndpoints()
	slog.Info("listening on", "port", port)
	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
