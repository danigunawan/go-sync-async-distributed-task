// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/danigunawan/go-sync-async-distributed-task/gocelery"
)

func get_value_from_multiple_return(val ...interface{}) []interface{} {
	return val
}

// Run Celery Worker First!
// celery -A worker_py worker --without-heartbeat --without-gossip --without-mingle --loglevel=INFO
func main() {

	// create RabbitMQ connection
	url := "amqp://guest:guest@localhost:5672/"

	// initialize celery client and configure 1 workers
	cli, _ := gocelery.NewCeleryClient(
		gocelery.NewAMQPCeleryBroker(url),
		gocelery.NewAMQPCeleryBackend(url),
		1, // number of workers
	)

	// prepare arguments
	taskName := "worker_py.add"
	argA := 5456
	argB := 2878

	// run task
	asyncResult, err := cli.Delay(taskName, argA, argB)
	if err != nil {
		fmt.Println(err)
	}
	// TaskID := asyncResult.TaskID
	// fmt.Println(get_value_from_multiple_return(asyncResult.Ready())[0])

	// check if result is ready
	// isReady, _ := asyncResult.Ready()
	// fmt.Printf("Ready status: %v\n", isReady)

	// AsyncResultReady := asyncResult.Ready()
	// if AsyncResultReady != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(asyncResult.Ready())
	fmt.Println(asyncResult.TaskID)
	// fmt.Println(asyncResult.AsyncGet())
	// fmt.Println(asyncResult.GetResult("a70e8c97-a109-40f3-8acb-49ab749e58d3"))
	// fmt.Println(cli)
	// get results from backend with timeout

	// get results from backend with timeout
	res, err := asyncResult.Get(10 * time.Second)
	if err != nil {
		panic(err)
	}

	log.Printf("result: %+v of type %+v", res, reflect.TypeOf(res))

}
