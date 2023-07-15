package satellite

import "time"

type Satellite struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SatelliteTLE struct {
	Satellite Satellite `json:"satellite"`
	Date      time.Time `json:"date"`
	Line1     string    `json:"line1"`
	Line2     string    `json:"line2"`
}
