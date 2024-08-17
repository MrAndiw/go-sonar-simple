package pkg

import "testing"

func TestGetUser(t *testing.T) {
	module := NewUserModule()
	user, err := module.GetUser("mrandiiw@gmail.com")
	if err != nil {
		t.Errorf("error when calling GetUser: %v", err)
	}

	result := user
	expected := "The User Email is mrandiiw@gmail.com"

	if result != expected {
		t.Errorf("getUser failed, expected %s, got %s", expected, result)
	}
}

func TestGetProfile(t *testing.T) {
	module := NewUserModule()
	user, err := module.GetProfile("mrandiiw@gmail.com")
	if err != nil {
		t.Errorf("error when calling GetUser: %v", err)
	}

	result := user
	expected := "The User Profile is mrandiiw@gmail.com"

	if result != expected {
		t.Errorf("getUser failed, expected %s, got %s", expected, result)
	}
}
