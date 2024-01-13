package yandex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/text/language"
)

type Yandex struct {
	apiKey   string
	folderId string
}

type translateRequest struct {
	FolderId           string   `json:"folderId"`
	Texts              []string `json:"texts"`
	SourceLanguageCode string   `json:"sourceLanguageCode"`
	TargetLanguageCode string   `json:"targetLanguageCode"`
}

type translation struct {
	Text                 string `json:"text"`
	DetectedLanguageCode string `json:"detectedLanguageCode"`
}
type translateResponse struct {
	Translations []translation `json:"translations"`
}

func New(apiKey, folderId string) *Yandex {
	return &Yandex{
		apiKey:   apiKey,
		folderId: folderId,
	}
}

func (y *Yandex) Translate(text string, sourceLang, targetLang language.Tag) (string, error) {
	body, err := json.Marshal(translateRequest{
		FolderId:           y.folderId,
		Texts:              []string{text},
		SourceLanguageCode: sourceLang.String(),
		TargetLanguageCode: targetLang.String(),
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, "https://translate.api.cloud.yandex.net/translate/v2/translate", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", strings.Join([]string{"Api-Key", y.apiKey}, " "))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("not status code 200 response, status code %v\n%v", resp.StatusCode, string(respBody))
	}

	var tr translateResponse
	err = json.Unmarshal(respBody, &tr)
	if err != nil {
		return "", err
	}

	return tr.Translations[0].Text, nil
}
