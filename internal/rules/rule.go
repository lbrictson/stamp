package rules

import "net/http"

// Rule is a shared interface that all rules must implement to be valid
// where eval() returns true if the rule is matched with a block and
// false if the rule was missed
// Generally you can treat the bool returned as a "should I block this" value
type Rule interface {
	Eval(c *http.Request) bool
	GetName() string
}
