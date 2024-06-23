package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Eche5/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error parsing JSON")
		return
	}
	newFeeds, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating Feeds:%v", err))
	}
	respondWithJSON(w, 200, databaseFeedsToFeeds(newFeeds))
}

func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {

	AllFeeds, err := apiCfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error fetching feeds:%v", err))
		return
	}
	var allData [2]interface{}
	allData[1] = databaseFeedsToFeedsArray(AllFeeds)
	allData[0] = fmt.Sprintf("there are %v feeds",len(databaseFeedsToFeedsArray(AllFeeds)))

	respondWithJSON(w, 200, allData)
}
