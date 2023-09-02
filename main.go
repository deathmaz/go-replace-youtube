package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for _, a := range os.Args[1:] {
		u, err := url.Parse(a)
		if err != nil {
			log.Fatal(err)
		}

		// https://youtu.be/xxxxxx
		isShortForm := strings.Contains(u.Host, "youtu.be")

		// https://www.youtube.com/watch?v=xxxxx
		isLongForm := strings.Contains(u.Host, "youtube")

		if !isShortForm && !isLongForm {
			fmt.Println("The given url can't be processed!")
			return
		}
		var videoId string
		queryVideoId := u.Query().Get("v")

		if isShortForm {
			videoId = strings.TrimPrefix(u.Path, "/")
		} else if queryVideoId != "" {
			videoId = queryVideoId
		}

		if len(videoId) != 0 {
			openUrl(videoId)
		}
	}
}

func openUrl(videoId string) {
	url := os.Getenv("RY_URL")
	if url == "" {
		log.Fatal("RY_URL variable is not set!")
	}
	out, err := exec.Command("xdg-open", url+videoId).Output()
	if err != nil {
		log.Fatal(err)
	} else {
		output := string(out[:])
		fmt.Println(output)
	}
}
