package datatype

import "time"

type DataType struct {
	Value     string        `json:"value"`
	Ttl       time.Duration `json:"ttl"`
	validTill time.Time
}

func New(value string, duration time.Duration) DataType {
	return DataType{value, duration, time.Now().Add(duration)}
}
