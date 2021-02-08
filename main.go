package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/core/runner"
)

func javaHandler(w http.ResponseWriter, r *http.Request) {
	code := runner.JavaExecutor{}
	err := json.NewDecoder(r.Body).Decode(&code)
	if err != nil {
		fmt.Println("err!")
	}
	result := code.JavaRunner()
	fmt.Fprintf(w, "%s", html.EscapeString(result))
}

func pythonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dev..")
}

func main() {
	http.HandleFunc("/java", javaHandler)
	http.HandleFunc("/python", pythonHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
