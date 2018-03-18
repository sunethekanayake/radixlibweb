package main

import (
	"net/http"
  "fmt"
)

func loginPage(resp http.ResponseWriter, req *http.Request)  {
  if req.Method != "POST" {
    http.ServeFile(resp, req, "login.html")
  }

  var username, password string
  username = req.FormValue("username")
  password = req.FormValue("password")

  fmt.Fprint(resp, "Username:", username, " Password:", password, "\n")



}

func homePage(resp http.ResponseWriter, req *http.Request)  {

    http.ServeFile(resp, req, "index.html")
}

func main() {
  http.HandleFunc("/login", loginPage)
  http.HandleFunc("/", homePage)
  http.ListenAndServe(":8080", nil)
}
