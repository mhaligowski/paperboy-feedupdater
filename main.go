package feedupdater

import (
	"net/http"
	"net/url"

	"google.golang.org/appengine"
	"google.golang.org/appengine/taskqueue"
	"github.com/mhaligowski/paperboy-feeds"
	"encoding/json"
)

func init() {
	feeds := []feeds.Feed{
		{"ff0b434e9374265b95316f6e5d09193eb4f81bbc760cc508308d6ee4da5af339",
			"Slashdot",
			"http://rss.slashdot.org/Slashdot/slashdotMainatom"},
	}

	http.HandleFunc("/updates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx := appengine.NewContext(r)

		for _, feed := range feeds {
			t := taskqueue.NewPOSTTask("/handle", url.Values{})
			v, err := json.Marshal(feed)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			t.Payload = v
			t.Header.Set("Content-Type", "application/json")


			if _, err := taskqueue.Add(ctx, t, "CrawlerQueue"); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	})

}
