FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /app

# Copy the code into the container
COPY . .

# Download dependency using go mod
RUN go mod download

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /build

# Copy binary from build to main folder
RUN cp /app/main .

# Copy file with env variables 
RUN cp /app/user-storage-service.env . || :

# Command to run when starting the container
CMD ["/build/main"]
