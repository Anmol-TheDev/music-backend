package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

type VideoInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	WebpageURL  string `json:"webpage_url"`
}

func Ytdlp(videoURL string) ([]string, error) {
	// Prepare the yt-dlp command
	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "--get-url", videoURL)

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error executing yt-dlp: %v", err)
	}

	// Parse the output: yt-dlp -g may return multiple lines for video and audio URLs
	downloadURLs := strings.Split(strings.TrimSpace(out.String()), "\n")
	data, err := toBytes(downloadURLs[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	return downloadURLs, nil
}

func toBytes(url string) ([]byte, error) {
	if url == "" {
		panic("empty url ")
	}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var binary bytes.Buffer

	_, err = io.Copy(&binary, resp.Body)

	if err != nil {
		panic(err)
	}

	return binary.Bytes(), err
}
