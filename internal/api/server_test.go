package api

import (
	"github.com/lbrictson/stamp/internal/logging"
	"github.com/lbrictson/stamp/internal/rulecache"
	"net/http"
	"testing"
)

func setupTestServer() {
	// Start server
	logging.Init()
	rulecache.Init()
	go RunServer()
}

func TestAPIServerHeartbeat(t *testing.T) {
	// Start server
	setupTestServer()
	// Setup request
	url := "http://localhost:8667/heartbeat"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("sample-h", "TEST")
	req.Header.Add("Host", "www.google.com")
	req.Header.Add("Referer", "www.google.com")
	res, _ := http.DefaultClient.Do(req)
	// Verify status code
	if res.StatusCode != 200 {
		t.Errorf("%v should have returned %v but returned %v", url, 200, res.StatusCode)
	}
	return
}

func TestAPIServerWithoutCacheHit(t *testing.T) {
	// Setup request
	url := "http://localhost:8667/authz"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("sample-h", "TEST")
	req.Header.Add("Host", "www.google.com")
	req.Header.Add("Referer", "www.google.com")
	res, _ := http.DefaultClient.Do(req)
	// Verify status code
	if res.StatusCode != 200 {
		t.Errorf("%v should have returned %v but returned %v", url, 200, res.StatusCode)
	}
	return
}
