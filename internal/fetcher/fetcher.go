package fetcher

import (
	"encoding/json"
	"errors"
	types "my-pubsub-app/utils"
	"net/http"
)

type RandomUserAPIResponse struct {
	Results []struct {
		Name struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Gender string `json:"gender"`
		Email  string `json:"email"`
	} `json:"results"`
}

func FetchUser() (types.User, error) {
	url := "https://randomuser.me/api/"
	resp, err := http.Get(url)
	if err != nil {
		return types.User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.User{}, errors.New("failed to fetch user data")
	}

	var apiResponse RandomUserAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return types.User{}, err
	}

	if len(apiResponse.Results) == 0 {
		return types.User{}, errors.New("no user data found")
	}

	result := apiResponse.Results[0]
	user := types.User{
		Name:   result.Name.Title + " " + result.Name.First + " " + result.Name.Last,
		Gender: result.Gender,
		Email:  result.Email,
	}

	return user, nil
}
