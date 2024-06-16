package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Eche5/rssagg/internal/auth"
	"github.com/Eche5/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error parsing JSON")
		return
	}
	newUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, "Error Creating User")
	}
	respondWithJSON(w, 201, databaseUserToUser(newUser))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error:%v", err))
		return
	}
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error fetching user details:%v", err))
		return
	}
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {

	_, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}
	users, err := apiCfg.DB.GetAllUsers(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error fetching user details:%v", err))
		return
	}
	respondWithJSON(w, 200, users)
}

func (apiCfg *apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "Missing user ID")
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	_, err = auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	err = apiCfg.DB.DeleteUser(r.Context(), userUUID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error deleting:%v", err))
		return
	}
	respondWithJSON(w, 200, "user deleted successfully")
}
