package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_getTitle(t *testing.T) {
	type args struct {
		body []byte
	}

	chatbot, err := os.ReadFile("./WebFocusedCrawlWorkV001/wikipages/Chatbot.html")
	if err != nil {
		fmt.Print(err)
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"chatbot", args{chatbot}, "Chatbot - Wikipedia"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTitle(tt.args.body); got != tt.want {
				t.Errorf("getTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
