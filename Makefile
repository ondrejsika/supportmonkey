build:
	go build

docker-build:
	docker build -t ondrejsika/supportmonkey .

docker-push:
	docker push ondrejsika/supportmonkey

docker: docker-build docker-push

helm-package:
	helm package ./helm/supportmonkey
