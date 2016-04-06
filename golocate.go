package golocate

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
)

type WifiAccessPoints struct {
	MacAddress string `json:"macAddress"`
	SignalStrength int `json:"signalStrength"`
	Age int `json:"age"`
	Channel int `json:"channel"`
	SignalToNoiseRatio int `json:"signalToNoiseRatio"`
}

type CellTowers struct {
	CellId int `json:"cellId"`
	LocationAreaCode int `json:"locationAreaCode"`
	MobileCountryCode int `json:"mobileCountryCode"`
	MobileNetworkCode int `json:"mobileNetworkCode"`
}

type Request struct {
	HomeMobileCountryCode int `json:"homeMobileCountryCode"`
	HomeMobileNetworkCode int `json:"homeMobileNetworkCode"`
	RadioType string `json:"radioType"`
	Carrier string `json:"carrier"`
	WifiAccessPoints []WifiAccessPoints `json:"wifiAccessPoints"`
	CellTowers []CellTowers `json:"cellTowers"`
}

type GeoClient struct {
	Url string
	ApiKey string
}

func NewGeoClient(key string) *GeoClient {
	gc := GeoClient{
		Url: "https://www.googleapis.com/geolocation/v1/geolocate?key=",
		ApiKey: key,
	}
	return &gc
}

func (gc *GeoClient) getRequestUrl() string {
	url := (gc.Url + gc.ApiKey)
	return url
}

func (gc *GeoClient) SendRequest(reqest Request) *http.Response {
	jsonStr, _ := json.Marshal(reqest)
	req, err := http.NewRequest("POST", gc.getRequestUrl(), bytes.NewBuffer(jsonStr))
	fmt.Print(req)
	if err != nil {
		fmt.Println(req)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp
}