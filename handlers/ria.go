package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/shinoxzu/news-provider-api/config"
	"github.com/shinoxzu/news-provider-api/models"

	"github.com/mmcdole/gofeed"
)

type RiaHandler struct {
	Config *config.Config
}

func (h RiaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(h.Config.Rss.Ria)

	var news []models.News

	for _, item := range feed.Items {
		news = append(news, models.News{Title: item.Title})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(news)
}
