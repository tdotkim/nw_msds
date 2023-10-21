package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type LoggingClient struct {
	log *log.Logger
}

func (c LoggingClient) RoundTrip(
	r *http.Request,
) (*http.Response, error) {
	c.log.Printf(
		"Sending a %s request to %s over %s\n",
		r.Method, r.URL, r.Proto,
	)
	resp, err := http.DefaultTransport.RoundTrip(r)
	c.log.Printf("Got back a response over %s\n", resp.Proto)

	return resp, err
}

func createHTTPClientWithTimeout(d time.Duration) *http.Client {
	client := http.Client{Timeout: d}
	return &client
}

func fetchRemoteResource(
	client *http.Client, url string,
) ([]byte, error) {
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}

type ScrapedFields struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func getTitle(body []byte) string {
	startkeys := "<title>"
	endkeys := "</title>"
	stringified := string(body)
	startindex := strings.Index(stringified, startkeys) + len(startkeys)
	endindex := strings.Index(stringified, endkeys)
	title := stringified[startindex:endindex]
	return string(title)
}

func main() {
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	//check to makesure we got two
	/* if len(os.Args) < 2 {
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to get data from")
		os.Exit(1)
	} */

	//make wikipages folder
	mkerr := os.Mkdir("./wikipages", 0777)
	if mkerr != nil {
		fmt.Println(mkerr)
	}

	//make items.jl file
	f, err := os.Create("items.jl")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	myTransport := LoggingClient{}
	l := log.New(os.Stdout, "", log.LstdFlags)
	myTransport.log = l

	client := createHTTPClientWithTimeout(15 * time.Second)
	client.Transport = &myTransport

	for _, link := range urls {
		filenameholder := strings.Split(link, "/")
		filename := filenameholder[len(filenameholder)-1]
		filepath := fmt.Sprintf("./wikipages/%v.html", filename)

		body, err := fetchRemoteResource(client, link)
		if err != nil {
			fmt.Fprintf(os.Stdout, "%#v\n", err)
			os.Exit(1)
		}

		var structHolder ScrapedFields
		structHolder.Url = string(link)

		structHolder.Title = getTitle(body)

		stringholdA := string(body)
		stringholdB := strings.ReplaceAll(stringholdA, "\n", "")
		stringhold := strings.ReplaceAll(stringholdB, "\t", "")
		re := regexp.MustCompile("<[^>]*>")
		newstr := re.ReplaceAllString(stringhold, "")
		re2 := regexp.MustCompile(`<script\b[^<]*(?:(^<\/script>)<[^<]*)*<\/script>`)
		newstr2 := re2.ReplaceAllString(newstr, "")
		structHolder.Text = newstr2

		//fmt.Println(newstr2)

		jsonstruct, errorthree := json.Marshal(structHolder)
		if errorthree != nil {
			fmt.Println(errorthree)
		}

		fmt.Fprintln(f, string(jsonstruct))

		errtwo := os.WriteFile(filepath, body, 0777)
		if errtwo != nil {
			fmt.Fprintf(os.Stdout, "%#v\n", errtwo)
			os.Exit(1)
		}
	}
}
