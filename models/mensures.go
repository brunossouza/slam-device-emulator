package models

import "encoding/json"

// Mensure struct type
type Mensure struct {
	Voltage float64 `json:"voltage"`
	Current float64 `json:"current"`
	Power   float64 `json:"power"`
	Energy  float64 `json:"energy"`
	Angle   float64 `json:"angle"`
}

// UnmarshalMensure json to object
func UnmarshalMensure(data []byte) (Mensure, error) {
	var r Mensure
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal object to json
func (r *Mensure) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
