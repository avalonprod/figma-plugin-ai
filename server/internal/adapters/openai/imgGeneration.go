package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type imgGenerator struct {
	ApiKey string
	ApiUrl string
}

func NewImgGeneration(apiKey string, apiUrl string) *imgGenerator {
	return &imgGenerator{
		ApiKey: apiKey,
		ApiUrl: apiUrl,
	}

}

type imgGenerationInput struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type imgItem struct {
	Url string `json:"url"`
}

type imgGenerationResponse struct {
	Created int `json:"created"`
	Data    []imgItem
}

func (g *imgGenerator) NewImg(prompt string) ([]imgItem, error) {
	client := &http.Client{}

	input := imgGenerationInput{
		Prompt: prompt,
		N:      4,
		Size:   "512x512",
	}
	inputBytes, err := json.Marshal(input)
	req, err := http.NewRequest("POST", g.ApiUrl, bytes.NewBuffer(inputBytes))

	if err != nil {
		return []imgItem{}, errors.New("error creating request.")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", g.ApiKey))
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)

	if err != nil {
		return []imgItem{}, err
		fmt.Print(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return []imgItem{}, err
	}

	var output imgGenerationResponse
	json.Unmarshal(body, &output)

	content := output

	return content.Data, nil
}
