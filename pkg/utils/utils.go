package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akshatphumbhra/device-tracker/pkg/models"
	"github.com/jinzhu/gorm"
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

func SyncDataFromApi(db *gorm.DB) error {
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
		models.CreateOrUpdateDevices(db, &device)
	}
	return nil
}

func ValidateIconUrls(db *gorm.DB) error {
	devices, err := models.GetAllDevices(db)
	if err != nil {
		return CustomError{Message: "Error fetching devices from database"}
	}

	for _, device := range devices {
		if device.IconUrl != "" {
			imageFilePath := filepath.Join("../../frontend/src/assets/", device.IconUrl)
			if !isFileExists(imageFilePath) {
				err := models.UpdateDeviceIcon(db, device.DeviceId, "")
				if err != nil {
					return CustomError{Message: "Error updating device icon"}
				}
			}
		}
	}
	return nil
}

func isFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
