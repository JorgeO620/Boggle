package main

import (
	"fmt"
)

type Account struct { //stores data used for log-in, not public
	Type     string `json:"type,omitempty"`
	Pid      string `json:"pid,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

//self explanatory helper functions
func emailUpdate(account *Account, email string) {
	account.Email = email
}

func passwordUpdate(account *Account, pWord string) {
	account.Password = pWord
}

type Profile struct { //stores primary profile info, public
	Type         string `json:"type,omitempty"`
	Firstname    string `json:"firstname,omitempty"`
	Lastname     string `json:"lastname,omitempty"`
	currJobTitle string `json:"currjobtitle,omitempty"`
	pageLink     string `json:"pagelink,omitempty"`
}

//self explanatory helper functions
func fNameUpdate(profile *Profile, fName string) {
	profile.Firstname = fName
}

func lNameUpdate(profile *Profile, lName string) {
	profile.Lastname = lName
}

func jobUpdate(profile *Profile, job string) {
	profile.currJobTitle = job
}

type Resume struct { //Additional profile info, public
	Type           string `json:"type,omitempty"`
	pastJobTitles  string `json:"pastjobtitle,omitempty"`
	education      string `json:"education,omitempty"`
	certifications string `json:"certifications,omitempty"`
	projects       string `json:"projects,omitempty"`
}

//self explanatory helper functions
func addPastJob(resume *Resume, pastJob string) {
	resume.pastJobTitles += pastJob + ";"
}

func addEducation(resume *Resume, school string) {
	resume.education += school + ";"
}

func addCertification(resume *Resume, cert string) {
	resume.certifications += cert + ";"
}

func addProject(resume *Resume, proj string) {
	resume.projects += proj + ";"
}

type Session struct { //Info used for multiple users, not visible to user
	Type string `json:"type,omitempty"`
	Pid  string `json:"pid,omitempty"`
}

type Contacts struct {
	// List of previous coworkers/employers, public. Friends stores basic info, Links
	// stores addresses of their respective page. A profile in Friends and its respective
	// address in Links share an address
	Friends []Profile
	Links   []string
}

//adds profile to Friends, then adds link to profile page at the same index of Links
func addContact(contacts *Contacts, profile Profile) {
	contacts.Friends = append(contacts.Friends, profile)
	contacts.Links = append(contacts.Links, profile.pageLink)
}

func main() {
	fmt.Println("TESTING")
}
