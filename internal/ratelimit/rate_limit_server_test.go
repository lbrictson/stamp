package ratelimit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/lbrictson/stamp/internal/logging"
)

func setupTestServer() {
	// Start server
	logging.Init()
	go RunServer()
}
func TestRunServer(t *testing.T) {
	// Run sever
	setupTestServer()
	// Setup request
	url := "http://localhost:8668/heartbeat"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "text/xml")
	res, _ := http.DefaultClient.Do(req)
	// Verify status code
	if res.StatusCode != 200 {
		t.Errorf("%v should have returned %v but returned %v", url, 200, res.StatusCode)
	}
	return
}

func TestRateLimitOK(t *testing.T) {
	// Setup request
	url := "http://localhost:8668/api/v1/limitip"
	payload := IPPayload{
		IPAddr: "8.8.8.8",
		Domain: "fakedomain.net",
		Limit:  3,
	}
	bytePayload, _ := json.Marshal(&payload)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytePayload))
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	// Verify status code
	if res.StatusCode != 200 {
		t.Errorf("%v should have returned %v but returned %v", url, 200, res.StatusCode)
	}
	return
}

func TestRateLimitBadPayload(t *testing.T) {
	// Setup request
	url := "http://localhost:8668/api/v1/limitip"
	payload := IPPayload{
		IPAddr:  "8.8.8.8",
		Domain:  "fakedomain.net",
		Limit:   3,
		Minutes: 1,
	}
	bytePayload, _ := json.Marshal(&payload)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytePayload))
	req.Header.Add("Content-Type", "txt/xml")
	res, _ := http.DefaultClient.Do(req)
	// Verify status code
	if res.StatusCode != 500 {
		t.Errorf("%v should have returned %v but returned %v", url, 500, res.StatusCode)
	}
	return
}

func TestRateLimitBreached(t *testing.T) {
	// Setup request
	url := "http://localhost:8668/api/v1/limitip"
	payload := IPPayload{
		IPAddr:  "8.8.8.8",
		Domain:  "ratelimitme.net",
		Limit:   3,
		Minutes: 1,
	}
	bytePayload, _ := json.Marshal(&payload)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytePayload))
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	// Verify status code, should be okay
	if res.StatusCode != 200 {
		t.Errorf("%v should have returned %v but returned %v", url, 200, res.StatusCode)
	}
	req, _ = http.NewRequest("POST", url, bytes.NewReader(bytePayload))
	req.Header.Add("Content-Type", "application/json")
	res, _ = http.DefaultClient.Do(req)
	// Verify status code, should be okay
	if res.StatusCode != 200 {
		t.Errorf("%v should have returned %v but returned %v", url, 200, res.StatusCode)
	}
	req, _ = http.NewRequest("POST", url, bytes.NewReader(bytePayload))
	req.Header.Add("Content-Type", "application/json")
	res, _ = http.DefaultClient.Do(req)
	// Verify status code, should be okay
	if res.StatusCode != 200 {
		t.Errorf("%v should have returned %v but returned %v", url, 200, res.StatusCode)
	}
	req, _ = http.NewRequest("POST", url, bytes.NewReader(bytePayload))
	req.Header.Add("Content-Type", "application/json")
	res, _ = http.DefaultClient.Do(req)
	// Verify status code, should get rate limited because it is request 3 of 4
	if res.StatusCode != 492 {
		t.Errorf("%v should have returned %v but returned %v", url, 492, res.StatusCode)
	}
	return
}
