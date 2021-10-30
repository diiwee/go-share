package main

import (
	"fmt"
	"sync"
	"time"
)

// Single 线程安全的单例模式
type Single struct {
}

var (
	singleInstance *Single
	once           sync.Once
)

func getInstance() *Single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &Single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}

	if singleInstance == nil {
		fmt.Println("Single error.")
	}

	return singleInstance
}

func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}
	time.Sleep(time.Second * 10)
}
