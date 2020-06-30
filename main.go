package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Device map
type Device struct {
	ID                 int64       `json:"id"`
	RegistryDate       int64       `json:"registryDate"`
	Local              interface{} `json:"local"`
	Device             string      `json:"device"`
	Token              string      `json:"token"`
	Status             string      `json:"status"`
	DateLastAlter      interface{} `json:"dateLastAlter"`
	DateActivation     interface{} `json:"dateActivation"`
	Uptime             string      `json:"uptime"`
	DataCadastroString string      `json:"dataCadastroString"`
}

var (
	host        string
	equipamento string
	devices     []Device
)

// UnmarshalDevice json to object
func UnmarshalDevice(data []byte) (Device, error) {
	var r Device
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal object to json
func (r *Device) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func getDeviceConfig() (device Device) {

	for {
		response, err := http.Post("http://"+host+"/api/v1/devices/registry", "application/json", bytes.NewBufferString(equipamento))
		if err != nil {
			log.Println(err)
		}

		if response != nil && response.StatusCode == 200 {
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println(err)
			}
			device, err = UnmarshalDevice(body)
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

func init() {
	flag.StringVar(&host, "h", "localhost:8080", "ip e porta do servidor")
	flag.StringVar(&equipamento, "e", "não informado", "equipamento que o consumo será emulado, exemplo televisão")
}

func main() {
	devices = append(devices, getDeviceConfig())
	for _, dev := range devices {
		fmt.Println("Device:\t", dev)
	}
	fmt.Println("done")
}
