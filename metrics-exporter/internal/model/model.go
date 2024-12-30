package model

import "time"

type Metric struct {
	// Date and Time
	Date time.Time `json:"date"`
	Time time.Time `json:"time"`

	// GPS Data
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	GPSSpeed    float32 `json:"gps_speed"`
	GPSAlt      float32 `json:"gps_alt"`
	GPSHeading  float32 `json:"gps_heading"`
	GPSDistance int32   `json:"gps_distance"`

	// Speed and Power
	Speed        float32 `json:"speed"`
	Voltage      float32 `json:"voltage"`
	PhaseCurrent float32 `json:"phase_current"`
	Current      float32 `json:"current"`
	Power        float32 `json:"power"`
	Torque       float32 `json:"torque"`
	PWM          float32 `json:"pwm"`

	// Battery and Distance
	BatteryLevel  int32 `json:"battery_level"`
	Distance      int32 `json:"distance"`
	TotalDistance int32 `json:"totaldistance"`

	// Temperature
	SystemTemp int32 `json:"system_temp"`
	Temp2      int32 `json:"temp2"`

	// Orientation
	Tilt float32 `json:"tilt"`
	Roll float32 `json:"roll"`

	// Mode and Alerts
	Mode  int32 `json:"mode"`
	Alert int32 `json:"alert"`
}
