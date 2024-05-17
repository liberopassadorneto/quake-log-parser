# Quake Parser

This project is a CLI application to process Quake game log files and generate reports in JSON format. It uses the Cobra library for the CLI and Docker for building and running the environment.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Project Structure
- **cmd/**
    - `root.go`: Defines the root command for the CLI application.
    - `upload.go`: Defines the 'upload' command for processing log files.
- **handler/**
    - `handler.go`: Contains the `Handler` struct and methods for processing log files.
- **logger/**
    - `logger.go`: Configures the global logger using logrus.
- **models/**
    - `models.go`: Defines the data models for `Game` and `Player`.
    - `models_test.go`: Unit tests for `models.go`.
- **parser/**
    - `parser.go`: Contains the `Parser` struct and methods for parsing log files.
    - `parser_test.go`: Unit tests for `parser.go`.
- `main.go`: The entry point for the CLI application.
- **report/**: Directory to store output files (created manually).
- **logs/**: Directory containing the log files.
    - `qgames.log`: Example log file to be processed.
- `Dockerfile`: Dockerfile to build the Docker image.
- `docker-compose.yml`: Docker Compose file to build and run the Docker container.

## Setup

1. Clone the repository to your local environment:
    ```bash
    git clone https://github.com/liberopassadorneto/quake-parser.git
    cd quake-parser
    ```

2. Run the following command to build the Docker image:
    ```bash
    docker-compose build
    ```

## Running the Docker Container

1. Run the Docker container, mounting the necessary volumes to provide the log file and store the output file:
    ```bash
    docker-compose run quake upload --file /logs/qgames.log --output output.json
    ```

   In the command above:
    - `upload --file /logs/qgames.log --output output.json` runs the upload command to process the log file and save the report in the mounted `report` directory.

## Mounted Directories

- `./logs`: This directory on the host is mounted as `/logs` in the container and should contain the log file to be processed.
- `./report`: This directory on the host is mounted as `/report` in the container and will store the output JSON file.

## Complete Example (with Docker)

Here are the complete steps to set up, build, and run the project:

1. Clone the repository:
    ```bash
    git clone https://github.com/liberopassadorneto/quake-parser.git
    cd quake-parser
    ```

2. Create the directories and place the log file in the `logs` directory (if necessary):
    ```bash
    mkdir report
    mkdir logs
    # Place your example.log file in ./logs
    ```

3. Build the Docker image:
    ```bash
    docker-compose build
    ```

4. Run the Docker container:
    ```bash
    docker-compose run quake upload --file /logs/qgames.log --output output.json
    ```

   The resulting JSON will be saved in `./report/output.json`.

## Build Locally and Run
To build and run the project without using Docker, follow the instructions below:

### Build the Project
To build the project, execute the following command:
```bash
go build -o quake .
```
### Run the Project
To run the quake binary, use the command below:
```bash
./quake upload --file logs/qgames.log --output output.json
```

## Testing
To run all the unit tests in the project, navigate to the root directory and execute the following command:
```bash
go test -v ./...
```