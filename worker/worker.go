package worker

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/javatour/runserver/runner"
)

const (
	Maxtime   = 10
	MaxWorker = 4
)

type Job struct {
	Code runner.CodeFile
	End  chan Result
}

// Worker is the routine who work
type Worker struct {
	ID            int
	WorkerChannel chan chan Job
	Channel       chan Job
}

// Workers 는 Worker들에게 일을 부여하는
// 채널로 구성됩니다.
// End는 한 일을 돌려줘야 하는 채널이 정해져 있습니다.

type Workers struct {
	Do chan Job
}

type Result struct {
	result string
	err    error
}

func (w Worker) Start() {
	var n = 0
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			job := <-w.Channel
			fmt.Println("[ Worker", w.ID, "] 가 일을 수주해 갔습니다!")
			fmt.Println("[ Worker", w.ID, "] 현재 워커가 메모리 누수를 일으키는 숫자는", n, "입니다.")
			id := w.ID
			path := strconv.Itoa(id)
			result, err := job.Code.Run(path)
			if err != nil {
				fmt.Println("[ Worker", w.ID, "] 가 받은 일이 에러일 가능성이 높습니다. 종료!.")
				job.End <- Result{"", errors.New("에러!!")}
			} else {
				fmt.Println("[ Worker", w.ID, "] 가 일을 마쳤습니다.")
				job.End <- Result{result, nil}
			}
		}
	}()
}

func (w Workers) WorkStart() {
	WorkerChannel := make(chan chan Job)
	for i := 0; i < MaxWorker; i++ {
		worker := Worker{i, WorkerChannel, make(chan Job)}
		worker.Start()
	}
	go func() {
		for {
			work := <-w.Do
			worker := <-WorkerChannel
			worker <- work
		}
	}()
}

func MakeWorkers() (Workers, error) {
	dochannel := make(chan Job)
	return Workers{dochannel}, nil
}

func (wk *Workers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wk.Handle(w, r)
}
