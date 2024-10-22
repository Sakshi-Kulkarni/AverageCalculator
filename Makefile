APP_NAME = averagecalculator
IMAGE_NAME = quay.io/rh-ee-sakkulka/$(APP_NAME):latest

all: build

build:
	go build -o main

run: build
	./main

test:
	go test ./... -v

podman-build:
	podman build -t $(IMAGE_NAME) .

podman-run:
	podman run -p 8080:9901 $(IMAGE_NAME)

# Push podman --> quay.io
quay-push:
    podman login quay.io .
	podman push $(IMAGE_NAME)

# Deploy to OpenShift  
openshift-deploy:
	oc new-app . --strategy=docker --name=$(APP_NAME)

openshift-route:
	oc expose service $(APP_NAME)
	oc get routes.route.openshift.io | awk '{ print $2 }'

openshift-clean:
	oc delete all --selector app=$(APP_NAME)

help:
	@echo "Makefile Help:"
	@echo "  build             Build the Go application"
	@echo "  run               Build and run the Go application"
	@echo "  test              Run all test cases"
	@echo "  podman-build      Build the Docker image using Podman"
	@echo "  podman-run        Run the Docker container using Podman"
	@echo "  quay-push         Push the image to Quay.io"
	@echo "  openshift-deploy  Deploy the app to OpenShift using Docker strategy"
	@echo "  openshift-route   Expose the service and display the route"
	@echo "  openshift-clean   Remove all OpenShift resources for the app"
