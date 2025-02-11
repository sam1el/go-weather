package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func GetWeather(zipcode string) (string, error) {
	url := fmt.Sprintf("http://wttr.in/%s?format=j1", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the response to extract city, state, and temperature
	var weatherData map[string]interface{}
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return "", err
	}

	nearestArea := weatherData["nearest_area"].([]interface{})[0].(map[string]interface{})
	city := nearestArea["areaName"].([]interface{})[0].(map[string]interface{})["value"].(string)
	state := nearestArea["region"].([]interface{})[0].(map[string]interface{})["value"].(string)
	currentCondition := weatherData["current_condition"].([]interface{})[0].(map[string]interface{})
	tempF := currentCondition["temp_F"].(string)

	return fmt.Sprintf("The current weather conditions in %s, %s is %sºF.", city, state, tempF), nil
}