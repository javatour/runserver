package worker

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/javatour/runserver/runner"
)

func (wk *Workers) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	code := new(runner.CodeFile)
	err := json.NewDecoder(r.Body).Decode(&code)
	if err != nil {
		fmt.Fprintf(w, "%s", "No understandable code")
	}
	job := Job{*code, make(chan Result)}
	wk.Do <- job
	rr := <-job.End
	result, err := rr.result, rr.err
	if err != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", html.EscapeString(err.Error()))
		return
	}
	fmt.Fprintf(w, "%s", html.EscapeString(result))
	return
}
