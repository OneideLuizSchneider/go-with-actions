############################
# STEP 1 build executable binary
############################
FROM golang:1.15.3-alpine as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update 
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
RUN update-ca-certificates

# Create apiuser
ENV USER=apiuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR $GOPATH/src/api/
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /go/bin/api .

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy our static executable
COPY --from=builder /go/bin/api /go/bin/api
COPY .env /go/bin/.env

# Use an unprivileged user.
USER apiuser:apiuser

ENV key=value

# Run the hello binary.
ENTRYPOINT ["/go/bin/api"]