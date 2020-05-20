package website

import (
	"fmt"
	"log"
	"monitor/request"
	"net/http"
)

func StartSite(ch chan request.Crawl) {
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "website/index.html")
	})
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Something failed parsing your form: %v", err)
			return
		}
		email := r.Form.Get("email")
		url := r.Form.Get("url")
		if email != "" && url != "" {
				ch <- request.Crawl{
					Email: email,
					Url:   url,
				}
		}
		fmt.Fprint(w, "Got it. Will email you when the site goes down, then again when it goes up.")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
