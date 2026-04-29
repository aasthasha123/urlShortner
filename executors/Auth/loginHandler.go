package auth_login

import (
	"fmt"
	"log"
	"net/http"
	"urlShortner/auth"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	password := r.FormValue("password")

	passwordDb := "password"
	hashpass, err := auth.HashPassword(passwordDb)
	if err != nil {
		log.Panic(err)
	}
	checkPass := auth.CheckPassword(password, hashpass)
	if !checkPass {
		http.Error(w, "Password not matching", http.StatusUnauthorized)
	}
	token, err := auth.GenerateToken(userName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)

		return

	}
	w.Write([]byte(token))

}
