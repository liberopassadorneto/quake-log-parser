# STEP 1: Build executable binary
FROM golang:1.22.1-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum from host to container
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Create the /tmp directory to store the games.log while running
RUN mkdir -p /tmp

# Copy everything from the current directory of the host to the Working Directory in the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o quake-parser .

# STEP 2: Build a small image
FROM scratch

# Copy the executable from the builder stage
COPY --from=builder /app/quake-parser /quake-parser

# Copy the report directory from the builder stage
COPY --from=builder /app/report /report

# Copy the temporary directory from the builder stage
COPY --from=builder /tmp /tmp

# Set the working directory
WORKDIR /

# Run the quake-parser binary
ENTRYPOINT ["/quake-parser"]
