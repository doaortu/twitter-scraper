package twitterscraper_test

import (
	"testing"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func TestParseUserQuery(t *testing.T) {
	queryMap := twitterscraper.ParseUserQuery("testtos")

	if queryMap["username"] != "testtos" {
		t.Errorf("Expected username to be testtos, got %s", queryMap["username"])
	}
	if queryMap["max_id"] != "" {
		t.Errorf("Expected max_id to be empty, got %s", queryMap["max_id"])
	}

	queryMap = twitterscraper.ParseUserQuery("testtos max_id:1321421")

	if queryMap["username"] != "testtos" {
		t.Errorf("Expected username to be testtos, got %s", queryMap["username"])
	}
	if queryMap["max_id"] != "1321421" {
		t.Errorf("Expected max_id to be 1321421, got %s", queryMap["max_id"])
	}
}
