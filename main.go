package main

import (
	"errors"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	for _, a := range os.Args[1:] {
		u, err := url.Parse(a)
		if err != nil {
			log.Fatal(err)
		}

		videoId, err := getVideoId(u)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = openUrl(videoId)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getVideoId(u *url.URL) (string, error) {
	// https://youtu.be/xxxxxx
	isShortForm := strings.Contains(u.Host, "youtu.be")

	// https://www.youtube.com/watch?v=xxxxx
	isLongForm := strings.Contains(u.Host, "youtube")

	if !isShortForm && !isLongForm {
		return "", errors.New("The given url can't be processed!")
	}
	var videoId string
	queryVideoId := u.Query().Get("v")

	if isShortForm {
		videoId = strings.TrimPrefix(u.Path, "/")
	} else if queryVideoId != "" {
		videoId = queryVideoId
	}

	if videoId == "" {
		return "", errors.New("Video id can't be parsed")
	}

	return videoId, nil
}

func openUrl(videoId string) error {
	url := os.Getenv("RY_URL")
	if url == "" {
		return errors.New("RY_URL variable is not set!")
	}
	return open.Run(url + videoId)
}
