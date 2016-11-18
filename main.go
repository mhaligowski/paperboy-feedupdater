package main

import (
    "net/http"
    "encoding/json"
    "fmt"
)

type Feed struct {
    FeedId string `json:"feed_id"`
    FeedUrl string `json:"feed_url"`
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

        output, err := json.Marshal(feeds)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "%s", output)
    })

}

