package datastruct

import (
	"container/list"
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	taskList := NewTaskList()
	pool := NewGoroutinePool(1000)
	for _, task := range taskList {
		pool.Register(task)
	}
	pool.Go()

	end := time.Now().UnixNano()
	//
	//for _, task := range taskList {
	//	fmt.Println("index:", task.source, ":", task.result)
	//}
	fmt.Print("time: ", end-start)
}

func testDefer() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		fmt.Println("a")
		panic("error")
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(3)
	fmt.Println(3)
	fmt.Println(3)
	fmt.Println(3)
	fmt.Println(3)
}

func testList() {
	queue := list.New()
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack("3sd")

	fmt.Println(queue.Remove(queue.Front()).(int))
	fmt.Println(queue.Remove(queue.Front()).(int))
	fmt.Println(queue.Remove(queue.Front()))

	fmt.Println("end")
}
