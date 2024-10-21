### `Makefile`

```Makefile
APP_NAME = averagecalculator
IMAGE_NAME = quay.io/your-username/$(APP_NAME):latest

all: build

build:
	go build -o main

run: build
	./main

test:
	go test ./... -v


Podman-build:
	docker build -t $(IMAGE_NAME) .

Podman-run:
	podman run -p 8080:9901 $(IMAGE_NAME)

# Push podman --> quay.io
Quay-push:
	podman push $(IMAGE_NAME)-t ${TAG}

# Deploy to OpenShift  
openshift-deploy:
	oc new-app . --strategy=docker --name=$(APP_NAME)
	oc start-build $(APP_NAME) --from-dir=.

openshift-clean:
	oc delete all --selector app=$(APP_NAME)
