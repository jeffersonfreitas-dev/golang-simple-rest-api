package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"simple-rest-api/internal/usecases/products"
	"simple-rest-api/internal/usecases/users"
)

type Handlers struct {
	userCases    *users.UsersCase
	productCases *products.ProductCase
}

func New(userCases *users.UsersCase, productCases *products.ProductCase) *Handlers {
	return &Handlers{
		userCases:    userCases,
		productCases: productCases,
	}
}

func (h Handlers) Listen(port int) error {
	h.registerUserEndpoints()
	h.registerProductEndpoints()
	slog.Info("listening on", "port", port)
	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
