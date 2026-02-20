package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-rest-api/internal/models"
)

func main() {
	req := models.CreateUserRequest{
		Name:  "Jeff",
		Email: "jeff@email.com",
	}

	b, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8090/users", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusCreated {
		panic("error to create user")
	}

	var responseApi models.CreateUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseApi); err != nil {
		panic(err)
	}

	fmt.Println("new user created", responseApi)
}
