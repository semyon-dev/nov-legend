package util

import (
	"bytes"
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"nov-legend/app/config"
)

func Translate(toLang, text string) string {

	if toLang == "" {
		toLang = "en"
	}

	URL := "https://translate.api.cloud.yandex.net/translate/v2/translate"

	message := map[string]interface{}{
		"format":             "PLAIN_TEXT",
		"sourceLanguageCode": "ru",
		"folderId":           config.FolderId,
		"targetLanguageCode": toLang,
		"texts":              []string{text},
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", "Bearer "+config.IAMTOKEN)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)

	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)

	}
	result := gjson.Get(string(body), "translations.0.text").String()

	return result
}
