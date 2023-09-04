package main

import (
	"net/url"
	"testing"
)

func TestGetVideoIdFromShortUrl(t *testing.T) {
	videoId := "xxxxx"
	videoUrl := "https://youtu.be/" + videoId

	u, _ := url.Parse(videoUrl)
	parsedVideoId, err := getVideoId(u)

	if parsedVideoId == "" || parsedVideoId != videoId || err != nil {
		t.Fatalf(`Parsed video id doesn't match %s`, videoId)
	}
}

func TestEmptyGetVideoIdFromShortUrl(t *testing.T) {
	videoUrl := "https://youtu.be/"

	u, _ := url.Parse(videoUrl)
	parsedVideoId, err := getVideoId(u)

	if parsedVideoId != "" || err == nil {
		t.Fatalf(`Parsed video id or error should not be empty when parsing %s`, videoUrl)
	}
}

func TestGetVideoIdFromLongUrl(t *testing.T) {
	videoId := "xxxxx"
	videoUrl := "https://www.youtube.com/watch?v=" + videoId

	u, _ := url.Parse(videoUrl)
	parsedVideoId, err := getVideoId(u)

	if parsedVideoId == "" || parsedVideoId != videoId || err != nil {
		t.Fatalf(`Parsed video id doesn't match %s`, videoId)
	}
}

func TestEmptyGetVideoIdFromLongUrl(t *testing.T) {
	videoUrl := "https://www.youtube.com/watch?v="

	u, _ := url.Parse(videoUrl)
	parsedVideoId, err := getVideoId(u)

	if parsedVideoId != "" || err == nil {
		t.Fatalf(`Parsed video id or error should not be empty when parsing %s`, videoUrl)
	}
}
