package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"time"
	"zhichen/ollama_client"
)

var query string

func main() {
	fmt.Println("Server up!")
	homePage := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Index")
		http.ServeFile(w, r, "./templates/index.html")
	}

	chat := func(w http.ResponseWriter, r *http.Request) {
		query = r.FormValue("query")
		b, err := os.ReadFile("./templates/user_message.html")
		if err != nil {
			fmt.Print(err)
		}
	
		template := string(b)
		fmt.Fprintf(w, template, query)
	}

	generate := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second) 
		b, err := os.ReadFile("./templates/bot_message.html")
		if err != nil {
			fmt.Print(err)
		}
		template := string(b)

		response := ollama_client.Generate(query)

		fmt.Fprintf(w, template, response)
	}

	http.HandleFunc("GET /", homePage)
	http.HandleFunc("POST /chat", chat)
	http.HandleFunc("GET /generate", generate)
	http.Handle("GET /img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}