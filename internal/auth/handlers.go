package auth

import (
	"encoding/json"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var reqBody LoginRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	if reqBody.Username == "admin" && reqBody.Password == "password" {
		json.NewEncoder(w).Encode(map[string]string{
			"message":  "Login successful",
			"username": reqBody.Username,
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid username or password",
		})
	}

}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
