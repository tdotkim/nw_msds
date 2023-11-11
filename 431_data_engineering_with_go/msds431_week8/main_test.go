package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_homeHandlerHTTP(t *testing.T) {
	if !assert.HTTPSuccess(t, homeHandler, "GET", "/search", nil) {
		// if the result is not correct print error
		t.Error("HTTP connection unsuccessful")
	}
}

func Test_searchHandlerHTTP(t *testing.T) {
	if !assert.HTTPSuccess(t, searchPage, "GET", "/search", nil) {
		// if the result is not correct print error
		t.Error("HTTP connection unsuccessful")
	}
}

func Test_scrapeHandlerHTTP(t *testing.T) {
	if !assert.HTTPSuccess(t, scrapeIt, "GET", "/dosearch", nil) {
		// if the result is not correct print error
		t.Error("HTTP connection unsuccessful")
	}
}

func Test_homeHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	//fmt.Println(response.Body.String())

	t.Run("test1", func(t *testing.T) {
		homeHandler(response, request)
	})
	//fmt.Println(response.Body.String())

	got := response.Body.String()
	want := `<a href="/search">Go to search page</a>`
	if got != want {

		// if the result is not correct print error
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test_searchPage(t *testing.T) {

	formStatic := `<label>URL:</label><br />`

	if !assert.HTTPBodyContains(t, searchPage, "GET", "/search", nil, formStatic) {
		// if the result is not correct print error
		t.Error("search page missing form")
	}

}

func Test_scrapeIt(t *testing.T) {
	form := url.Values{}
	form.Add("Url", "https://www.merriam-webster.com/dictionary/pickle	")
	form.Add("Word", "pickle")
	form.Add("Countt", "45")
	request, _ := http.NewRequest("post", "/dosearch", strings.NewReader(form.Encode()))
	response := httptest.NewRecorder()
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{"mwdict", args{response, request}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scrapeIt(tt.args.w, tt.args.req)
		})
	}
}
