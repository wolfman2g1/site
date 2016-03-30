package main

import (
	"io"
	"net/http"
	"os"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("signup", signupHandler)
	http.ListenAndServe(":6060", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "test")
	f, err := os.Open("./public/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.Copy(w, f)

}
// handle the about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./public/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.Copy(w, f)
}

// handle the user account sign up
func signupHandler(w http.ResponseWriter, r *http.Request) {
     err :=	r.ParseForm() // Parse the url perameters
	if err != nil{
		// if there is an error do something
		// TODO figure out how to handle this if it evaluates to be true. probably a redirect, must maintain the state of the form

	}
	// parse the form elements
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	password2 := r.PostFormValue("password2")


	createAccount( username,password, password2)

	// TODO print sucess or fail message after the create account function runs
}
// handle the account creation
func createAccount( username string, password string, password2 string) {
	// TODO check if user account exists. if so alert user.

	//TODO  should i be handling password hashing on the server or should this be done in the client ?
	if password != password2 {
		// return some sort of message to the user.. Javascript?
		// TODO figure out how to handle this if it evaluates to be true
	}


	// hash the password
	hashedpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
//TODO make a comparison function for user sign in

	return
}
