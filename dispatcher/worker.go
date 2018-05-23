package dispatcher

import (
	"crypto/rand"
	"encoding/base64"
	"fedomn/converter/processor"
	"fmt"
	"strings"
)

type (
	JobType int

	JobInput struct {
		Type    JobType
		Context string
	}
	JobInputQueue chan JobInput

	JobOutput struct {
		Input  string
		Output string
		Type   JobType
	}
	JobOutputQueue chan JobOutput

	Worker struct {
		id             string
		processor      processor.Processor
		quit           chan bool
		jobInputQueue  JobInputQueue
		jobOutputQueue JobOutputQueue
		jobType        JobType
	}
)

func genId() string {
	b := make([]byte, 8)
	rand.Read(b)
	s := base64.URLEncoding.EncodeToString(b)
	return strings.Trim(s, " ")
}

func newWorker(jobType JobType, p processor.Processor) Worker {
	return Worker{
		id:             genId(),
		processor:      p,
		jobInputQueue:  make(JobInputQueue, 10),
		jobOutputQueue: make(JobOutputQueue, 10),
		jobType:        jobType,
		quit:           make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.jobInputQueue:
				fmt.Printf("\033[32mWorker Log:\033[0m %v got job %+v\n", w.id, job)
				output := w.processor.Process(job.Context)
				w.jobOutputQueue <- JobOutput{job.Context, output, w.jobType}
			case <-w.quit:
				fmt.Printf("worker: %+v stop ok!\n", w.jobType)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
		close(w.jobInputQueue)
		close(w.jobOutputQueue)
	}()
}

// 保证任务队列顺序 阻塞进入
func (w Worker) AddJob(job JobInput) {
	w.jobInputQueue <- job
}

// 保证 jobOutputQueue <- output 不阻塞 在goroutines遍历结果
func (w Worker) AcquireJobOutput() JobOutputQueue {
	return w.jobOutputQueue
}
