# Use the official MariaDB image as the base image
#FROM mariadb:latest

# Set environment variables
#ENV MYSQL_ROOT_PASSWORD=mysecretpassword
#ENV MYSQL_DATABASE=db-test-api
#ENV MYSQL_USER=user-test-api
#ENV MYSQL_PASSWORD=mypassword

# Expose port 3306 to allow external connections to the database
#EXPOSE 3306

# Use an official Golang runtime as the base image
FROM golang:1.22

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Install dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
