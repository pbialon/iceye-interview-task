package satellite

import (
	"encoding/json"
	"time"
)

type Parser struct {
	timeFormat string
}

func (p *Parser) parse(data []byte) (SatelliteTLE, error) {
	raw := rawTLE{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return SatelliteTLE{}, err
	}

	date, err := time.Parse(p.timeFormat, raw.Date)
	if err != nil {
		return SatelliteTLE{}, err
	}

	return SatelliteTLE{
		Satellite: Satellite{
			ID:   raw.ID,
			Name: raw.ID,
		},
		Date:  date,
		Line1: raw.Line1,
		Line2: raw.Line2,
	}, nil
}

type rawTLE struct {
	ID    string `json:"id"`
	Date  string `json:"date"`
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
}
