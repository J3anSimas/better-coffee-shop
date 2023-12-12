FROM mysql:latest

# Copy the SQL file to the Docker image
COPY create.sql /docker-entrypoint-initdb.d/