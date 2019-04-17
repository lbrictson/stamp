package rules

import (
	"net"
	"net/http"

	"github.com/yl2chen/cidranger"
)

// CIDRRule inspects the request remote address then passes or fails
// the rule.  WhiteListed indicates if the ip cidr value matching is a positive or negative
type CIDRRule struct {
	Name        string
	WhiteListed bool
	Block       cidranger.Ranger // A valid CIDR block, parsed from Kubernetes so we know its valid before its here
}

// Eval implements the rule interface
func (r CIDRRule) Eval(c *http.Request) bool {
	// Pull our remote calling IP
	remote := c.RemoteAddr
	// Check in in specified CIDR
	contains, _ := r.Block.Contains(net.ParseIP(remote))
	// Inverse whitelist
	if r.WhiteListed {
		return !contains
	}
	return contains
}
