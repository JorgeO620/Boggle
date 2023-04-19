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
  Friends []string
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
  http.ListenAndServe(":8000", router)
}

// encodes slice of all usernames and their mathching passwords
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

// creates new account
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

// encodes account with desired username
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

// updates the password of a specific account based on account's username
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

// delete accounts with desired username
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

// encodes friends and links slices of account with desired username
func friendsPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT friends FROM accounts WHERE username = ?", params["username"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var account Account
	for result.Next() {
		err := result.Scan(&account.Friends)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(account)
}
