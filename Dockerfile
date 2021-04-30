############################
# STEP 1 build executable binary
############################
FROM golang:1.16 AS builder
# ARG GITHUB_TOKEN

# Create appuser.
ENV USER=appuser
ENV UID=10001 

# create ssh directory
# RUN mkdir ~/.ssh
# RUN touch ~/.ssh/known_hosts
# RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/src/github.com/anchamber/genetics-system
COPY . .

# Using go get.
RUN go get -d -v 
# Build the binary.
RUN go build -o /go/bin/genetics-systems

############################
# STEP 2 build a small image
############################
FROM alpine 

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/genetics-systems /genetics-systems

# Use an unprivileged user.
USER appuser:appuser

EXPOSE 10000

# Run the binary.
ENTRYPOINT ["/genetics-systems"]