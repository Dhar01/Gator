package config

import "testing"

func TestSetUser(t *testing.T) {
	cfg := Config{}

	err := cfg.SetUser("test_user")
	if err != nil {
		t.Fatalf("SetUser failed: %v", err)
	}

	if cfg.User_name != "test_user" {
		t.Errorf("expected username to be 'test_user', got %s", cfg.User_name)
	}
}
