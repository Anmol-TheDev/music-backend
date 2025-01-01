package utils

import (
	"bytes"
	"fmt"
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
	return downloadURLs, nil
}
