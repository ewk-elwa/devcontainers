package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Name string
}

type TaskQueue struct {
	mu    sync.Mutex
	tasks []Task
}

func (q *TaskQueue) AddTask(task Task) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.tasks = append(q.tasks, task)
}

func (q *TaskQueue) GetTask() *Task {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.tasks) == 0 {
		return nil
	}
	task := q.tasks[0]
	q.tasks = q.tasks[1:]
	return &task
}

func processTask(task *Task) {
	fmt.Printf("Processing task ID: %d, Name: %s\n", task.ID, task.Name)
	time.Sleep(1 * time.Second) // Simulate task processing
}

func worker(queue *TaskQueue, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task := queue.GetTask()
		if task == nil {
			break
		}
		processTask(task)
	}
}

func main() {
	queue := &TaskQueue{}
	for i := 1; i <= 10; i++ {
		queue.AddTask(Task{ID: i, Name: fmt.Sprintf("Task-%d", i)})
	}

	var wg sync.WaitGroup
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(queue, &wg)
	}
	wg.Wait()
	fmt.Println("All tasks processed")
}
