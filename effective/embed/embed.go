package main

import (
	"fmt"
	"log"
	"os"
)

type Job struct {
	Command string
	*log.Logger
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

// 可以覆盖原方法
func (job *Job) Printf(format string, args ...interface{}) {
	// 如果属性没有赋值，可以通过类型名来调用(job.Logger)
	// 注意最后的...
	job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}

var loggerStderr = log.New(os.Stderr, "Job: ", log.Ldate)

func main() {
	job := NewJob("listen", loggerStderr)
	// 直接调用Logger的方法
	job.Println("starting now...")

	job.Printf("%s %d", "counter", 2022)
}
