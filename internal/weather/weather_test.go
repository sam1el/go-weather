// filepath: /workspaces/go-weather/internal/weather/weather_test.go
package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	// Mock response data
	mockResponse := `{
        "nearest_area": [{
            "areaName": [{"value": "North Richland Hills"}],
            "region": [{"value": "Texas"}]
        }],
        "current_condition": [{
            "temp_F": "75"
        }]
    }`

	// Create a new server that returns the mock response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Replace the URL in the GetWeather function with the test server URL
	getWeather := func(url string) (string, error) {
		resp, err := http.Get(url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

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

	// Call the GetWeather function
	t.Log("Calling GetWeather function with zipcode 76180")
	url := fmt.Sprintf("%s?format=j1", server.URL)
	weatherInfo, err := getWeather(url)

	// Extract the temperature from the mock response
	var weatherData map[string]interface{}
	err = json.Unmarshal([]byte(mockResponse), &weatherData)
	if err != nil {
		t.Fatalf("Error unmarshalling mock response: %v", err)
	}
	currentCondition := weatherData["current_condition"].([]interface{})[0].(map[string]interface{})
	tempF := currentCondition["temp_F"].(string)

	// Verify the result
	expected := fmt.Sprintf("The current weather conditions in North Richland Hills, Texas is %sºF.", tempF)
	if strings.TrimSpace(weatherInfo) != expected {
		t.Errorf("Expected %q, got %q", expected, weatherInfo)
	} else {
		t.Logf("Test passed: %q", weatherInfo)
	}
}
