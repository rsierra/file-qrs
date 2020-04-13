FROM golang:1.14.2-alpine3.11 as builder

WORKDIR /go/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# There is a roblem with net lib bindings and CGO_ENABLED is needed
# https://stackoverflow.com/questions/36279253/go-compiled-binary-wont-run-in-an-alpine-docker-container-on-ubuntu-host
RUN CGO_ENABLED=0 go build -o server -v .

# ==============================
# Stage 2: Run the isolated build in a lightweight image
# ==============================

FROM alpine:3.11

WORKDIR /app

EXPOSE 8100

ARG HTPASSWD_FILE
ENV HTPASSWD_FILE ${HTPASSWD_FILE:-""}

ENTRYPOINT ["/app/server"]
CMD ["-p", "8100", "-d", "/files"]

COPY statics/ ./statics
COPY templates/ ./templates
COPY --from=builder /go/src/app/server ./

VOLUME /files
