package main
import (
  "github.com/gorilla/mux"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"
)
type Account struct {
  Username string `json:"username"`
  Password string `json:"password"`
}
type JobPost struct {
  PostID string `json:"postid"`
  Company string `json:"company"`
  PostTitle string `json:"posttitle"`
  Post string `json:"post"`
  Username string `json:"username"` 
}
var db *sql.DB
var err error
func main() {
db, err = sql.Open("mysql", "root:David620@tcp(127.0.0.1:3306)/demo")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()
  router := mux.NewRouter()
  router.HandleFunc("/accounts", getPosts).Methods("GET")
  router.HandleFunc("/accounts", createPost).Methods("POST")
  router.HandleFunc("/accounts/{username}", getPost).Methods("GET")
  router.HandleFunc("/accounts/{username}", updatePost).Methods("PUT")
  router.HandleFunc("/accounts/{username}", deletePost).Methods("DELETE")
  router.HandleFunc("/posts", getJobPosts).Methods("GET")
  router.HandleFunc("/posts", createJobPost).Methods("POST")
  //router.HandleFunc("/posts/{postid}", getIDJobPost).Methods("GET")
  router.HandleFunc("/posts/{username}", getUserJobPost).Methods("GET")
  router.HandleFunc("/posts/{postid}", updateJobPost).Methods("PUT")
  router.HandleFunc("/posts/{postid}", deleteJobPost).Methods("DELETE")
  http.ListenAndServe(":8000", router)
}
func getPosts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var accounts []Account
  result, err := db.Query("SELECT username, password from accounts")
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var account Account
    err := result.Scan(&account.Username, &account.Password)
    if err != nil {
      panic(err.Error())
    }
    accounts = append(accounts, account)
  }
  json.NewEncoder(w).Encode(accounts)
}
func getJobPosts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var jobposts []JobPost
  result, err := db.Query("SELECT postid, company, posttitle, post, username from posts")
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var jobpost JobPost
    err := result.Scan(&jobpost.PostID, &jobpost.Company, &jobpost.PostTitle, &jobpost.Post, &jobpost.Username)
    if err != nil {
      panic(err.Error())
    }
    jobposts = append(jobposts, jobpost)
  }
  json.NewEncoder(w).Encode(jobposts)
}
func createPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  stmt, err := db.Prepare("INSERT INTO accounts(username, password) VALUES(?, ?)")
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  username := keyVal["username"]
  password := keyVal["password"]
  _, err = stmt.Exec(username, password)
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "New account was created")
}
func createJobPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  stmt, err := db.Prepare("INSERT INTO posts(company, posttitle, post, username) VALUES(?, ?, ?, ?)")
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  company := keyVal["company"]
  posttitle := keyVal["posttitle"]
  post := keyVal["post"]
  username := keyVal["username"]
  _, err = stmt.Exec(company, posttitle, post, username)
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "New post was created")
}
func getPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  result, err := db.Query("SELECT username, password FROM accounts WHERE username = ?", params["username"])
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  var account Account
  for result.Next() {
    err := result.Scan(&account.Username, &account.Password)
    if err != nil {
      panic(err.Error())
    }
  }
  json.NewEncoder(w).Encode(account)
}
// func getIDJobPost(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
//   params := mux.Vars(r)
//   result, err := db.Query("SELECT postid, company, posttitle, post, username FROM posts WHERE postid = ?", params["postid"])
//   if err != nil {
//     panic(err.Error())
//   }
//   defer result.Close()
//   var jobpost JobPost
//   for result.Next() {
//     err := result.Scan(&jobpost.PostID, &jobpost.Company, &jobpost.PostTitle, &jobpost.Post, &jobpost.Username)
//     if err != nil {
//       panic(err.Error())
//     }
//   }
//   json.NewEncoder(w).Encode(jobpost)
// }
 func getUserJobPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var jobposts []JobPost
  params := mux.Vars(r)
  result, err := db.Query("SELECT postid, company, posttitle, post, username FROM posts WHERE username = ?", params["username"])
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var jobpost JobPost
    err := result.Scan(&jobpost.PostID, &jobpost.Company, &jobpost.PostTitle, &jobpost.Post, &jobpost.Username)
    if err != nil {
      panic(err.Error())
    }
    jobposts = append(jobposts, jobpost)
  }
  json.NewEncoder(w).Encode(jobposts)
}
func updatePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  stmt, err := db.Prepare("UPDATE accounts SET password = ? WHERE username = ?")
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  newTitle := keyVal["password"]
  _, err = stmt.Exec(newTitle, params["username"])
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "Account with username = %s was updated", params["username"])
}
func updateJobPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  stmt, err := db.Prepare("UPDATE posts SET post = ? WHERE postid = ?")
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  newPost := keyVal["post"]
  _, err = stmt.Exec(newPost, params["postid"])
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "Post with post ID = %s was updated", params["postid"])
}
func deletePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  stmt, err := db.Prepare("DELETE FROM accounts WHERE username = ?")
  if err != nil {
    panic(err.Error())
  }
  _, err = stmt.Exec(params["username"])
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "Account with username = %s was deleted", params["username"])
}
func deleteJobPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  stmt, err := db.Prepare("DELETE FROM posts WHERE postid = ?")
  if err != nil {
    panic(err.Error())
  }
  _, err = stmt.Exec(params["postid"])
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "Post with post ID = %s was deleted", params["postid"])
}