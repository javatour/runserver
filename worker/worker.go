package worker

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/javatour/runserver/runner"
)

var (
	Maxtime   = 10
	MaxWorker = runtime.NumCPU()
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

func (w Worker) start() {
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

// WorkStart() 는 구체적인 worker들을 생성하고, 일을 할 수 있는 준비상태로 변환하고
// 쉬고있는 [Job 채널]을 전송하는 채널을 통해 worker를 관리합니다.
func (w Workers) WorkStart() {
	WorkerChannel := make(chan chan Job)
	fmt.Println("You hire " + strconv.Itoa(MaxWorker) + " employees")
	for i := 0; i < MaxWorker; i++ {
		worker := Worker{i, WorkerChannel, make(chan Job)}
		worker.start()
	}
	go func() {
		for {
			work := <-w.Do
			worker := <-WorkerChannel
			worker <- work
		}
	}()
}

// MakeWorkers() 는 일을 공급하는 채널을 초기화 합니다.
func MakeWorkers() (Workers, error) {
	employeeChannel := make(chan Job)
	return Workers{employeeChannel}, nil
}

func (wk *Workers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wk.Handle(w, r)
}
