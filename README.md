![Average calculator](https://media.licdn.com/dms/image/D5612AQFXS5VRN7cQvg/article-cover_image-shrink_720_1280/0/1713275727837?e=2147483647&v=beta&t=s9EZeH9HptwvrbUMMGzXl59r25UteIfOn--p8C_0c24)


# Average Calculator   
![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![Go Report Card](https://goreportcard.com/badge/github.com/sakshi-kulkarni/AverageCalculator)
![Go Version](https://img.shields.io/github/go-mod/go-version/sakshi-kulkarni/AverageCalculator?color=green&logo=go)
![OpenShift Version](https://img.shields.io/badge/OpenShift-v4.10+-green?logo=redhatopenshift)
![License: MIT](https://img.shields.io/badge/License-MIT-green.svg?style=flat)


> A simple, Go-based web service for calculating averages from a list of given numbers. Choose to calculate averages for all numbers, even numbers, or odd numbers.
---

## Introduction

This web service allows users to calculate the average of numbers by providing different query parameters. It's built with **Go** and provides RESTful endpoints for easy interaction. The application is designed to handle errors and ensures valid input from users.

### How It Works

1. **Input**: Users send a `POST` request with a space-separated list of numbers.
2. **Processing**: The server calculates the average based on the specified type (`n=1`, `n=2`, or `n=3`).
3. **Output**: Returns the calculated average or an error message in case of invalid input.

---
## 🚀 Getting Started

### Prerequisites

Make sure you have the following installed:

- [Go 1.21+](https://golang.org/doc/install)
- [Podman](https://podman.io/docs/installation)
- [OpenShift CLI (oc)](https://docs.openshift.com)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/sakshi-kulkarni/AverageCalculator.git
   cd AverageCalculator
   ```
2. **Build the application**:
   ```bash
   go build -o main
   ```
3. **Run the application locally**:
   ```bash
   ./main
   ```
---
### API Usage

- **Endpoint**
 ```bash
 POST http://localhost:9901/average?n={1|2|3}
```
- **Parameters**
  ```bash
  n=1: Calculate the average of even numbers.
  n=2: Calculate the average of odd numbers.
  n=3: Calculate the average of all numbers.
  ```
- **Example Request**
```bash
curl -X POST -d "10 20 30 40" "http://localhost:9901/average?n=3"
```
- **Example Response**
```bash
The average is 25.00
```
- **Error Handling**
  
If the input is invalid or missing, the API will respond with a 400 Bad Request and an error message like:

 ```bash
no numbers provided
```
---
### Testing

The following sections cover how to run tests, what tests are included, and what to expect.

**NOTE:** To run the tests, navigate to the specific folders containing the test files and execute the following commands.

**Running Tests**

1. **Run all tests** using the Makefile:

   ```bash
   make test
   ```

   This command runs all unit tests and integration tests defined for the application.

2. **Run tests manually**:

   You can also use the `go` testing command directly:

   ```bash
   go test ./...
   ```

3. **Benchmark Tests**:

   To run benchmarks for the `AverageHandler`, use:

   ```bash
   go test -bench=.
   ```

   This is useful to understand the performance of the application, especially when processing large datasets.

### Test Coverage

- To ensure the codebase is fully tested, you can generate a coverage report:

```bash
go test -coverprofile=coverage.out 
go tool cover -html=coverage.out
```

This will produce an HTML file that shows which lines of code are covered by tests.

---
### Included Tests

The testing suite includes:

- **Unit Tests**: Focused on testing individual functions and handlers in isolation.Testing the `AverageHandler` with different inputs to ensure end-to-end functionality.
- **Error Handling Tests**: Validates that the API returns appropriate error messages and status codes for invalid inputs (e.g., missing or malformed data).


### Podman Usage
**Build Image**
```bash
make podman-build
```
**Run Container**
```bash
make podman-run
```
The app will be available at:http://localhost:9901.

---
## Deployment on OpenShift
- **Deploying the Application**
```bash
oc new-app . --strategy=docker --name=average
```
Access the application via the OpenShift route created.
To get the route, refer to the openshift-route target in the Makefile. 

- Simply run:

```bash
make openshift-route
```

---
## 📄 License
This project is licensed under the MIT License. See the LICENSE file for more information.