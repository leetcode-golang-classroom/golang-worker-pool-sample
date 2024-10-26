package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/leetcode-golang-classroom/golang-worker-pool-sample/internal/work"
)

func main() {
	wp, err := work.NewPool(5, 5)
	if err != nil {
		log.Fatal(err)
	}
	// 設定離開時，先送　context cancel 訊號給 subroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wp.Start(ctx)
	// 加入 20　個　Task
	for range 20 {
		task := work.NewTask(func() error {
			const urlString = "https://google.com"
			res, err := http.Get(urlString)
			if err != nil {
				return err
			}
			fmt.Printf("%s returned status code %d\n", urlString, res.StatusCode)
			return nil
		}, func(err error) {
			fmt.Println(err)
		})

		// wp.AddTask(task)
		wp.AddTaskNonBlocking(task)
	}
	// 利用　counter 來計算有多少 task 完成
	counter := 0
	for completed := range wp.TasksCompleted() {
		if completed {
			counter++
		}
		// 當 Task 數量等於建立數量，代表所有任務都完成
		if counter == 20 {
			wp.Stop()
			return
		}
	}
}
