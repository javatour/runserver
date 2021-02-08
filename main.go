package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/javatour/runserver/runner"
)

type Result struct {
	result string
	err    error
}

var (
	m   sync.Mutex
	num int
)

// @dev 함수를 나누거나 파일을 나눌 예정
// context를 사용하여, 무한루프도는 고루틴 강제 종료 예정
// 예외 처리를 조금 더 깔끔히 변경 예정
func handler(w http.ResponseWriter, r *http.Request) {
	code := new(runner.CodeFile)
	err := json.NewDecoder(r.Body).Decode(&code)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan Result)
	go func(code runner.CodeFile) {
		m.Lock()
		num++
		m.Unlock()
		result, err := code.Run(strconv.Itoa(num))
		ch <- Result{result, err}
	}(*code)
	select {
	case t := <-ch:
		result, err := t.result, t.err
		fmt.Println(result)
		if err != nil {
			log.Fatal(err)
			fmt.Fprintf(w, "%s", html.EscapeString(err.Error()))
		}
		fmt.Fprintf(w, "%s", html.EscapeString(result))
	case <-time.After(8 * time.Second):
		fmt.Println("time out")
		fmt.Fprintf(w, "%s", html.EscapeString("Time out"))
	}
}

// @dev
// 큐 구조로 변경할 예정
// ListenAndServe가 고루틴으로 구현되어 있으므로,
// 로직은 각 핸들러는 큐에 작업을 넣고
// worker가 그 큐에서 작업을 빼서 처리하여 예정 응답하는 구조가 더 효과적일 수 있음
// 현재 구조는 요청이 급증하게 되면 고루틴의 수가 많아지므로
// 해당 처리가 늦게 되어 타임아웃의 가능성이 존재

func main() {
	http.HandleFunc("/code", handler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
