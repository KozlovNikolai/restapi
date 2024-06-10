package model

import "testing"

// TestUser ...
func TestUser(_ *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
