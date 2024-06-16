package main

import (
	"encoding/json"
	"net/http"
)

func (apiCfg *apiConfig) handlerCreateUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string "name"
	}
	json.NewDecoder(r.body)
	respondWithJSON(w, 200, struct{}{})
}
