package rules

import (
	"net/http"
	"strings"
)

// HeaderRule inspects a header and it's value and then passes or fails
// the rule.  WhiteListed indicates if the header value matching is a positive or negative
type HeaderRule struct {
	Name          string
	Header        string
	WhiteListed   bool
	Value         string
	ExactMatch    bool
	CaseSensitive bool
}

// Eval Implements the rule interface
func (r HeaderRule) Eval(c *http.Request) bool {
	inValue := false
	// Pull request header in
	requestHeader := c.Header.Get(r.Header)
	// Flip case sensitive
	if r.CaseSensitive != true {
		r.Value = strings.ToLower(r.Value)
		requestHeader = strings.ToLower(requestHeader)
	}
	// Perform matching
	if r.ExactMatch {
		inValue = (requestHeader == r.Value)
	} else {
		inValue = strings.Contains(requestHeader, r.Value)
	}
	// Do whitelisted swap
	if r.WhiteListed == true {
		return !inValue
	}
	return inValue
}

// GetName just returns the name of the rule
func (r HeaderRule) GetName() string {
	return r.Name
}
