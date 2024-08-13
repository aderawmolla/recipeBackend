package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/machinebox/graphql"
	"golang.org/x/crypto/bcrypt"
	services "golang/start/go/config"
	"golang/start/queries"
	"net/http"
	"os"
	"time"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var reqBody requestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := queries.User
	req := graphql.NewRequest(query)
	req.Var("username", reqBody.Username)

	// Updated response structure
	var respData struct {
		Data struct {
			Users []struct {
				ID       string `json:"id"`
				Password string `json:"password"`
			} `json:"users"`
		} `json:"data"`
	}

	ctx := context.Background()
	if err := services.RunGraphQLQuery(ctx, req, &respData); err != nil {
		http.Error(w, "Failed to check username", http.StatusInternalServerError)
		return
	}

	if len(respData.Data.Users) > 0 {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := hashPassword(reqBody.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	mutation := queries.SignupMutation
	req = graphql.NewRequest(mutation)
	req.Var("username", reqBody.Username)
	req.Var("password", hashedPassword)

	var mutationResp struct {
		InsertUsers struct {
			Returning []struct {
				ID string `json:"id"`
			} `json:"returning"`
		} `json:"insert_users"`
	}

	if err := services.RunGraphQLQuery(ctx, req, &mutationResp); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User created successfully")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var reqBody requestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := queries.User
	req := graphql.NewRequest(query)
	req.Var("username", reqBody.Username)

	var respData struct {
		Users []struct {
			ID       string `json:"id"`
			Password string `json:"password"`
		} `json:"users"`
	}

	ctx := context.Background()
	if err := services.RunGraphQLQuery(ctx, req, &respData); err != nil {
		http.Error(w, "Failed to login", http.StatusInternalServerError)
		return
	}

	if len(respData.Users) == 0 || !checkPasswordHash(reqBody.Password, respData.Users[0].Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	userID := respData.Users[0].ID

	payload := jwt.MapClaims{
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles": []string{"admin", "user", "anonymous"},
			"x-hasura-default-role":  "user",
			"x-hasura-user-id":       userID,
		},
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(os.Getenv("HASURA_GRAPHQL_JWT_SECRET")))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    tokenString,
		Expires:  time.Now().Add(72 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Login successful",
	}
	json.NewEncoder(w).Encode(response)
}
