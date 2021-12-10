startDb:
	/bin/bash script/runDb.sh
testing:
	cd test/ && go test .
stopDb:
	cd  deploy/local && sudo docker-compose down

