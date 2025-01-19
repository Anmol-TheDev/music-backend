package utils

import (
	"encoding/json"
	"fetch-spotify/modles"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var URL = "https://saavn.dev/api/search/songs?query="

type JioSavanData struct {
	Data struct {
		Results []struct {
			Name        string `json:"name"`
			Id          string `json:"id"`
			DownloadUrl []struct {
				Quality string `json:"quality"`
				Url     string `json:"url"`
			} `json:"downloadUrl"`
			Arists struct {
				Primary []struct {
					Name string `json:"name"`
				}
			}
		} `json:"results"`
	} `json:"data"`
}

func GetTrackfromJio(tracks []modles.TrackStr) {
	var resTracks []modles.TrackStr
	var wg sync.WaitGroup
	wg.Add(len(tracks))

	for _, track := range tracks {
		name := track.Name
		fullUrl := URL + name + "&limit=1"
		go func(url string) {
			defer wg.Done() // Decrement counter when goroutine completes
			makeRequest(url, &resTracks)
		}(fullUrl) // Pass the current url as an argument to avoid closure issues
	}

}

func makeRequest(url string, restracks *[]modles.TrackStr) {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error while jio savan ", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var resBody JioSavanData
	err = json.Unmarshal(body, &resBody)
	var temp modles.TrackStr
	temp.DownloadUrl = resBody.Data.Results[0].DownloadUrl[4].Url
	if err != nil {
		fmt.Println(err)
	}
	*restracks = append(*restracks, )
}
