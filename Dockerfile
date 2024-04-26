# Use the official MariaDB image as the base image
FROM mariadb:latest

# Set environment variables
ENV MYSQL_ROOT_PASSWORD=mysecretpassword
ENV MYSQL_DATABASE=db-test-api
ENV MYSQL_USER=user-test-api
ENV MYSQL_PASSWORD=mypassword

# Expose port 3306 to allow external connections to the database
EXPOSE 3306