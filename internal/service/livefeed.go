package service

import (
    "live-event-dashboard/internal/model"
    "live-event-dashboard/internal/store"
	"log"
	"github.com/gocolly/colly" 
)

const WEB_URL = "https://www.bet365.com/#/IP/B1"

type LiveFeedService struct {
    Store *store.Database
}

func NewLiveFeedService(db *store.Database) *LiveFeedService {
    return &LiveFeedService{Store: db}
}

func (lfs *LiveFeedService) GetLiveEvents() ([]model.Event, error) {
    return lfs.Store.GetLiveEvents()
}

func getLiveData() {
	c := colly.NewCollector()
	c.Visit(WEB_URL)
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting %s", r.URL)
	})

	c.OnHTML("div.ovm-CompetitionList ", func(e *colly.HTMLElement) {
		log.Print(e)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})


}
