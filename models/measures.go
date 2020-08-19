package models

import "encoding/json"

// Measure struct type
type Measure struct {
	Voltage float64 `json:"voltage"`
	Current float64 `json:"current"`
	Power   float64 `json:"power"`
	Energy  float64 `json:"energy"`
	Angle   float64 `json:"angle"`
}

// UnmarshalMeasure json to object
func UnmarshalMeasure(data []byte) (Measure, error) {
	var r Measure
	err := json.Unmarshal(data, &r)
	return r, err
}

// MeasureMarshal object to json
func MeasureMarshal(r *Measure) ([]byte, error) {
	return json.Marshal(r)
}
