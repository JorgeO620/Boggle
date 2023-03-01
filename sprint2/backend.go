package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/couchbase/gocb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Account struct { //stores data used for log-in, not public
	Type     string `json:"type,omitempty"`
	Pid      string `json:"pid,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

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

func addContact(contacts *Contacts, profile Profile) {
	contacts.Friends = append(contacts.Friends, profile)
	contacts.Links = append(contacts.Links, profile.pageLink)
}

func server(sHost string, sType string, sPort string) {
	server, err := net.Listen(sType, sHost+":"+sPort)
	if err != nil {
		os.Exit(1)
	}
	defer server.Close()
	for {
		connection, err := server.Accept()
		if err != nil {
			os.Exit(1)
		}
	}
}

// type Blog struct {
// 	Type      string `json:"type,omitempty"`
// 	Pid       string `json:"pid,omitempty"`
// 	Title     string `json:"title,omitempty"`
// 	Content   string `json:"content,omitempty"`
// 	Timestamp int    `json:"timestamp,omitempty"`
// }

var bucket *gocb.Bucket

//func Validate(next http.HandlerFunc) http.HandlerFunc {}

func RegisterEndpoint(w http.ResponseWriter, req *http.Request) {
	var data map[string]interface{}
	_ = json.NewDecoder(req.Body).Decode(&data)
	id := uuid.NewV4().String()
	passHash, _ := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), 20)
	account := Account{
		Type:     "account",
		Pid:      id,
		Email:    data["email"].(string),
		Password: string(passHash),
	}
	profile := Profile{
		Type:         "profile",
		Firstname:    data["firstname"].(string),
		Lastname:     data["lastname"].(string),
		currJobTitle: data["jobtitle"].(string),
	}
	_, err := bucket.Insert(id, profile, 0)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}
	_, err = bucket.Insert(data["email"].(string), account, 0)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(account)
}
func LoginEndpoint(w http.ResponseWriter, req *http.Request)   {}
func AccountEndpoint(w http.ResponseWriter, req *http.Request) {}
func BlogsEndpoint(w http.ResponseWriter, req *http.Request)   {}
func BlogEndpoint(w http.ResponseWriter, req *http.Request)    {}

func main() {
	fmt.Println("Starting the Go server...")
	router := mux.NewRouter()
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ = cluster.OpenBucket("default", "")
	//router.HandleFunc("/account", RegisterEndpoint).Methods("POST")
	//router.HandleFunc("/login", LoginEndpoint).Methods("POST")
	//router.HandleFunc("/account", Validate(AccountEndpoint)).Methods("GET")
	// router.HandleFunc("/blogs", Validate(BlogsEndpoint)).Methods("GET")
	// router.HandleFunc("/blog", Validate(BlogEndpoint)).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
