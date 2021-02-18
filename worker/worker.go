package worker

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

// Wokers 는 Worker들에게 일을 부여하는
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
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			fmt.Println("Worker", w.ID, "takes job!")
			job := <-w.Channel
			id := w.ID
			path := strconv.Itoa(id)
			tempChanel := make(chan Result)
			go func() {
				result, err := job.Code.Run(path)
				tempChanel <- Result{result, err}
			}()
			select {
			case success := <-tempChanel:
				fmt.Println("Worker", w.ID, "finished job!")
				fmt.Println(success.result)
				job.End <- success
			case <-time.After(10 * time.Second):
				fmt.Println("누군가.. 이걸로 비트코인 채굴이라도 하나보다.")
				job.End <- Result{"", errors.New("무한루프 금지!!")}

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
