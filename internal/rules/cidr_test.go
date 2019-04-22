package rules

import (
	"net"
	"net/http"
	"testing"

	"github.com/yl2chen/cidranger"
)

func TestCIDREval(t *testing.T) {
	// Setup CIDR rangers for rules
	_, network1, _ := net.ParseCIDR("10.0.0.0/8")
	ranger1 := cidranger.NewPCTrieRanger()
	ranger1.Insert(cidranger.NewBasicRangerEntry(*network1))
	_, network2, _ := net.ParseCIDR("10.0.0.0/8")
	ranger2 := cidranger.NewPCTrieRanger()
	ranger2.Insert(cidranger.NewBasicRangerEntry(*network2))
	_, network3, _ := net.ParseCIDR("10.0.0.77/32")
	ranger3 := cidranger.NewPCTrieRanger()
	ranger3.Insert(cidranger.NewBasicRangerEntry(*network3))
	_, network4, _ := net.ParseCIDR("10.0.0.77/32")
	ranger4 := cidranger.NewPCTrieRanger()
	ranger4.Insert(cidranger.NewBasicRangerEntry(*network4))
	// Make mock rules using CIDRs
	ruleTable := make(map[CIDRRule]bool)
	mockRuleOne := CIDRRule{
		Name:        "easy-pass",
		WhiteListed: true,
		CIDRRange:   ranger1,
	}
	ruleTable[mockRuleOne] = false
	mockRuleTwo := CIDRRule{
		Name:        "block-cidr-blacklist",
		WhiteListed: false,
		CIDRRange:   ranger2,
	}
	ruleTable[mockRuleTwo] = true
	mockRuleThree := CIDRRule{
		Name:        "pass-single-ip",
		WhiteListed: true,
		CIDRRange:   ranger3,
	}
	ruleTable[mockRuleThree] = false
	mockRuleFour := CIDRRule{
		Name:        "block-single-ip",
		WhiteListed: false,
		CIDRRange:   ranger4,
	}
	ruleTable[mockRuleFour] = true
	request, _ := http.NewRequest("GET", "/authz", nil)
	// Mock the IP address the request is coming from
	request.RemoteAddr = "10.0.0.77"
	for k, v := range ruleTable {
		blocked := k.Eval(request)
		if blocked != v {
			t.Errorf("CIDRRule rule %v failed because blocked was %v but should have been %v", k.Name, blocked, v)
		}
	}
}

func TestCIDRGetNameMethod(t *testing.T) {
	rule1 := CIDRRule{
		Name:        "test-name",
		WhiteListed: true,
	}
	name := rule1.GetName()
	if name != rule1.Name {
		t.Errorf("Get name should have returned %v but returned %v instead", rule1.Name, name)
	}

}
