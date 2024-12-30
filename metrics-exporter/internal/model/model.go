package model

type Metric struct {
	// Date and Time
	Date string `csv:"date"`
	Time string `csv:"time"`

	// GPS Data
	Latitude    float64 `csv:"latitude"`
	Longitude   float64 `csv:"longitude"`
	GPSSpeed    float32 `csv:"gps_speed"`
	GPSAlt      float32 `csv:"gps_alt"`
	GPSHeading  float32 `csv:"gps_heading"`
	GPSDistance int32   `csv:"gps_distance"`

	// Speed and Power
	Speed        float32 `csv:"speed"`
	Voltage      float32 `csv:"voltage"`
	PhaseCurrent float32 `csv:"phase_current"`
	Current      float32 `csv:"current"`
	Power        float32 `csv:"power"`
	Torque       float32 `csv:"torque"`
	PWM          float32 `csv:"pwm"`

	// Battery and Distance
	BatteryLevel  int32 `csv:"battery_level"`
	Distance      int32 `csv:"distance"`
	TotalDistance int32 `csv:"totaldistance"`

	// Temperature
	SystemTemp int32 `csv:"system_temp"`
	Temp2      int32 `csv:"temp2"`

	// Orientation
	Tilt float32 `csv:"tilt"`
	Roll float32 `csv:"roll"`

	// Mode and Alerts
	Mode  int32 `csv:"mode"`
	Alert int32 `csv:"alert"`
}
