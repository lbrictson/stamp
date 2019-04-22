package transformations

import (
	"net/http"
	"testing"
)

func TestHeaderInjectionGetNameMethod(t *testing.T) {
	rule1 := HeaderTransformation{
		Name: "test-name",
	}
	name := rule1.GetName()
	if name != rule1.Name {
		t.Errorf("Get name should have returned %v but returned %v instead", rule1.Name, name)
	}

}

func TestHeaderInjectionEval(t *testing.T) {
	// Happy path
	headerRule1 := HeaderTransformation{
		Name: "client-rule",
		Header: "referer",
		Value: "example.com",
		CaseSensitive: false,
		ExactMatch: true,
		InjectedHeaders: map[string]string{"client-slug": "example-client"},
	}
	request, _ := http.NewRequest("GET", "/authz", nil)
	request.Header.Add("test-header", "pass")
	request.Header.Add("referer", "example.com")
	headersToAdd := headerRule1.Eval(request)
	if len(headersToAdd) != 1 {
		t.Errorf("rule %v should have returned 1 header to add but returned %v", headerRule1.Name, len(headersToAdd))
	}
	// Non case match
	headerRule2 := HeaderTransformation{
		Name: "client-rule",
		Header: "referer",
		Value: "EXAMPLE.com",
		CaseSensitive: true,
		ExactMatch: true,
		InjectedHeaders: map[string]string{"client-slug": "example-client"},
	}
	request, _ = http.NewRequest("GET", "/authz", nil)
	request.Header.Add("test-header", "pass")
	request.Header.Add("referer", "example.com")
	headersToAdd = headerRule2.Eval(request)
	if len(headersToAdd) != 0 {
		t.Errorf("rule %v should have returned 0 header to add but returned %v", headerRule2.Name, len(headersToAdd))
	}
	// Non exact match
	headerRule3 := HeaderTransformation{
		Name: "client-rule",
		Header: "referer",
		Value: "EXAMPLE.com",
		CaseSensitive: true,
		ExactMatch: false,
		InjectedHeaders: map[string]string{"client-slug": "example-client"},
	}
	request, _ = http.NewRequest("GET", "/authz", nil)
	request.Header.Add("test-header", "pass")
	request.Header.Add("referer", "example.com")
	headersToAdd = headerRule3.Eval(request)
	if len(headersToAdd) != 0 {
		t.Errorf("rule %v should have returned 0 header to add but returned %v", headerRule3.Name, len(headersToAdd))
	}
	return
}