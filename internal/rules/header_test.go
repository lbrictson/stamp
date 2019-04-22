package rules

import (
	"net/http"
	"testing"
)

func TestHeaderEval(t *testing.T) {
	ruleTable := make(map[HeaderRule]bool)
	mockRuleOne := HeaderRule{
		Name:          "easy-pass",
		WhiteListed:   true,
		Header:        "test-header",
		Value:         "pass",
		ExactMatch:    true,
		CaseSensitive: false,
	}
	ruleTable[mockRuleOne] = false
	mockRuleTwo := HeaderRule{
		Name:          "fail-blacklist-casesensitive-false",
		WhiteListed:   false,
		Header:        "test-header",
		Value:         "PaSS",
		ExactMatch:    false,
		CaseSensitive: false,
	}
	ruleTable[mockRuleTwo] = true
	mockRuleThree := HeaderRule{
		Name:          "fail-because-case-mismatch",
		WhiteListed:   true,
		Header:        "test-header",
		Value:         "PAsS",
		ExactMatch:    true,
		CaseSensitive: true,
	}
	ruleTable[mockRuleThree] = true
	mockRuleFour := HeaderRule{
		Name:          "pass-because-case-mismatch-but-whitelisted",
		WhiteListed:   false,
		Header:        "test-header",
		Value:         "PASS",
		ExactMatch:    true,
		CaseSensitive: true,
	}
	ruleTable[mockRuleFour] = false
	request, _ := http.NewRequest("GET", "/authz", nil)
	request.Header.Add("test-header", "pass")
	for k, v := range ruleTable {
		blocked := k.Eval(request)
		if blocked != v {
			t.Errorf("Header rule %v failed because blocked was %v but should have been %v", k.Name, blocked, v)
		}
	}
}

func TestHeaderGetNameMethod(t *testing.T) {
	rule1 := HeaderRule{
		Name:        "test-name",
		WhiteListed: true,
	}
	name := rule1.GetName()
	if name != rule1.Name {
		t.Errorf("Get name should have returned %v but returned %v instead", rule1.Name, name)
	}

}
