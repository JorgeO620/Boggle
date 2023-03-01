package main

import (
	"reflect"
	"testing"
)

func TestFNameUpdate(t *testing.T) {
	tProfile := Profile{
		Type:         "profile",
		Firstname:    "John",
		Lastname:     "Smith",
		currJobTitle: "Doctor",
		pageLink:     "www.website.com",
	}
	fProfile := Profile{
		Type:         "profile",
		Firstname:    "James",
		Lastname:     "Smith",
		currJobTitle: "Doctor",
		pageLink:     "www.website.com",
	}
	var x *Profile
	x = &tProfile
	fNameUpdate(x, "James")
	got := tProfile
	want := fProfile
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLNameUpdate(t *testing.T) {
	tProfile := Profile{
		Type:         "profile",
		Firstname:    "John",
		Lastname:     "Smith",
		currJobTitle: "Doctor",
		pageLink:     "www.website.com",
	}
	fProfile := Profile{
		Type:         "profile",
		Firstname:    "John",
		Lastname:     "Denver",
		currJobTitle: "Doctor",
		pageLink:     "www.website.com",
	}
	var x *Profile
	x = &tProfile
	lNameUpdate(x, "Denver")
	got := tProfile
	want := fProfile
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestJobUpdate(t *testing.T) {
	tProfile := Profile{
		Type:         "profile",
		Firstname:    "John",
		Lastname:     "Smith",
		currJobTitle: "Doctor",
		pageLink:     "www.website.com",
	}
	fProfile := Profile{
		Type:         "profile",
		Firstname:    "John",
		Lastname:     "Smith",
		currJobTitle: "Funemployed",
		pageLink:     "www.website.com",
	}
	var x *Profile
	x = &tProfile
	jobUpdate(x, "Funemployed")
	got := tProfile
	want := fProfile
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestEmailUpdate(t *testing.T) {
	tAccount := Account{
		Type:     "Account",
		Pid:      "00000001",
		Email:    "123@gmail.com",
		Password: "password",
	}
	fAccount := Account{
		Type:     "Account",
		Pid:      "00000001",
		Email:    "abc@hotmail.com",
		Password: "password",
	}
	var x *Account
	x = &tAccount
	emailUpdate(x, "abc@hotmail.com")
	got := tAccount
	want := fAccount
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPasswordUpdate(t *testing.T) {
	tAccount := Account{
		Type:     "Account",
		Pid:      "00000001",
		Email:    "123@gmail.com",
		Password: "password",
	}
	fAccount := Account{
		Type:     "Account",
		Pid:      "00000001",
		Email:    "123@gmail.com",
		Password: "securePassword1!",
	}
	var x *Account
	x = &tAccount
	passwordUpdate(x, "securePassword1!")
	got := tAccount
	want := fAccount
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAddContact(t *testing.T) {
	tProfile := Profile{
		Type:         "profile",
		Firstname:    "John",
		Lastname:     "Smith",
		currJobTitle: "Doctor",
		pageLink:     "www.website.com",
	}
	var tContacts Contacts
	fContacts := Contacts{
		Friends: []Profile{tProfile},
		Links:   []string{"www.website.com"},
	}
	var x = &tContacts
	addContact(x, tProfile)
	got := tContacts
	want := fContacts
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
