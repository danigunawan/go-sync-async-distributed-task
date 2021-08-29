// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"fmt"

	gocelery "github.com/danigunawan/go-sync-async-distributed-task/gocelery"
)

func get_value_from_multiple_return(val ...interface{}) []interface{} {
	return val
}

// Run Celery Worker First!
// celery -A worker worker --loglevel=debug --without-heartbeat --without-mingle
func main() {

	// // create RabbitMQ connection
	url := "amqp://guest:guest@localhost:5672/"

	// // initialize celery client and configure 1 workers
	celery_client, _ := gocelery.NewCeleryClient(
		gocelery.NewAMQPCeleryBroker(url),
		gocelery.NewAMQPCeleryBackend(url),
		1, // number of workers
	)

	// prepare arguments
	taskName := "worker_go.add"
	argA := 2
	argB := 2

	// run task
	asyncResult, err := celery_client.Delay(taskName, argA, argB)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(asyncResult)
	fmt.Println(celery_client)

	// TaskID := asyncResult.TaskID
	// fmt.Println(get_value_from_multiple_return(asyncResult.Ready())[0])

	// check if result is ready
	// isReady, _ := asyncResult.Ready()
	// fmt.Printf("Ready status: %v\n", isReady)

	// AsyncResultReady := asyncResult.Ready()
	// if AsyncResultReady != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(asyncResult.Ready()
	var id int = 1cfd33da-6523-4618-874e-cb8220d705c6
	fmt.Println(AsyncResult(id))
	
	// fmt.Println(asyncResult.TaskID)
	// fmt.Println(asyncResult.GetResult("a70e8c97-a109-40f3-8acb-49ab749e58d3"))
	// fmt.Println(cli)
	// get results from backend with timeout

	// res, err := AsyncResult.backend.GetResult(AsyncResult.TaskID)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)
	// queueName := strings.Replace("efbaf301-2387-44b3-beb6-b5033cc12cfa", "-", "", -1)
	// fmt.Println(asyncResult.GetByID("efbaf301-2387-44b3-beb6-b5033cc12cfa"))

	// get results from backend with timeout

	// run task
	// asyncResultID, err := gocelery.Delay("09afe5c6-5005-4250-829c-93bf3856b93d")
	// if err != nil {
	// fmt.Println(err)
	// }

	// fmt.Println("HASIL BY ID")
	// out := asyncResult.GetByID(asyncResult.TaskID)
	// fmt.Println(out)
	// fmt.Println(asyncResultID)

	// fmt.Println("HASIL BY TIME")

	// res, err := asyncResult.Get(10 * time.Second)
	// if err != nil {
	// 	panic(err)
	// }

	// log.Printf("result: %+v of type %+v", res, reflect.TypeOf(res))

}
