package main

import (
	"lendr/internal/routes"
	"lendr/pkg/db"
	"log"
	"net/http"
	"os"
	"unicode"
)

// StrEmpty checks whether string contains only whitespace or not
func StrEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	r := []rune(s)
	l := len(r)

	for l > 0 {
		l--
		if !unicode.IsSpace(r[l]) {
			return false
		}
	}

	return true
}


// *************************
// Get Database Env Variables
// *************************
var username = os.Getenv("MYSQL_USER")
var password = os.Getenv("MYSQL_PASSWORD")
var hostname = os.Getenv("MYSQL_HOST")
var database = os.Getenv("MYSQL_DATABASE")

func init() {
	// *************************
	// Verify All Env Variables Found
	// *************************
	if StrEmpty(username) { panic("missing env variable MYSQL_USER") }
	if StrEmpty(password) { panic("missing env variable MYSQL_PASSWORD") }
	if StrEmpty(hostname) { panic("missing env variable MYSQL_HOST") }
	if StrEmpty(database) { panic("missing env variable MYSQL_DATABASE") }
}

func main() {
	// *************************
	// Setup DB
	// *************************
	if err := db.Open(username, password, hostname, database); err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	// *************************
	// Initialize Web Server
	// *************************
	web := routing.Register()
	if err := http.ListenAndServe(":8080", web); err != nil {
		log.Fatal(err)
	}
}
