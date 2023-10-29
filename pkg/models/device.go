package models

import (
	"fmt"

	"github.com/akshatphumbhra/device-tracker/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type DeviceState struct {
	DriveStatus string `json:"drive_status"`
}

type LatestDevicePoint struct {
	Latitude    float64     `json:"lat"`
	Longitude   float64     `json:"lng"`
	DeviceState DeviceState `json:"device_state" gorm:"embedded"`
}

type Device struct {
	gorm.Model
	DeviceId          string `json:"device_id" gorm:"primary_key"`
	Name              string `json:"display_name"`
	ActiveStatus      string `json:"active_state"`
	IconUrl           string
	Visible           bool
	LatestDevicePoint LatestDevicePoint `json:"latest_device_point" gorm:"embedded"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Device{})
}

func (d *Device) CreateDevice() (*Device, error) {
	result := db.Create(&d)
	return d, result.Error
}

func GetAllDevices() ([]Device, error) {
	var devices []Device
	result := db.Find(&devices)
	return devices, result.Error
}

func (d *Device) UpdateExistingDevice(newDeviceDetails *Device) (*Device, error) {
	result := db.Model(&d).Updates(&newDeviceDetails)
	return newDeviceDetails, result.Error
}

func CreateOrUpdateDevices(device *Device) error {
	var existingDevice Device
	result := db.Where("device_id = ?", device.DeviceId).First(&existingDevice)
	if result.Error == gorm.ErrRecordNotFound {
		// Device doesn't exist in the database, create a new record
		device.Visible = true
		device.IconUrl = ""
		_, err := device.CreateDevice()
		if err != nil {
			return err
		}
		fmt.Printf("Device with DeviceId %s added to the database\n", device.DeviceId)
	} else if result.Error == nil {
		// Device exists, update its information
		_, err := existingDevice.UpdateExistingDevice(device)
		if err != nil {
			return err
		}
		fmt.Printf("Device with DeviceId %s updated in the database\n", device.DeviceId)
	} else {
		fmt.Printf("Error querying the database: %v\n", result.Error)
	}
	return result.Error
}

func UpdateDeviceIcon(deviceId string, iconUrl string) error {
	var existingDevice Device
	result := db.Where("device_id = ?", deviceId).First(&existingDevice)
	if result.Error == nil {
		existingDevice.IconUrl = iconUrl
		saveResult := db.Save(&existingDevice)
		if saveResult.Error != nil {
			return saveResult.Error
		}
		fmt.Printf("Device with DeviceId %s updated in the database\n", deviceId)
	} else {
		fmt.Printf("Error querying the database while updating device icon: %v\n", result.Error)
	}
	return result.Error
}

func UpdateDeviceVisibility(deviceId string, visible bool) error {
	var existingDevice Device
	result := db.Where("device_id = ?", deviceId).First(&existingDevice)
	if result.Error == nil {
		existingDevice.Visible = visible
		saveResult := db.Save(&existingDevice)
		if saveResult.Error != nil {
			return saveResult.Error
		}
		fmt.Printf("Device with DeviceId %s updated in the database\n", deviceId)
	} else {
		fmt.Printf("Error querying the database while updating visibility: %v\n", result.Error)
	}
	return result.Error
}
