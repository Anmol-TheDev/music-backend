package utils

import (
	"encoding/json"
	"fetch-spotify/modles"
	"fmt"
	"io"
	"net/http"
	"strings"
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
			Artists struct {
				Primary []struct {
					Name string `json:"name"`
				}
			} `json:"artists"`
		} `json:"results"`
	} `json:"data"`
}

func GetTrackfromJio(tracks *[]modles.TrackStr) {
	// var resTracks []modles.TrackStr
	var wg sync.WaitGroup
	wg.Add(len(*tracks))

	for i := range *tracks {
		name := (*tracks)[i].Name
		fullUrl := URL + name + "&limit=1"
		go func(url string) {
			artistName := (*tracks)[i].Artist[0]
			defer wg.Done() // Decrement counter when goroutine completes
			temp := makeRequest(url)
			if strings.EqualFold(artistName, temp.artist) {
				(*tracks)[i].Id = temp.savanID
				(*tracks)[i].DownloadUrl = temp.downloadUrl
			}
		}(fullUrl) // Pass the current url as an argument to avoid closure issues
	}
	defer wg.Wait()
}

type makeRequestStruct struct {
	// orgID string
	savanID     string
	downloadUrl string
	artist      string
}

func makeRequest(url string) makeRequestStruct {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error while jio savan ", err)
	}
	if res.StatusCode != http.StatusOK {
		return makeRequestStruct{}
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	var resBody JioSavanData
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		panic(err)
	}
	var temp makeRequestStruct
	temp.downloadUrl = resBody.Data.Results[0].DownloadUrl[4].Url
	temp.savanID = resBody.Data.Results[0].Id

	if len(resBody.Data.Results[0].Artists.Primary) > 0 {
		temp.artist = resBody.Data.Results[0].Artists.Primary[0].Name
	}
	return temp
}
