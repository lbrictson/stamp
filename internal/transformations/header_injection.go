package transformations

import (
	"net/http"
	"strings"
)

// HeaderTransformation will match a header and value combination and inject the specified header and value
type HeaderTransformation struct {
	Name            string
	Header          string
	Value           string
	InjectedHeaders map[string]string
	ExactMatch      bool
	CaseSensitive   bool
}

// Eval Implements the transformation interface returning a map of headers to inject into a request if the request matches
// the transformation rules
func (r HeaderTransformation) Eval(c *http.Request) map[string]string {
	match := false
	headerValue := c.Header.Get(r.Header)
	valueToMatch := r.Value
	if r.CaseSensitive == false {
		valueToMatch = strings.ToLower(valueToMatch)
		headerValue = strings.ToLower(headerValue)
	}
	if r.ExactMatch {
		match = (valueToMatch == headerValue)
	} else {
		match = strings.Contains(headerValue, valueToMatch)
	}
	if match == true {
		return r.InjectedHeaders
	}
	return nil
}

// GetName just returns the name of the rule
func (r HeaderTransformation) GetName() string {
	return r.Name
}
