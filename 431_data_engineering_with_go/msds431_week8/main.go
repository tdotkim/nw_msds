package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"week8/scraper"
)

type Page struct {
	Title string
	Body  []byte
}

func searchPage(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./templates/basic.html")
	t.Execute(w, nil)

}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `<a href="/search">Go to search page</a>`)
}

func scrapeIt(w http.ResponseWriter, req *http.Request) {
	a := scraper.Scrape(req.FormValue("url"), req.FormValue("word"))
	varmap := map[string]interface{}{
		"Url":   string(a.Url),
		"Word":  string(a.Word),
		"Count": fmt.Sprint(a.Count),
	}
	t, _ := template.ParseFiles("./templates/response.html")
	t.Execute(w, varmap)
}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/search", searchPage)
	mux.HandleFunc("/dosearch", scrapeIt)
	mux.HandleFunc("/", homeHandler)

}

func main() {

	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	mux := http.NewServeMux()
	setupHandlers(mux)

	log.Fatal(http.ListenAndServe(listenAddr, mux))

}
