package go_holiday

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Holiday(date string) (string, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/guangrei/APIHariLibur_V2/main/holidays.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data map[string]map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if summary, ok := data[date]["summary"]; ok {
		return summary, nil
	}

	return "", nil
}
