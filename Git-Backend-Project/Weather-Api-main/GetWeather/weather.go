package GetWeather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	custommodels "weather-api/CustomModels"
)

const weatherApiUrl = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline"

// --------------------------------------------------------------------------+
func GetWeather(location string) (*custommodels.WeatherResponse, error) {
	ApiKey := os.Getenv("WEATHER_API_KEY")
	ApiUrl := fmt.Sprintf("%s/%s?key=%s", weatherApiUrl, location, ApiKey)
	log.Printf("Fetch weather for %s/n", location)

	ApiUrlBuilder, err := url.Parse(ApiUrl)
	if err != nil {
		return nil, fmt.Errorf("error in #block [2]| ApiBuilder\n%v", err)
	}

	//-------------------------------------------+-------------+
	parameter := url.Values{}                   //| add query
	parameter.Add("key", ApiKey)                //| parameters
	parameter.Add("contentType", "json")        //|
	parameter.Add("unitGroup", "metric")        //|
	ApiUrlBuilder.RawQuery = parameter.Encode() //|
	//--------------------------------------------+-------------+

	Response, err := http.Get(ApiUrlBuilder.String())
	if err != nil {
		return nil, fmt.Errorf("error in #block [3] | Create get req\n%v", err)
	}

	defer Response.Body.Close()

	if Response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in block [4] | unexpected status code\n%v", err)
	}

	var ApiResponse custommodels.WeatherResponse
	if err := json.NewDecoder(Response.Body).Decode(&ApiResponse); err != nil {
		return nil, fmt.Errorf("error in #block [5] | decoding json\n%v", err)
	}

	log.Printf("Code: '200'\n fetched weather")
	return &ApiResponse, nil
}

//---------------------------------------------------------------------------+
