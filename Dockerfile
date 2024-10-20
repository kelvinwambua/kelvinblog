FROM heroku/heroku:22

# Install Go
RUN curl -sL https://golang.org/dl/go1.21.3.linux-amd64.tar.gz | tar -C /usr/local -xzf -
ENV PATH=$PATH:/usr/local/go/bin

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Run the Go app
CMD ["go", "run", "cmd/server/main.go"]