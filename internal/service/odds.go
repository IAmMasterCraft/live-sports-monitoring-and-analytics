package service

import (
    "live-event-dashboard/internal/model"
    "math/rand"
)

type OddsService struct {}

func NewOddsService() *OddsService {
    return &OddsService{}
}

func (os *OddsService) CalculateOdds(event model.Event) float64 {
    return rand.Float64() * 10
}
