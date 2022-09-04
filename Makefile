

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

migrateup:
	migrate -path database/migration -database "postgresql://vijay:zmxmcmvbn@localhost:5432/echo_rest?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://vijay:zmxmcmvbn@localhost:5432/echo_rest?sslmode=disable" -verbose down
		