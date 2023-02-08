package main

import (
	"encoding/json"
	"fmt"
	// "log"
	// "net/http"
	// "strings"
	// "time"

	// "golang.org/x/crypto/bcrypt"

	"github.com/couchbase/gocb"
	// "github.com/gorilla/context"
	// "github.com/gorilla/handlers"
	// "github.com/gorilla/mux"
	// uuid "github.com/satori/go.uuid"
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
}

// type Session struct {
// 	Type string `json:"type,omitempty"`
// 	Pid  string `json:"pid,omitempty"`
// }

// type Blog struct {
// 	Type      string `json:"type,omitempty"`
// 	Pid       string `json:"pid,omitempty"`
// 	Title     string `json:"title,omitempty"`
// 	Content   string `json:"content,omitempty"`
// 	Timestamp int    `json:"timestamp,omitempty"`
// }

// var bucket *gocb.Bucket

// func Validate(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		authorizationHeader := req.Header.Get("authorization")
// 		if authorizationHeader != "" {
// 			bearerToken := strings.Split(authorizationHeader, " ")
// 			if len(bearerToken) == 2 {
// 				var session Session
// 				_, err := bucket.Get(bearerToken[1], &session)
// 				if err != nil {
// 					w.WriteHeader(401)
// 					w.Write([]byte(err.Error()))
// 					return
// 				}
// 				context.Set(req, "pid", session.Pid)
// 				bucket.Touch(bearerToken[1], 0, 3600)
// 				next(w, req)
// 			}
// 		} else {
// 			w.WriteHeader(401)
// 			w.Write([]byte("An authorization header is required"))
// 			return
// 		}
// 	})
// }

// func RegisterEndpoint(w http.ResponseWriter, req *http.Request) {
// 	var data map[string]interface{}
// 	_ = json.NewDecoder(req.Body).Decode(&data)
// 	id := uuid.NewV4().String()
// 	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), 10)
// 	account := Account{
// 		Type:     "account",
// 		Pid:      id,
// 		Email:    data["email"].(string),
// 		Password: string(passwordHash),
// 	}
// 	profile := Profile{
// 		Type:      "profile",
// 		Firstname: data["firstname"].(string),
// 		Lastname:  data["lastname"].(string),
// 	}
// 	_, err := bucket.Insert(id, profile, 0)
// 	if err != nil {
// 		w.WriteHeader(401)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	_, err = bucket.Insert(data["email"].(string), account, 0)
// 	if err != nil {
// 		w.WriteHeader(401)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	json.NewEncoder(w).Encode(account)
// }

// func LoginEndpoint(w http.ResponseWriter, req *http.Request) {
// 	var data Account
// 	var account Account
// 	_ = json.NewDecoder(req.Body).Decode(&data)
// 	_, err := bucket.Get(data.Email, &account)
// 	if err != nil {
// 		w.WriteHeader(401)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(data.Password))
// 	if err != nil {
// 		w.WriteHeader(401)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	session := Session{
// 		Type: "session",
// 		Pid:  account.Pid,
// 	}
// 	var result map[string]interface{}
// 	result = make(map[string]interface{})
// 	result["sid"] = uuid.NewV4().String()
// 	_, err = bucket.Insert(result["sid"].(string), &session, 3600)
// 	if err != nil {
// 		w.WriteHeader(401)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	json.NewEncoder(w).Encode(result)
// }

// func AccountEndpoint(w http.ResponseWriter, req *http.Request) {
// 	pid := context.Get(req, "pid").(string)
// 	var profile Profile
// 	_, err := bucket.Get(pid, &profile)
// 	if err != nil {
// 		w.WriteHeader(401)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	json.NewEncoder(w).Encode(profile)
// }

// func BlogsEndpoint(w http.ResponseWriter, req *http.Request) {
// 	var n1qlParams []interface{}
// 	n1qlParams = append(n1qlParams, context.Get(req, "pid").(string))
// 	query := gocb.NewN1qlQuery("SELECT `" + bucket.Name() + "`.* FROM `" + bucket.Name() + "` WHERE type = 'blog' AND pid = $1")
// 	query.Consistency(gocb.RequestPlus)
// 	rows, err := bucket.ExecuteN1qlQuery(query, n1qlParams)
// 	if err != nil {
// 		w.WriteHeader(401)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	var row Blog
// 	var result []Blog
// 	for rows.Next(&row) {
// 		result = append(result, row)
// 		row = Blog{}
// 	}
// 	rows.Close()
// 	if result == nil {
// 		result = make([]Blog, 0)
// 	}
// 	json.NewEncoder(w).Encode(result)
// }

// func BlogEndpoint(w http.ResponseWriter, req *http.Request) {
// 	var blog Blog
// 	_ = json.NewDecoder(req.Body).Decode(&blog)
// 	blog.Type = "blog"
// 	blog.Pid = context.Get(req, "pid").(string)
// 	blog.Timestamp = int(time.Now().Unix())
// 	_, err := bucket.Insert(uuid.NewV4().String(), blog, 0)
// 	if err != nil {
// 		w.WriteHeader(401)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	json.NewEncoder(w).Encode(blog)
// }

func main() {
	// fmt.Println("Starting the Go server...")
	// router := mux.NewRouter()
	cluster, _ := gocb.Connect("couchbase://127.0.0.1")
	bucket, _ := cluster.OpenBucket("default", "")
	var Professional Profile
 	bucket.Upsert("JohnSmith", Profile{
		Type: "Engineer",
		Firstname: "John",
		Lastname: "Smith",
	}, 0)
	bucket.Get("JohnSmith", &Professional)
	jsonBytes, _ := json.Marshal(Professional)
	fmt.Println(string(jsonBytes))
	// router.HandleFunc("/account", RegisterEndpoint).Methods("POST")
	// router.HandleFunc("/login", LoginEndpoint).Methods("POST")
	// router.HandleFunc("/account", Validate(AccountEndpoint)).Methods("GET")
	// router.HandleFunc("/blogs", Validate(BlogsEndpoint)).Methods("GET")
	// router.HandleFunc("/blog", Validate(BlogEndpoint)).Methods("POST")
	// log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
