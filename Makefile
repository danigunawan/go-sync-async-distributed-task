run-worker:
	go run worker.go
run-client:
	go run client.go
run-celery-py:
	celery -A worker_py worker --without-heartbeat --without-gossip --without-mingle --loglevel=INFO
run-celery-go:
	celery -A worker_go worker --without-heartbeat --without-gossip --without-mingle --loglevel=INFO
# cd example 
#pip install celery
# pip install celery==4.4.7
#go run worker.go
#go run worker.go&
#go run client.go
#go run client.go&
#celery -A worker worker
#celery -A worker worker&

#buat repo : github.com/danigunawan/go-sync-async-distributed-task
#go mod init github.com/danigunawan/go-sync-async-distributed-task
#git init
#git add .
#git commit -m "update module go"
#git remote add origin git@github.com:danigunawan/go-sync-async-distributed-task