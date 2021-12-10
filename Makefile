startDb:
	/bin/bash script/runDb.sh
testing:
	cd test/ && go test .
stopDb:
	cd  deploy/local && sudo docker-compose down

run:
	/bin/bash script/runDb.sh
	sleep 3;
	go run main.go

