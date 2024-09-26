package userTest

import (
	"testing"
)

func TestUser(t *testing.T) {
	t.Log("testing user")
}
func TestAddUser(t *testing.T) {
	got := addUser({Name: "test", FullName: "test", Contact: 1234567890, Email: "test@test.com", Address: "test", Gender: "male", Password: "test"});
	want := {user}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
	t.Log("testing add user")
}
func TestUpdateUser(t *testing.T) {
	t.Log("testing update user")
}
func TestDeleteUser(t *testing.T) {
	t.Log("testing delete user")
}
