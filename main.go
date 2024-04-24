package main

import (
	"fmt"
	server "groupie-tracker/back"
	"log"
	"net/http"
)

func main() {
	fmt.Println("http://localhost:4747/")
	fs := http.FileServer(http.Dir("./front/css_img"))
	http.Handle("/css_img/", http.StripPrefix("/css_img", fs))
	http.HandleFunc("/", server.MainPage)
	http.HandleFunc("/artists/", server.InfoAboutArtist)
	err := http.ListenAndServe(":4747", nil)
	if err != nil {
		log.Fatal(err)
	}
}
