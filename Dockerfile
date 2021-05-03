############################
# STEP 1 build executable binary
############################
ARG USER=genetics
FROM golang:1.16 AS builder
# ARG GITHUB_TOKEN
ARG USER

# Create user
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# create ssh directory
RUN mkdir ~/.ssh
RUN touch ~/.ssh/known_hosts
RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

WORKDIR $GOPATH/src/github.com/anchamber/genetics-system
COPY . .

# Using go get.
RUN go get -d -v 
# Build the binary.
RUN go build -o /genetics-system
RUN chown ${USER}:${USER} /genetics-system

############################
# STEP 2 build a small image
############################
FROM golang:1.16
ARG USER

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /genetics-system /genetics-system

# Use an unprivileged user.
USER ${USER}:${USER}

# Run the binary.
ENTRYPOINT /genetics-system