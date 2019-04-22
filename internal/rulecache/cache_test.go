package rulecache

import (
	"testing"

	"github.com/lbrictson/stamp/internal/rules"
)

func TestCache(t *testing.T) {
	Init()
	rule1 := rules.HeaderRule{
		Name: "demo-header-1",
	}
	rule2 := rules.HeaderRule{
		Name: "demo-header-1",
	}
	rule3 := rules.HeaderRule{
		Name: "demo-header-1",
	}
	rule4 := rules.HeaderRule{
		Name: "demo-header-1",
	}
	rule5 := rules.HeaderRule{
		Name: "demo-header-1",
	}
	host1 := "google.com"
	host2 := "facebook.com"
	host3 := "example.com"
	UpdateCacheHostRules(host1, []rules.Rule{rule1, rule2, rule3, rule4, rule5})
	// test looking up host1 in cache
	found, rulesFound, _ := GetCacheHostRules(host1)
	if found != true {
		t.Errorf("%v host should have been found in cache but was not", host1)
	}
	if len(rulesFound) != 5 {
		t.Errorf("%v host rule list did not match expected count of 5", host1)
	}
	// test looking up host that doesn't exist
	found, _, _ = GetCacheHostRules("fake.com")
	if found != false {
		t.Error("fake.com should not have been found in cache but was")
	}
	// Create host2 and modify it, see if new rule is there
	UpdateCacheHostRules(host2, []rules.Rule{rule1})
	UpdateCacheHostRules(host2, []rules.Rule{rule1, rule2})
	_, rulesFound, _ = GetCacheHostRules(host2)
	if len(rulesFound) != 2 {
		t.Errorf("%v host rule list did not match expected count of 2", host2)
	}
	// test deleting from cache
	UpdateCacheHostRules(host3, []rules.Rule{rule1, rule2, rule3, rule4, rule5})
	DeleteCacheHostRules(host3)
	found, _, _ = GetCacheHostRules(host3)
	if found != false {
		t.Errorf("Should not have found host %v in cache because it was deleted", host3)
	}
}
