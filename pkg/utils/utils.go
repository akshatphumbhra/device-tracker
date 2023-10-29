package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akshatphumbhra/device-tracker/pkg/models"
)

type APIResponse struct {
	ResultList []models.Device `json:"result_list"`
}

type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

func SyncDataFromApi() error {
	apiKey := os.Getenv("ONE_STEP_GPS_API_KEY")
	url := fmt.Sprintf("https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=%s", apiKey)

	response, err := http.Get(url)
	if err != nil {
		return CustomError{Message: "Error fetching data from API"}
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return CustomError{Message: "Error reading the response body"}
	}

	var apiResponse APIResponse

	if err := json.Unmarshal(responseBody, &apiResponse); err != nil {
		return CustomError{Message: "Error decoding JSON"}
	}

	for _, device := range apiResponse.ResultList {
		models.CreateOrUpdateDevices(&device)
	}
	return nil
}

func ValidateIconUrls(devices []models.Device) ([]models.Device, error) {
	for _, device := range devices {
		if device.IconUrl != "" {
			imageFilePath := filepath.Join("../../frontend/src/assets/", device.IconUrl)
			if !isFileExists(imageFilePath) {
				err := models.UpdateDeviceIcon(device.DeviceId, "")
				if err != nil {
					return nil, CustomError{Message: "Error updating device icon"}
				}
				device.IconUrl = ""
			}
		}
	}
	return devices, nil
}

func isFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
