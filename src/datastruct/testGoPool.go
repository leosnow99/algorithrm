package main

import "time"

type MyTask struct {
	source int
	result int
}

func (t *MyTask) Process() {
	t.result = t.source + 1
	time.Sleep(time.Millisecond * 2)
}

func NewTaskList() []*MyTask {
	result := make([]*MyTask, 10000)
	for i := 0; i < 10000; i++ {
		result[i] = &MyTask{
			source: i,
		}
	}

	return result
}
