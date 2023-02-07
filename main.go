package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/couchbase/gocb"
	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Type     string `json:"type,omitempty"`
	Pid      string `json:"pid,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Profile struct {
	Type      string `json:"type,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	currJobTitle string `json:"currjobtitle,omitempty"`
}

type Resume struct {
	Type      string `json:"type,omitempty"`
	pastJobTitles string `json:"pastjobtitle,omitempty"`
	education string `json:"education,omitempty"`
	certifications string `json:"certifications,omitempty"`
	projects string `json:"projects,omitempty"`
}

type Session struct {
	Type string `json:"type,omitempty"`
	Pid  string `json:"pid,omitempty"`
}

// type Blog struct {
// 	Type      string `json:"type,omitempty"`
// 	Pid       string `json:"pid,omitempty"`
// 	Title     string `json:"title,omitempty"`
// 	Content   string `json:"content,omitempty"`
// 	Timestamp int    `json:"timestamp,omitempty"`
// }

var bucket *gocb.Bucket

func Validate(next http.HandlerFunc) http.HandlerFunc {}

func RegisterEndpoint(w http.ResponseWriter, req *http.Request) {}
func LoginEndpoint(w http.ResponseWriter, req *http.Request) {}
func AccountEndpoint(w http.ResponseWriter, req *http.Request) {}
func BlogsEndpoint(w http.ResponseWriter, req *http.Request) {}
func BlogEndpoint(w http.ResponseWriter, req *http.Request) {}

func main() {
	fmt.Println("Starting the Go server...")
	router := mux.NewRouter()
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ = cluster.OpenBucket("default", "")
	router.HandleFunc("/account", RegisterEndpoint).Methods("POST")
	router.HandleFunc("/login", LoginEndpoint).Methods("POST")
	router.HandleFunc("/account", Validate(AccountEndpoint)).Methods("GET")
	// router.HandleFunc("/blogs", Validate(BlogsEndpoint)).Methods("GET")
	// router.HandleFunc("/blog", Validate(BlogEndpoint)).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}