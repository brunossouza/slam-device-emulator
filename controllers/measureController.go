package controllers

import (
	"bytes"
	"math"
	"math/rand"
	"net/http"
	"slam-device-emulator/models"
	"slam-device-emulator/utils"
	"time"
)

func generateFakeMeasureValue(min float64, max float64) float64 {
	return math.Min((rand.Float64()*min)+min, max)
}

// CreateMeasure cria dados fakes da leitura
func CreateMeasure() models.Measure {
	measure := models.Measure{}
	measure.Voltage = generateFakeMeasureValue(200, 240)
	measure.Current = generateFakeMeasureValue(7, 10)
	measure.Power = generateFakeMeasureValue(1400, 2400)
	measure.Energy = generateFakeMeasureValue(1, 2370)
	measure.Angle = generateFakeMeasureValue(0.8, 1)
	return measure
}

// RegistryMeasure registra device
func RegistryMeasure(token string, host string) {

	measure := CreateMeasure()

	dataToSend, err := models.MeasureMarshal(&measure)
	utils.CheckError(err)

	for { //https://slam-gui.herokuapp.com/

		response, err := http.Post("http://"+host+"/api/v1/"+token+"/measures", "application/json", bytes.NewBuffer(dataToSend))
		// response, err := http.Post("https://slam-gui.herokuapp.com/api/v1/devices/registry", "application/json", bytes.NewBufferString(equipamento))
		utils.CheckError(err)

		if response != nil && response.StatusCode == 200 {
			defer response.Body.Close()

			break
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
