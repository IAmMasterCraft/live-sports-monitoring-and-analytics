package service

import (
	"context"
	"live-event-dashboard/internal/model"
	"live-event-dashboard/internal/store"
	"log"
	"strings"
	"github.com/chromedp/chromedp"
	"github.com/PuerkitoBio/goquery"
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

func GetLiveData() {
	log.Print("Getting live data")
	// Start Chrome Headless
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()
	var renderedHTML string
	err := chromedp.Run(ctx,
        chromedp.Navigate(WEB_URL),
        chromedp.OuterHTML("html", &renderedHTML, chromedp.ByQuery),
    )
    if err != nil {
        log.Fatalf("Failed to render page: %v", err)
    }

	log.Printf("Page rendered successfully: %s", renderedHTML)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(renderedHTML))
    if err != nil {
        log.Fatalf("Error loading HTTP response body: %v", err)
    }
	
    doc.Find("div.hm-MainHeaderWide").Each(func(i int, s *goquery.Selection) {
        log.Printf("Data found: %s", s.Text())
    })

	log.Print("Finished getting live data")

}
