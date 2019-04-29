package ratelimit

// IPPayload contains an ip and domain to be verified against the rate limit rules
type IPPayload struct {
	IPAddr  string
	Domain  string
	Limit   int
	Minutes int
}
