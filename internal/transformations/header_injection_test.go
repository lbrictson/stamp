package transformations

import "testing"

func TestHeaderInjectionGetNameMethod(t *testing.T) {
	rule1 := HeaderTransformation{
		Name: "test-name",
	}
	name := rule1.GetName()
	if name != rule1.Name {
		t.Errorf("Get name should have returned %v but returned %v instead", rule1.Name, name)
	}

}
