package tle_api_client

import (
	"fmt"
	"github.com/delabania/iceye-interview-task/src/satellite"
	"io"
	"log"
	"net/http"
)

const TimeFormat = "2023-07-14T17:11:48+00:00"

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type DataParser interface {
	Parse(data []byte) (satellite.SatelliteTLE, error)
}

type Client struct {
	endpoint   string
	httpClient HttpClient
	parser     DataParser
}

func NewClient(endpoint string, httpClient HttpClient, parser DataParser) *Client {
	return &Client{
		endpoint:   endpoint,
		httpClient: httpClient,
		parser:     parser,
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

	return c.parser.Parse(responseData)
}

func (c *Client) endpointForSatellite(satellite satellite.Satellite) string {
	return fmt.Sprintf("%s/%s", c.endpoint, satellite.ID)
}
