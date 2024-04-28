package service

import (
    "live-event-dashboard/internal/model"
    "live-event-dashboard/internal/store"
)

type LiveFeedService struct {
    Store *store.Database
}

func NewLiveFeedService(db *store.Database) *LiveFeedService {
    return &LiveFeedService{Store: db}
}

func (lfs *LiveFeedService) GetLiveEvents() ([]model.Event, error) {
    return lfs.Store.GetLiveEvents()
}
