package dispatcher

import (
	"fedomn/converter/processor"
	"fmt"
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
		processor      processor.Processor
		quit           chan bool
		jobInputQueue  JobInputQueue
		jobOutputQueue JobOutputQueue
		jobType        JobType
	}
)

func newWorker(p processor.Processor) Worker {
	return Worker{
		processor:      p,
		jobInputQueue:  make(JobInputQueue),
		jobOutputQueue: make(JobOutputQueue),
		quit:           make(chan bool),
	}
}

func (w Worker) Start() {
	fmt.Println("worker start")
	go func() {
		for {
			select {
			case job := <-w.jobInputQueue:
				output := w.processor.Process(job.Context)
				if output != "" {
					w.jobOutputQueue <- JobOutput{job.Context, output, w.jobType}
				}
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
