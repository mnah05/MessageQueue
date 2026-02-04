run:
	go run cmd/workers/main.go & \
	go run cmd/api/main.go

test:
	chmod +x ./req.sh & \
	./req.sh