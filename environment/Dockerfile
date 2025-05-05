FROM golang:1.24.2

# Set destination for COPY
WORKDIR /

# Download Go modules
COPY go.mod ./
RUN go mod download

COPY *.go ./

# Build
RUN go build -o /app

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8118

# Run
CMD ["/app"]