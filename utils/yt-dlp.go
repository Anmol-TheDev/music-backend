package utils

import (
	"bytes"
	"os/exec"
)

type VideoInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	WebpageURL  string `json:"webpage_url"`
}

func Ytdlp(videoURL string) (string, error) {

	err := ytdlpCheck()
	if err != nil {
		return "", err
	}
	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "--get-url", videoURL)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	// Parse the output: yt-dlp -g may return multiple lines for video and audio URLs
	downloadURL := out.String()

	return downloadURL, nil
}

func ytdlpCheck() error {
	err := exec.Command("yt-dlp", "--version").Run()
	return err
}
