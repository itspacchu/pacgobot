package anilisthandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserData struct {
	Data struct {
		User struct {
			ID          int    `json:"id"`
			Avatar      Avatar `json:"avatar"`
			BannerImage string `json:"bannerImage"`
			SiteURL     string `json:"siteUrl"`
			Statistics  struct {
				Anime Anime `json:"anime"`
				Manga Manga `json:"manga"`
			} `json:"statistics"`
		} `json:"User"`
	} `json:"data"`
}

type Avatar struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

type Anime struct {
	Count     int     `json:"count"`
	MeanScore float64 `json:"meanScore"`
}

type Manga struct {
	Count     int     `json:"count"`
	MeanScore float64 `json:"meanScore"`
}

func sendRawGraphQLQuery(graphql string) UserData {
	url := "https://graphql.anilist.co"
	requestBody, err := json.Marshal(map[string]string{
		"query": graphql,
	})
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return UserData{}
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return UserData{}
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data UserData
		err := json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			fmt.Println("Error decoding JSON response:", err)
			return UserData{}
		}
		return data
	} else {
		fmt.Println("Error:", resp.Status)
		return UserData{}
	}
}

type AnimeData struct {
}

func GetAnimeData(animeName string) AnimeData {

	return AnimeData{}
}

func GetUserData(username string) UserData {
	query := fmt.Sprintf(`
    {
        User(name: "%s") {
            id
            avatar {
                large
                medium
            }
            bannerImage
            siteUrl
            statistics {
                anime {
                    count
                    meanScore
                }
                manga {
                    count
                    meanScore
                }
            }
        }
    }
    `, username)
	return sendRawGraphQLQuery(query)
}
