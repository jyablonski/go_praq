FROM golang:1.22.5-alpine AS builder

WORKDIR /app

# example: https://martynassubonis.substack.com/p/optimizing-docker-images-for-python
# Copy dependency files first: changes less often, improving cache efficiency
COPY go.mod go.sum* ./

# Downloads Go module dependencies 
RUN go mod download
COPY *.go ./

# Builds the Go application with several optimizations:
# - `CGO_ENABLED=0` disables the use of cgo
# - `-ldflags="-w -s"` Reduces the size of the binary by stripping debug information and symbol tables.
# - `-a -installsuffix cgo` Forces a rebuild of packages with the cgo suffix.
# - `-o main` sets the output binary name
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix cgo -o main .

FROM scratch AS runtime

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]