package main

import (
	"net/http"
  "fmt"
)

func loginPage(w http.ResponseWriter, r *http.Request)  {
  if r.Method != "POST" {
    http.ServeFile(w, r, "login.html")
  }

  var username, password string
  username = r.FormValue("username")
  password = r.FormValue("password")
  fmt.Fprint(w, "Username:", username, " Password:", password, "\n")
}

func homePage(w http.ResponseWriter, r *http.Request)  {
    http.ServeFile(w, r, "./RadixRtrGui/public_html/index.html")
}

func main() { 
  fmt.Print("Hello Go Lang  !!!!!!!")
  http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./RadixRtrGui/public_html/js"))))
  http.HandleFunc("/", homePage)
  http.HandleFunc("/login", loginPage)
  http.HandleFunc("/addbook", addBook)
  http.ListenAndServe(":8080", nil)
}
