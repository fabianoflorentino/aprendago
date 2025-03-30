package basecontent

import (
	"testing"
)

// TestNew tests the New function to ensure it creates a new BaseContent instance.
func TestNew(t *testing.T) {
	baseContent := New()
	if baseContent == nil {
		t.Errorf("Expected non-nil BaseContent instance, got nil")
	}
}

// TestMenu tests the Menu function to ensure it generates menu options correctly.
func TestMenu(t *testing.T) {
	baseContent := New()
	optionFlags := []string{"Option1", "Option2", "Option3"}
	optionFuncs := []func(){
		func() { t.Log("Option1 executed") },
		func() { t.Log("Option2 executed") },
		func() { t.Log("Option3 executed") },
	}

	menuOptions := baseContent.Menu(optionFlags, optionFuncs)

	if len(menuOptions) != len(optionFlags) {
		t.Errorf("Expected %d menu options, got %d", len(optionFlags), len(menuOptions))
	}

	for i, option := range menuOptions {
		if option.Options != optionFlags[i] {
			t.Errorf("Expected option flag %s, got %s", optionFlags[i], option.Options)
		}
		if option.ExecFunc == nil {
			t.Errorf("Expected non-nil ExecFunc for option %s", option.Options)
		}
	}
}
