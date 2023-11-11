# Web Applications with Go

## Project Summary

For week 8 we were tasked with creating a simple web app with Go that could be used as the basis for an information agent. 

For this week we leveraged the Go-Colly package as a webscraper and used the standard http library for tthe web application itself.

THere is a simple three page structure with one page serving as a root page "/". Another page serving as the input page "/search" which takes in a URL and the word that is to be searched for. The final page is a simple results page that returns the URL, the word, and the total number of occurrences in the body of the page.

This lays the groundwork for additional decisions to be made by the client in regards to flagging a website as low/medium/high for contextual relevance. Marketing clients can determine if a blog is a good target for a display ad based on this hypothetical contextual decision engine. Internal teams could set context scores for their sites as part of a recommendation system for employees that might be looking around their confluence space for helpful docs.

## Findings

Go's webserver is highly performant even with using Go-Colly instead of our base http app created in Week 5. The app could be rewritten with the scraper module being fully replaced by the Week 5 module to get even better performance.

The project closely resembles the flask app for python and is structured to follow a similar format with templates being served through handlers and data shared between handlers via get/post. 

## Overall Recommendation

THe company should be able to swap over to go to serve as the backend for the webapp. The ability ot natively handle everything offers significant advantages over having to import an external library. And if external libraries are going to be used, a package like Wails offers similar funcitonality to Electron. 

For the actual advanced decisioning, a different language should be used unless there is a strong desire to recreate what most likely already exists for Python in Go. 

## Go Setup


*.:* \
Root folder contains the go program + testing.

*./main.go:* \
The main program, starts a local web app on port 8080.

*./main_test.go:* \
Testing script. Contains a few unit tests for the handlers to confirm routing works and outputs work. ALso leverages the testify package to help do unit tests for HTTP responses. 

*./templates:* \
This folder hosts two templates, basic.html and response.html. Basic serves as the basic template for the input page. Response serves as the template for the returned results.

*./scraper:* \
Contains the scraper built on Go-Colly. Main go file references this package locally. 


## Instructions

To recreate this please run the week8 executable and navigate to http://localhost:8080/

## References

https://github.com/gocolly/colly

https://github.com/stretchr/testify