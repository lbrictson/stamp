package api

import (
	"net/http"
	"testing"

	"github.com/lbrictson/stamp/internal/logging"
	"github.com/lbrictson/stamp/internal/rulecache"
	"github.com/lbrictson/stamp/internal/rules"
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

func TestAPIServerWithCacheHitBlock(t *testing.T) {
	// Setup rule
	mockRuleOne := rules.HeaderRule{
		Name:          "this-header-is-not-okay",
		WhiteListed:   true,
		Header:        "test-header",
		Value:         "pass",
		ExactMatch:    true,
		CaseSensitive: false,
	}
	// Put rule in cache
	rulecache.UpdateCacheHostRules("localhost", []rules.Rule{mockRuleOne})
	// Setup request
	url := "http://localhost:8667/authz"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("test-header", "notokay")
	req.Header.Add("Host", "www.google.com")
	req.Header.Add("Referer", "localhost")
	res, _ := http.DefaultClient.Do(req)
	// Verify status code
	if res.StatusCode != 501 {
		t.Errorf("%v should have returned %v because of rule %v but returned %v", url, 501, mockRuleOne.Name, res.StatusCode)
	}
	return
}

func TestAPIServerWithCacheHitOkay(t *testing.T) {
	// Setup rule
	mockRuleOne := rules.HeaderRule{
		Name:          "this-header-is-okay",
		WhiteListed:   true,
		Header:        "test-header",
		Value:         "pass",
		ExactMatch:    true,
		CaseSensitive: false,
	}
	// Put rule in cache
	rulecache.UpdateCacheHostRules("localhost", []rules.Rule{mockRuleOne})
	// Setup request
	url := "http://localhost:8667/authz"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("test-header", "pass")
	req.Header.Add("Host", "www.google.com")
	req.Header.Add("Referer", "localhost")
	res, _ := http.DefaultClient.Do(req)
	// Verify status code
	if res.StatusCode != 200 {
		t.Errorf("%v should have returned %v because of rule %v but returned %v", url, 200, mockRuleOne.Name, res.StatusCode)
	}
	return
}
