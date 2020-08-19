package models

import "encoding/json"

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

// UnmarshalDevice json to object
func UnmarshalDevice(data []byte) (Device, error) {
	var r Device
	err := json.Unmarshal(data, &r)
	return r, err
}

// DeviceMarshal object to json
func (r *Device) DeviceMarshal() ([]byte, error) {
	return json.Marshal(r)
}
