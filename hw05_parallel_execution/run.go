package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errCounter int32

	tasksCh := make(chan Task, len(tasks))
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(errCounter *int32) {
			for {
				task, ok := <-tasksCh
				mu.Lock()
				if !ok || *errCounter >= int32(m) {
					mu.Unlock()
					break
				}
				mu.Unlock()
				if err := task(); err != nil {
					mu.Lock()
					*errCounter++
					mu.Unlock()
				}
			}
			wg.Done()
		}(&errCounter)
	}

	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)
	wg.Wait()
	if m != 0 && errCounter >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
