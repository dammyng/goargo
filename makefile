build:
	docker build -t goargoserver .
run:
	docker run -p 8080:8080 -e ENVIRONMENT=development goargoserver
login: 
	docker login -u=cabhub -p=OyindamoLA1.
tag:
	docker tag goargoserver cabhub/goargoserver:latest
push: 
	docker push cabhub/goargoserver:latest

