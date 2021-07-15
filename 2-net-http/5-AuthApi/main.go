package main

import (
	"fmt"
	"net/http"

	helper "./helpers"
)

func main() {
	mux := http.NewServeMux()

	userName, email,password, passwordConfirm := "","","",""
	//register
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		userName = r.FormValue("userName")
		email = r.FormValue("email")
		password = r.FormValue("password")
		passwordConfirm = r.FormValue("passwordConfirm")

		userNameCheck := helper.IsEmpty(userName)
		emailCheck := helper.IsEmpty(email)
		passwordCheck := helper.IsEmpty(password)
		passwordConfirmCheck := helper.IsEmpty(passwordConfirm)

		if userNameCheck || emailCheck || passwordCheck || passwordConfirmCheck {
			fmt.Fprintf(w, "Error any field is empty!")
			return
		}

		if password == passwordConfirm{
			fmt.Fprintf(w, "Successful")
		}else{
			fmt.Fprintf(w, "Passwords must be same")
		}
	})

	//login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()

			userName = r.FormValue("userName")
			password = r.FormValue("password")

		userNameCheck := helper.IsEmpty(userName)
		passwordCheck := helper.IsEmpty(password)
		if userNameCheck || passwordCheck {
			fmt.Fprintf(w, "Error any field is empty!")
			return
		}

		dbUser := "tunahan"
		dbPassw := "123123"

		if dbUser == userName && dbPassw == password{
			fmt.Fprint(w,"Success!")
		}else {
			fmt.Fprint(w,"Check your infos!!")

		}
	})

	http.ListenAndServe(":8080", mux)
}

//for key, value := range r.Form{
//	fmt.Printf("%s = %s \n", key, value)
//}