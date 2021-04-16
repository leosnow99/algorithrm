package main

import "sync"

type Task interface {
	Process()
}

type GoroutinePool struct {
	num        int
	routineNum int
	index      int
	mutex      sync.Mutex
	tasks      []Task
}

func NewGoroutinePool(num int) *GoroutinePool {
	if num <= 0 || num >= 10000 {
		panic("go routine num should larger than 0 and less than 1000")
	}
	return &GoroutinePool{
		num:        num,
		mutex:      sync.Mutex{},
		routineNum: 1000,
	}
}

func (g *GoroutinePool) SetGoroutineNum(num int) {
	if num <= 0 || num > 500 {
		panic("error goroutine num")
	}
	g.routineNum = num
}

func (g *GoroutinePool) Register(task Task) {
	g.mutex.Lock()
	g.tasks = append(g.tasks, task)
	g.mutex.Unlock()
}

//func (g *GoroutinePool) Go() {
//	if len(g.tasks) == 0 {
//		return
//	}
//
//	wg := &sync.WaitGroup{}
//	count := len(g.tasks)
//
//	if count > g.num {
//		count = g.num
//	}
//
//	for i := 0; i < count; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//
//			for {
//				g.mutex.Lock()
//				if g.index == len(g.tasks) {
//					g.mutex.Unlock()
//					return
//				}
//				task := g.tasks[g.index]
//				g.index++
//				g.mutex.Unlock()
//				task.Process()
//			}
//		}()
//	}
//
//	wg.Wait()
//	g.index = 0
//	g.tasks = g.tasks[:0]
//}

func (g *GoroutinePool) Go() {
	if len(g.tasks) == 0 {
		return
	}

	wg := &sync.WaitGroup{}
	count := len(g.tasks)

	for i := 0; i < count; i += g.routineNum {
		taskNum := count
		if taskNum > g.routineNum {
			taskNum = g.routineNum
		}

		for j := 0; j < taskNum; j++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				g.mutex.Lock()
				task := g.tasks[g.index]
				g.index++
				g.mutex.Unlock()
				task.Process()
			}()

		}

		wg.Wait()
	}

	g.index = 0
	g.tasks = g.tasks[:0]
}
