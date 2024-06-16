package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Eche5/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params:=parameters{}
	err:=decoder.Decode(&params)
	if err !=nil{
		respondWithError(w,400,"Error parsing JSON")
		return
	}
	newUser ,err:=apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil{
		respondWithError(w,400,"Error Creating User")
	}
	respondWithJSON(w, 200, newUser)
}
