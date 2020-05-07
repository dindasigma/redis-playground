package dto

// data transfer object

import (
	"encoding/gob"
	"time"
)

type SensorMessage struct {
	Name string
	Value float64
	Timestamp time.Time
}

func init() {
	// gob encoding to encode the message
	gob.Register(SensorMessage{})
}