package tle_api_client

import (
	"encoding/json"
	"fmt"
	"github.com/delabania/iceye-interview-task/src/satellite"
	"io"
	"log"
	"net/http"
	"time"
)

const TimeFormat = "2023-07-14T17:11:48+00:00"

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type Client struct {
	Endpoint   string
	httpClient HttpClient
}

func NewClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

func (c *Client) GetSatelliteTLE(s satellite.Satellite) (satellite.SatelliteTLE, error) {
	response, err := c.httpClient.Get(c.endpointForSatellite(s))
	if err != nil {
		return satellite.SatelliteTLE{}, err
	}

	if err != nil {
		return satellite.SatelliteTLE{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return satellite.SatelliteTLE{}, err
	}

	return parseTLEData(responseData)
}

func (c *Client) endpointForSatellite(satellite satellite.Satellite) string {
	return fmt.Sprintf("%s/%s", c.Endpoint, satellite.ID)
}

func parseTLEData(data []byte) (satellite.SatelliteTLE, error) {
	raw := rawTLE{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return satellite.SatelliteTLE{}, err
	}

	date, err := time.Parse(TimeFormat, raw.Date)
	if err != nil {
		return satellite.SatelliteTLE{}, err
	}

	return satellite.SatelliteTLE{
		Satellite: satellite.Satellite{
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
