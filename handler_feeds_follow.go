package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Eche5/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedToFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error parsing JSON")
		return
	}
	newFeeds, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating Feeds folllow:%v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedsFollowToFeedsFollow(newFeeds))
}

func (apiCfg *apiConfig) handlerGetFeedToFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	newFeeds, err := apiCfg.DB.ViewFeedsByUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating Feeds folllow:%v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedsToFollowArray(newFeeds))
}

func (apiCfg *apiConfig) handlerDeleteFeedToFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedfollowID")

	feedfollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("couldn't parse Feeds folllow:%v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedsToFollow(r.Context(), database.DeleteFeedsToFollowParams{
		ID:     feedfollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Deleting Feeds folllow:%v", err))
		return
	}
	respondWithJSON(w, 200, fmt.Sprintf("error:%v",err))
}
