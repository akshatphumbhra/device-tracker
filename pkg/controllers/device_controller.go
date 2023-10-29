package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akshatphumbhra/device-tracker/pkg/config"
	"github.com/akshatphumbhra/device-tracker/pkg/models"
	"github.com/akshatphumbhra/device-tracker/pkg/utils"
	"github.com/google/uuid"
)

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&models.Device{})
}

func FetchDeviceData(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	err := utils.SyncDataFromApi(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	queryParams := r.URL.Query()
	shouldValidateIconUrls := queryParams.Get("validateIconUrls")
	if shouldValidateIconUrls == "true" {
		err = utils.ValidateIconUrls(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	devices, err := models.GetAllDevices(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(devices)

	// Set CORS headers to allow requests from any origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateDeviceVisibility(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	var visibilityUpdates []struct {
		DeviceId string `json:"deviceId"`
		Visible  bool   `json:"visible"`
	}

	err := json.NewDecoder(r.Body).Decode(&visibilityUpdates)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	for _, visibilityUpdate := range visibilityUpdates {
		err := models.UpdateDeviceVisibility(db, visibilityUpdate.DeviceId, visibilityUpdate.Visible)
		if err != nil {
			http.Error(w, "Failed to update device visibility", http.StatusInternalServerError)
			return
		}
	}

	// Set CORS headers to allow requests from any origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
}

func UpdateDeviceIcon(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	err := r.ParseMultipartForm(10 * 1024 * 1024) // 10MB limit
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	// Loop through all uploaded files -- stored as key value pairs
	// where the key is the deviceID and the value is the image file
	for deviceID, headers := range r.MultipartForm.File {
		for _, header := range headers {
			file, err := header.Open()
			if err != nil {
				http.Error(w, "Failed to open uploaded file", http.StatusBadRequest)
				return
			}
			defer file.Close()

			// create a unique name for the file
			// save the file to the server, in the frontend/src/assets/icons folder
			uniqueID := uuid.New().String()
			fileExt := filepath.Ext(header.Filename)
			filePath := filepath.Join("../../frontend/src/assets/icons", uniqueID+fileExt)
			imgFile, err := os.Create(filePath)
			if err != nil {
				http.Error(w, "Failed to create image file: "+err.Error(), http.StatusInternalServerError)
				return
			}
			defer imgFile.Close()

			_, err = io.Copy(imgFile, file)
			if err != nil {
				http.Error(w, "Failed to save image", http.StatusInternalServerError)
				return
			}

			// set the device.IconUrl as relative path to the image file for frontend to use
			imageUrl := "icons/" + uniqueID + fileExt
			updateError := models.UpdateDeviceIcon(db, deviceID, imageUrl)
			if updateError != nil {
				http.Error(w, "Failed to update device icon", http.StatusInternalServerError)
				return
			}
		}
	}

	// Set CORS headers to allow requests from any origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
