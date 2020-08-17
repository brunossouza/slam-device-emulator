package controllers

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"slam-device-emulator/models"
	"time"
)

// GetDeviceConfig registra device
func GetDeviceConfig(equipamento string, host string) (device models.Device) {

	for { //https://slam-gui.herokuapp.com/
		response, err := http.Post("http://"+host+"/api/v1/devices/registry", "application/json", bytes.NewBufferString(equipamento))
		// response, err := http.Post("https://slam-gui.herokuapp.com/api/v1/devices/registry", "application/json", bytes.NewBufferString(equipamento))
		if err != nil {
			log.Println(err)
		}

		if response != nil && response.StatusCode == 200 {
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println(err)
			}
			device, err = models.UnmarshalDevice(body)
			if err != nil {
				log.Println(err)
			}
			break
		} else {
			time.Sleep(1 * time.Second)
		}
	}
	return device
}
