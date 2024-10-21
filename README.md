![Gopher Logo](https://media.licdn.com/dms/image/D5612AQFXS5VRN7cQvg/article-cover_image-shrink_720_1280/0/1713275727837?e=2147483647&v=beta&t=s9EZeH9HptwvrbUMMGzXl59r25UteIfOn--p8C_0c24)


# 🚀 Average Calculator
![Go Version](https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Podman](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Red Hat](https://img.shields.io/badge/Red%20Hat-Compatible-EE0000?style=for-the-badge&logo=red-hat&logoColor=white)
[![Unit Tests](https://img.shields.io/badge/Unit%20Tests-Passing-4CAF50?style=for-the-badge&logo=checkmarx&logoColor=white)](https://github.com/yourusername/even-number-average/actions)
![MIT License](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)


> A simple, Go-based web service for calculating averages from a list of given numbers. Choose to calculate averages for all numbers, even numbers, or odd numbers.

---

## 📚 Table of Contents

- [Introduction](#-introduction)
- [Features](#-features)
- [Getting Started](#-getting-started)
- [API Usage](#-api-usage)
- [Development](#-development)
- [Testing](#-testing)
- [Podman Usage](#-docker-usage)
- [Deployment on OpenShift](#-deployment-on-openshift)
- [Contributing](#-contributing)
- [License](#-license)

---

## 💡 Introduction

This web service allows users to calculate the average of numbers by providing different query parameters. It's built with **Go** and provides RESTful endpoints for easy interaction. The application is designed to handle errors and ensures valid input from users.

### How It Works

1. **Input**: Users send a `POST` request with a space-separated list of numbers.
2. **Processing**: The server calculates the average based on the specified type (`n=1`, `n=2`, or `n=3`).
3. **Output**: Returns the calculated average or an error message in case of invalid input.

---

## ✨ Features

- **Multiple Calculation Modes**: Choose between averages for all numbers, even numbers, or odd numbers.
- **Error Handling**: Validates inputs and provides meaningful error messages.
- **RESTful API**: Easy-to-use `POST` endpoints.
- **Docker Support**: Run the app in a container with ease.
- **OpenShift Integration**: Deploy directly to OpenShift with custom configurations.

---

## 🚀 Getting Started

### Prerequisites

Make sure you have the following installed:

- [Go 1.23+](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [OpenShift CLI (oc)](https://docs.openshift.com)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/AverageCalculator.git
   cd AverageCalculator
2. **Build the application**:

   ```bash
   go build -o main
3. **Run the application locally**:
   ```bash
   ./main
4. **Run tests**:
   ```bash
   make test

## 🔥 API Usage

### Endpoint
 ```bash
 POST /average?n={1|2|3}
```
### Parameters
  ```bash
  n=1: Calculate the average of even numbers.
  n=2: Calculate the average of odd numbers.
  n=3: Calculate the average of all numbers.
```
### Example Request
```bash
curl -X POST -d "10 20 30 40" "http://localhost:9901/average?n=3"
```
### Example Response
```bash
The average is 25.00
```
### Error Handling
If the input is invalid or missing, the API will respond with a 400 Bad Request and an error message like:

 ```bash
no numbers provided
```
## 🛠️ Development
Directory Structure
``` bash
AverageCalculator/
├── handler/          # Contains the main HTTP handler logic
├── tests/            # Unit and integration tests
├── Dockerfile        # Dockerfile for containerization
├── Makefile          # Makefile for common tasks
└── README.md         # Project documentation
```

## ✅ Testing

Testing is a crucial part of ensuring the reliability and correctness of this project. The following sections cover how to run tests, what tests are included, and what to expect.

### Running Tests

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

To ensure the codebase is fully tested, you can generate a coverage report:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

This will produce an HTML file that shows which lines of code are covered by tests.

### Included Tests

The testing suite includes:

- **Unit Tests**: Focused on testing individual functions and handlers in isolation.
- **Integration Tests**: Testing the `AverageHandler` with different inputs to ensure end-to-end functionality.
- **Error Handling Tests**: Validates that the API returns appropriate error messages and status codes for invalid inputs (e.g., missing or malformed data).


## 🐳 Podman Usage
**Build Podman Image**
```bash
make docker-build
```
Run Podman Container
```bash
make docker-run
```
The app will be available at:http://localhost:9901.

### 🚢 Deployment on OpenShift
**Deploying the Application**
```bash
oc new-app . --name avg-s2i --strategy=docker
```

```bash
oc new-app . --strategy=docker --name=average
```
Start the build:

```bash
oc start-build average --from-dir=.
```
Access the application via the OpenShift route created.

## 🤝 Contributing
We welcome contributions to make this project even better! Please fork the repository and create a pull request with a meaningful commit message.

## 📄 License
This project is licensed under the MIT License. See the #-LICENSE file for more information.