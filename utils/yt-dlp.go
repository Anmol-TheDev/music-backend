package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

type VideoInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	WebpageURL  string `json:"webpage_url"`
}

func Ytdlp(videoURL string) (string, error) {
	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "--get-url", videoURL)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	// Parse the output: yt-dlp -g may return multiple lines for video and audio URLs
	downloadURL := out.String()
	return downloadURL, nil
}
