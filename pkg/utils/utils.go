package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/akshatphumbhra/device-tracker/pkg/models"
)

type APIResponse struct {
	ResultList []models.Device `json:"result_list"`
}

func SyncDataFromApi() {
	apiKey := os.Getenv("ONE_STEP_GPS_API_KEY")
	url := fmt.Sprintf("https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=%s", apiKey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body", err)
		return
	}

	var apiResponse APIResponse

	if err := json.Unmarshal(responseBody, &apiResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, device := range apiResponse.ResultList {
		models.CreateOrUpdateDevices(&device)
	}
}
