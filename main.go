package main

import (
    "google.golang.org/appengine/taskqueue"
    "golang.org/x/net/context"
    "net/http"
    "net/url"
)

type Feed struct {
    FeedId string
    FeedUrl string
}

func init() {
    feeds := []Feed{
        Feed{"dummy1", "http://rss.slashdot.org/Slashdot/slashdotMainatom"},
    }

    http.HandleFunc("/updates", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPut {
            http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
            return
        }

        for _, feed := range feeds {
            t := taskqueue.NewPOSTTask("/handle", url.Values{
                "feed_id": {feed.FeedId},
                "feed_url": {feed.FeedUrl},
            })

            taskqueue.Add(context.Background(), t, "")
        }
    })

}

