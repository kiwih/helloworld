/* Copyright May Contain Awesome Ltd, 2015

This basic program is used in Lesson 3 of Project 1: A social network

It aims to show the use of templates and the web server. Students will extend this programs functionality in class.

Released under MIT license. */

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//templates global variable
var templates = template.Must(template.ParseFiles("template.html"))

//website handler function
func handler(w http.ResponseWriter, r *http.Request) {
	err := templates.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func main() {
	//set up the listening address
	serverAddr := ":8080"
	fmt.Println("Hello! Server listening at", serverAddr)

	//assign routes
	http.HandleFunc("/", handler)

	//run server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Fatal error:", err.Error())
	}
}
