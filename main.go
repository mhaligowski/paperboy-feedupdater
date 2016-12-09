package feedupdater

import (
	"net/http"
	"net/url"

	"google.golang.org/appengine"
	"google.golang.org/appengine/taskqueue"
)

type Feed struct {
	FeedId  string
	FeedUrl string
}

func init() {
	feeds := []Feed{
		{"dummy1", "http://rss.slashdot.org/Slashdot/slashdotMainatom"},
	}

	http.HandleFunc("/updates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx := appengine.NewContext(r)
		for _, feed := range feeds {
			t := taskqueue.NewPOSTTask("/handle", url.Values{
				"feed_id":  {feed.FeedId},
				"feed_url": {feed.FeedUrl},
			})

			if _, err := taskqueue.Add(ctx, t, "CrawlerQueue"); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	})

}
