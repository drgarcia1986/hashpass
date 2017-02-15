package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func generateHash(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}
	password := string(body)

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Cannot generate the hashed password: %s", err)
	} else {
		fmt.Fprintf(w, string(hashed))
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/", generateHash)
	log.Printf("Start HASHPASS server at port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
