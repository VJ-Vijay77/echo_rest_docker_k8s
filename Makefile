

up: 
	@echo "starting postgres container"
	docker compose down
	docker compose up -d
	@echo "Started"
	

down:
	@echo "stopping container"
	docker compose down
	
	
server:
	go build -o main .
	./main

stop:
	-pkill -SIGTERM -f "./main"
	rm main

