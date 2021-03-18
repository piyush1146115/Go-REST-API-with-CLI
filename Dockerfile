# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Piyush <piyush@appscode.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# install the dependencies
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app binaries from the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 10000

# Declare entry point of the docker command
ENTRYPOINT ["./main"]

# Command to run the executable
CMD ["startServer"]


#Type the following command to build the above image -
# $ docker build -t go-docker-api -f Dockerfile .

#To see the images
# $ docker image ls

#To run an image
# $ docker run -d -p 10000:10000 go-docker-api

#Finding running containers
# $ docker container ls

#To stop a running container
# $docker container stop <fff93d13a484(container_id)>