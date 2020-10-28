REPOSITORY?=
TAG?=

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o openstackweekly cmd/openstackweekly/*.go 
	docker build -t ${REPOSITORY}:${TAG} .
	rm openstackweekly

