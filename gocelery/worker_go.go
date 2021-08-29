// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"time"

	"github.com/danigunawan/go-sync-async-distributed-task/gocelery"
)

// Celery Task
func add(a int, b int) int {
	return a + b
}

func main() {

	// create AMQP & Redis connection pool
	url_amqp := "amqp://guest:guest@localhost:5672/"
	// url_redis :=

	// create broker and backend
	// celeryBroker := gocelery.NewRedisCeleryBroker("redis://localhost:6379")
	// celeryBackend := gocelery.NewRedisCeleryBackend("redis://localhost:6379")

	//use AMQP instead
	celeryBroker := gocelery.NewAMQPCeleryBroker(url_amqp)
	celeryBackend := gocelery.NewAMQPCeleryBackend(url_amqp)

	// initialize celery client
	cli, _ := gocelery.NewCeleryClient(celeryBroker, celeryBackend, 5) // number of workers

	// register task
	cli.Register("worker_go.add", add)

	// start workers (non-blocking call)
	cli.StartWorker()

	// wait for client request
	time.Sleep(100000 * time.Second)

	// stop workers gracefully (blocking call)
	cli.StopWorker()
}
