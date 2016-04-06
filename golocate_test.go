package golocate

import (
	"testing"
	"fmt"
	"io/ioutil"
)

func TestWifiAccessPointsContruct(t *testing.T) {
	wap := WifiAccessPoints{
		MacAddress: "00:04:F2:80:81:BC",
		SignalStrength: -34,
		Age: 0,
		Channel: 11,
		SignalToNoiseRatio: -34,
	}

	if wap.MacAddress != "00:04:F2:80:81:BC" {
		t.Error("Expected 00:04:F2:80:81:BC, got ", wap.MacAddress)
	}

}

func TestGeoClientUrl(t *testing.T) {
	gc := NewGeoClient("TestApiString")
	if gc.getRequestUrl() != "https://www.googleapis.com/geolocation/v1/geolocate?key=TestApiString" {
		t.Error("Expected https://www.googleapis.com/geolocation/v1/geolocate?key=TestApiString, got ", gc.getRequestUrl())
	}
}

func TestSimpleRequest(t *testing.T) {
	gc := NewGeoClient("TestKey")
	wap := []WifiAccessPoints{WifiAccessPoints{
		MacAddress: "01:23:45:67:89:AB",
		SignalStrength: 8,
		Age: 0,
		Channel: 8,
		SignalToNoiseRatio: -65,
	}}
	req := Request{
		HomeMobileCountryCode: 310,
		HomeMobileNetworkCode: 260,
		RadioType: "gsm",
		Carrier: "T-Mobile",
		WifiAccessPoints: wap,
	}

	resp := gc.SendRequest(req)

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(contents)
	}
}