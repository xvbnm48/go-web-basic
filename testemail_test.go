package main

import "testing"

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not email")
	}

	_, err = IsEmail("fariz@gmail.com")
	if err != nil {
		t.Error("fariz@gmail.com is your email")
	}
}
