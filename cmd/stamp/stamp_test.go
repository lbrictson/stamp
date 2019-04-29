package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/lbrictson/stamp/internal/rulecache"
	"github.com/lbrictson/stamp/internal/rules"

	"github.com/lbrictson/stamp/internal/logging"
)

func TestRunningStampAPI(t *testing.T) {
	os.Args = []string{"stamp", "api"}
	go main()
	// Verify API server started with just a simple heartbeat
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
	// Verify cache is accepting new records
	rulecache.UpdateCacheHostRules("localhost", []rules.Rule{})
	// Logger will panic if it failed to start
	logging.Logger.Info("Stamp should be running now")
	return

}

func TestRunningStampRate(t *testing.T) {
	os.Args = []string{"stamp", "rate"}
	go main()
	// Verify API server started with just a simple heartbeat
	// Setup request
	url := "http://localhost:8668/heartbeat"
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
	// Logger will panic if it failed to start
	logging.Logger.Info("Stamp should be running now")
	return

}
