# HOWTO start

- Install go-swagger
- swagger generate server -f swagger/worker.yaml -A worker -a worker
- go build ./cmd/worker-server
- ./worker-server --port 8080